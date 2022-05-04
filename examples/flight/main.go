package main

import (
	"github.com/marvin-hansen/goC8"
	"github.com/marvin-hansen/goC8/requests/collection_req"
	utils2 "github.com/marvin-hansen/goC8/utils"
)

const (
	// client config
	apiKey   = "email.root.secretkey"
	endpoint = "https://YOUR-ID-us-west.paas.macrometa.io"
	fabric   = "SouthEastAsia"
	timeout  = 5 // http connection timeout in seconds
	// collection & graph config
	delete           = false
	verbose          = true
	silent           = false
	graph            = "airline"
	collectionID     = "cities"
	edgeCollectionID = "flights"
)

func main() {
	println("Create new config ")
	config := goC8.NewConfig(apiKey, endpoint, fabric, timeout)

	println("Create new client with config ")
	c := goC8.NewClient(config)

	println("Run setup")
	setup(c)

	println("Query: Document & Graph")
	query(c)

	if delete {
		println("Teardown: Delete Graph & Data")
		teardown(c)
	}
}

func query(c *goC8.Client) {
	var q = ""
	var msg = ""

	q = GetAllCitiesQuery()
	msg = "Get all cities."
	runQuery(c, q, msg)

	q = GetBreadthFirstQuery()
	msg = "Get all cities with a direct flight to New York."
	runQuery(c, q, msg)

	q = GetShortestPathQuery()
	msg = "Get the shortest path from San Francisco to Paris."
	runQuery(c, q, msg)

	q = GetShortestDistanceQuery()
	msg = "Get the distance on the shortest path from San Francisco to Paris."
	runQuery(c, q, msg)

	q = GetNearestCities()
	msg = "Get the 2 nearest cities to a specified latitude and longitude."
	runQuery(c, q, msg)

	q = GetCitiesMaxDistance()
	msg = "Get the cities that are no more than 2500km away from houston."
	runQuery(c, q, msg)

}

func runQuery(c *goC8.Client, q, msg string) {
	println(msg)
	res, err := c.Query(fabric, q, nil, nil)
	utils2.CheckError(err, "Error Query: "+q)
	utils2.PrintQuery(res, verbose)
}

func setup(c *goC8.Client) {
	goC8.CreateCollection(c, fabric, collectionID, collection_req.DocumentCollectionType, false)
	// We have to create a geo index before importing geoJson
	field := "location"
	geoJson := true
	_, err := c.Index.CreateGeoIndex(fabric, collectionID, field, geoJson)
	utils2.CheckError(err, "Error CreateNewDocument")
	utils2.DbgPrint("Create GeoIndex on: "+field, verbose)

	goC8.ImportData(c, fabric, collectionID, GetCityData(), silent)
	goC8.CreateCollection(c, fabric, edgeCollectionID, collection_req.EdgeCollectionType, false)
	goC8.ImportData(c, fabric, edgeCollectionID, GetFlightData(), silent)
	goC8.CreateGraph(c, fabric, graph, GetAirlineGraph())
}

func teardown(c *goC8.Client) {
	goC8.TeardownGraph(c, fabric, graph, true)
	goC8.TeardownCollection(c, fabric, collectionID)
	goC8.TeardownCollection(c, fabric, edgeCollectionID)
}

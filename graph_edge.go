package goC8

import (
	"github.com/marvin-hansen/goC8/requests/graph_req"
	"strings"
	"time"
)

// GetAllEdges
// Lists all edge collections within this graph.
// https://macrometa.com/docs/api#/operations/ListEdgedefinitions
func (c Client) GetAllEdges(fabric, graphName string) (response *graph_req.ResponseForGetAllEdges, err error) {
	if benchmark {
		defer TimeTrack(time.Now(), "GetAllEdges")
	}

	req := graph_req.NewRequestForGetAllEdges(fabric, graphName)
	response = graph_req.NewResponseForGetAllEdges()
	if err = c.request(req, response); err != nil {
		return nil, err
	}
	return response, nil
}

// AddEdgeCollection
// Adds an additional edge definition to the graph. This edge definition has to contain a collection and an array of each from and to vertex collections.
// An edge definition can only be added if this definition is either not used in any other graph, or it is used with exactly the same definition.
// *  edgeCollectionName: The name of the edge collection to be used.
// *  sourceVertex (string): One or many vertex collections that can contain source vertices.
// *  destinationVertex (string): One or many vertex collections that can contain target vertices.
// https://macrometa.com/docs/api#/operations/AddEdgedefinition
func (c Client) AddEdgeCollection(fabric, graphName, edgeCollectionName, sourceVertex, destinationVertex string) (response *graph_req.ResponseForAddEdgeCollection, err error) {
	if benchmark {
		defer TimeTrack(time.Now(), "AddEdgeCollection")
	}

	req := graph_req.NewRequestForAddEdgeCollection(fabric, graphName, edgeCollectionName, sourceVertex, destinationVertex)
	response = graph_req.NewResponseForAddEdgeCollection()
	if err = c.request(req, response); err != nil {
		return nil, err
	}
	return response, nil
}

// GetEdge
// Gets an edge from the given collection.
// https://macrometa.com/docs/api#/operations/GetAnEdge
func (c Client) GetEdge(fabric, graphName, collectionName, edgeKey string) (response *graph_req.ResponseForGetEdge, err error) {
	if benchmark {
		defer TimeTrack(time.Now(), "GetEdge")
	}

	req := graph_req.NewRequestForGetEdge(fabric, graphName, collectionName, edgeKey)
	response = graph_req.NewResponseForGetEdge()
	if err = c.request(req, response); err != nil {
		return nil, err
	}
	return response, nil
}

// CheckEdgeExists
// returns true if the edge exists in the given collection
func (c Client) CheckEdgeExists(fabric, graphName, collectionName, edgeKey string) (exists bool, err error) {
	if benchmark {
		defer TimeTrack(time.Now(), "CheckEdgeExists")
	}

	req := graph_req.NewRequestForGetEdge(fabric, graphName, collectionName, edgeKey)
	response := graph_req.NewResponseForGetEdge()
	if err = c.request(req, response); err != nil {
		if strings.Contains(err.Error(), "1202") { //  Code=404, Number=1202,  Error Message=document not found
			return false, nil
		} else {
			return false, err // Any other error
		}
	}
	return true, nil
}

// CreateEdge
// Creates a new edge in the collection
// sourceVertex: The source vertex of this edge. Has to be valid within the used edge definition.
// destinationVertex: The target vertex of this edge. Has to be valid within the used edge definition.
func (c Client) CreateEdge(fabric, graphName, edgeCollectionName, sourceVertex, destinationVertex string, returnNew bool) (response *graph_req.ResponseForCreateEdge, err error) {
	if benchmark {
		defer TimeTrack(time.Now(), "CreateEdge")
	}

	req := graph_req.NewRequestForCreateEdge(fabric, graphName, edgeCollectionName, sourceVertex, destinationVertex, returnNew)
	response = graph_req.NewResponseForCreateEdge()
	if err = c.request(req, response); err != nil {
		return nil, err
	}
	return response, nil
}

func (c Client) ReplaceEdge(fabric, graphName, edgeCollectionName, sourceVertex, destinationVertex string, dropCollections bool) (response *graph_req.ResponseForReplaceEdge, err error) {
	if benchmark {
		defer TimeTrack(time.Now(), "ReplaceEdge")
	}

	req := graph_req.NewRequestForReplaceEdge(fabric, graphName, edgeCollectionName, sourceVertex, destinationVertex, dropCollections)
	response = graph_req.NewResponseForReplaceEdge()
	if err = c.request(req, response); err != nil {
		return nil, err
	}
	return response, nil
}

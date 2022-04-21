package goC8

import r "github.com/marvin-hansen/goC8/requests/index_req"

// GetIndexes
// Fetches the list of all indexes of a collection.
func (c Client) GetIndexes(fabric, collectionName string) (response *r.ResponseForGetAllIndices, err error) {
	req := r.NewRequestForGetAllIndices(fabric, collectionName)
	response = r.NewResponseForGetAllIndices()
	if err = c.request(req, response); err != nil {
		return nil, err
	}
	return response, nil
}

// CreateFulltextIndex
// Creates fulltext index, if it does not already exist.
// field: Attribute to index. Must be of text format.
// minLength: Minimum character length of words to index.
func (c Client) CreateFulltextIndex(fabric, collectionName, field string, minLength int) (response *r.ResponseForCreateFulltextIndex, err error) {
	req := r.NewRequestForCreateFulltextIndex(fabric, collectionName, field, minLength)
	response = r.NewResponseForCreateFulltextIndex()
	if err = c.request(req, response); err != nil {
		return nil, err
	}
	return response, nil
}

// CreateGeoIndex
// Creates a geo index.
// Note: Geo indexes are always sparse, meaning that documents that do not contain the index attributes or have non-numeric values in the index attributes will not be indexed.
func (c Client) CreateGeoIndex(fabric, collectionName, field string, geoJson bool) (response *r.ResponseForCreateGeoIndex, err error) {
	req := r.NewRequestForCreateGeoIndex(fabric, collectionName, field, geoJson)
	response = r.NewResponseForCreateGeoIndex()
	if err = c.request(req, response); err != nil {
		return nil, err
	}
	return response, nil
}

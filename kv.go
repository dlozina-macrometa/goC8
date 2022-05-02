package goC8

import (
	"github.com/marvin-hansen/goC8/requests/kv_req"
	"time"
)

func (c Client) GetAllKVCollections(fabric string) (response *kv_req.KVResult, err error) {
	if benchmark {
		defer TimeTrack(time.Now(), "GetAllKVCollections")
	}
	req := kv_req.NewRequestForGetAllKVCollections(fabric)
	response = kv_req.NewKVResponse()
	if err = c.request(req, response); err != nil {
		return nil, err
	}
	return response, nil
}

func (c Client) CreateNewKVCollection(fabric, collectionName string, expiration bool, options *kv_req.CreateKVOptions) (response *kv_req.KVResult, err error) {
	if benchmark {
		defer TimeTrack(time.Now(), "CreateNewKVCollection")
	}

	if options == nil {
		options = kv_req.GetDefaultCreateKVOptions()
	}

	req := kv_req.NewRequestForCreateKVCollection(fabric, collectionName, expiration, options)
	response = kv_req.NewKVResponse()
	if err = c.request(req, response); err != nil {
		return nil, err
	}
	return response, nil
}

func (c Client) CountKVCollection(fabric, collectionName string) (response *kv_req.KVResult, err error) {
	if benchmark {
		defer TimeTrack(time.Now(), "GetAllKVCollections")
	}
	req := kv_req.NewRequestForCount(fabric, collectionName)
	response = kv_req.NewKVResponse()
	if err = c.request(req, response); err != nil {
		return nil, err
	}
	return response, nil
}

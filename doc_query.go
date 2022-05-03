package goC8

import (
	"github.com/marvin-hansen/goC8/src/requests/query_req"
	"time"
)

func (c Client) Query(fabric, query string, bindVars map[string]interface{}, options *query_req.CursorOptions) (res *query_req.Cursor, err error) {
	if benchmark {
		defer TimeTrack(time.Now(), "Query")
	}

	if options == nil {
		options = query_req.NewDefaultCursorOptions()
	}
	ttl := c.getQueryTTL()
	req := query_req.NewRequestForCreateCursor(fabric, query, bindVars, options, ttl)

	// Create a cursor for the query onfrom the server
	response := query_req.NewResponseForCreateCursor()
	if err = c.request(req, response); err != nil {
		return nil, err
	}

	// We cast into a cursor here to allow for updates from NextCursor
	res = query_req.NewCursorFromCreateCursor(response)

	// check for more to come
	if response.HasMore {
		for {
			// request update for the cursor
			reqNext := query_req.NewRequestForReadNextCursor(fabric, res.Id)
			responseNext := query_req.NewResponseForReadNextCursor()
			if err = c.request(reqNext, responseNext); err != nil {
				return nil, err
			}

			// updated cursor with next result
			res.Update(responseNext)

			// stop when there are no more results, delete cursor and stop
			if responseNext.HasMore == false {

				// Delete the cursor from the server
				reqDel := query_req.NewRequestForDeleteCursor(fabric, response.Id)
				resDel := query_req.NewResponseForDeleteCursor()
				if err = c.request(reqDel, resDel); err != nil {
					return nil, err
				}

				// stop the loop
				break
			}
		}
	}

	return res, nil
}

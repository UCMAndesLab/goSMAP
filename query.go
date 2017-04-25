package gosMAP

import (
	"bytes"
	"encoding/json"
	"fmt"
	"hash/fnv"
	"io/ioutil"
	"net/http"

	"github.com/bradfitz/gomemcache/memcache"
)

func (conn *Connection) query(q string) []byte {
	url := fmt.Sprintf("%sapi/query?key=%s", conn.Url, conn.APIkey)
	response, err := http.Post(url, "text/smap", bytes.NewBufferString(q))
	if err != nil {
		return nil
	}

	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil
	}
	return contents
}

func queryKey(q string) string {
	h := fnv.New32a()
	h.Write([]byte("query"))
	h.Write([]byte(q))
	return fmt.Sprintf("%d", h.Sum32())
}

func (conn *Connection) cacheGet(q string) (b []byte, err error) {
	key := queryKey(q)

	if conn.Mc != nil {
		item, lerr := conn.Mc.Get(key)
		if lerr == nil && item != nil {
			return item.Value, nil
		}
		err = fmt.Errorf("Cache miss")
	} else {
		err = fmt.Errorf("Not using cache")
	}
	return b, err

}

// Query using sMAP querying language
//
// See http://www.cs.berkeley.edu/~stevedh/smap2/archiver.html#archiverquery for further
// documentation to retrieve data. The contents will be returned as json text if success,
// and on some errors a text file
func (conn *Connection) Query(q string) (results []Data, err error) {
	key := queryKey(q)
	var clean []Data
	b, err := conn.cacheGet(q)
	if err == nil {
		// Cache Hit
		err = json.Unmarshal(b, &clean)
	} else {
		// Cache Miss
		b := conn.query(q)

		raw := make([]RawData, 0)
		err = json.Unmarshal(b, &raw)
		clean = rawDataToClean(raw)

		if err == nil && conn.Mc != nil {
			var savebin []byte
			// Save in the cache
			savebin, err = json.Marshal(clean)
			// Save
			if err == nil {
				err = conn.Mc.Set(&memcache.Item{Key: key, Value: savebin, Expiration: 3600})
			} else {
				return clean, err
			}
		}
	}
	return clean, err
}

// QueryList is a query that returns a string array. This is necessary for
// for all ```select distinct``` queries.
func (conn *Connection) QueryList(q string) (results []string, err error) {
	key := queryKey(q)
	b, err := conn.cacheGet(q)
	if err == nil {
		// Cache hit
		err = json.Unmarshal(b, &results)

	} else {
		// Cache miss
		b := conn.query(q)
		err = json.Unmarshal(b, &results)

		//Save
		if err == nil && conn.Mc != nil {
			conn.Mc.Set(&memcache.Item{Key: key, Value: b, Expiration: 3600})
		}
	}
	return results, err

}

// Get data given a uuid.
//
// starttime and endtime are unix times in seconds based off of the epoch. A
// starttime of 0 will get data starting from the first entry and a endtime of
// 0 will have no upper bound. Limit is the number of values to be retrieved.
// Set to 0 if you do not want a limit.
//
// Although the return is an array of SMAPData, typically there should only be
// one value with the given uuid.
func (conn *Connection) Get(uuid string, starttime int, endtime int, limit int) ([]Data, error) {
	starttimeStr := smap_time(starttime)

	// endtime doesn't work
	if endtime == 0 {
		endtime = 2000000000000
	}
	endtimeStr := smap_time(endtime)

	url := fmt.Sprintf("%sapi/data/uuid/%s?startime=%s&endtime=%s&limit=%d", conn.Url, uuid, starttimeStr, endtimeStr, limit)

	return pullData(url)
}

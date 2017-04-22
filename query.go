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
		if lerr != nil && item != nil {
			return item.Value, nil
		}
		err = fmt.Errorf("Cache miss")
	} else {
		err = fmt.Errorf("Not using cache")
	}
	return b, err

}

// Use sMAP querying language
//
// See http://www.cs.berkeley.edu/~stevedh/smap2/archiver.html#archiverquery for further
// documentation to retrieve data. The contents will be returned as json text if success,
// and on some errors a text file
func (conn *Connection) Query(q string) (results []Data, err error) {
	key := queryKey(q)
	var clean []Data
	b, err := conn.cacheGet(q)
	if err == nil {
		fmt.Printf("CACHE Hit!%s%s\n", key, string(b))
		// Cache Hit
		err = json.Unmarshal(b, &clean)
	} else {
		err = nil
		// Cache Miss
		fmt.Printf("CACHE MISS!%s\n", key)
		b := conn.query(q)

		raw := make([]RawData, 0)
		err := json.Unmarshal(b, &raw)
		clean = rawDataToClean(raw)

		if err != nil {
			fmt.Printf("CACHE ERROR:!%s, Q:%s\n", err.Error(), q)
		}

		if conn.Mc == nil {
			fmt.Printf("Memcache not connected\n")
		}

		if err == nil && conn.Mc != nil {
			// Save in the cache
			b, err := json.Marshal(clean)

			// Save
			if err == nil {
				err := conn.Mc.Set(&memcache.Item{Key: key, Value: b, Expiration: 3600})
				if err != nil {
					fmt.Printf("CACHE Failed!%s\n", err.Error())
				} // Return error
			} else {
				fmt.Printf("CACHE ERROR:!%s, Query: %s\n", err.Error(), q)
				return clean, err
			}
		}
	}
	return clean, err
}

// Similar to Query, but QueryList returns a string array. This is necessary for
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
	starttime_str := smap_time(starttime)

	// endtime doesn't work
	if endtime == 0 {
		endtime = 2000000000000
	}
	endtime_str := smap_time(endtime)

	url := fmt.Sprintf("%sapi/data/uuid/%s?startime=%s&endtime=%s&limit=%d", conn.Url, uuid, starttime_str, endtime_str, limit)

	return pullData(url)
}

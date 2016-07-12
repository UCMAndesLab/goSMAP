package gosMAP

import (
  "bytes"
  "fmt"
  "net/http"
  "io/ioutil"
  "encoding/json"
)

func (conn *Connection) query(q string)([]byte){
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

// Use sMAP querying language
//
// See http://www.cs.berkeley.edu/~stevedh/smap2/archiver.html#archiverquery for further
// documentation to retrieve data. The contents will be returned as json text if success,
// and on some errors a text file
func (conn *Connection) Query(q string) ([]Data, error){
  b := conn.query(q)
  d := make([]RawData, 0)
  err := json.Unmarshal(b, &d)

  return rawDataToClean(d), err
}

// Similar to Query, but QueryList returns a string array. This is necessary for
// for all ```select distinct``` queries.
func (conn *Connection) QueryList(q string) ([]string, error){
  b := conn.query(q)
  d := make([]string, 0)
  err := json.Unmarshal(b, &d)

  return d, err
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
func (conn *Connection) Get(uuid string, starttime int, endtime int, limit int) ([]Data, error){
  starttime_str := smap_time(starttime)

  // endtime doesn't work
  if endtime == 0 {
    endtime = 2000000000000
  }
  endtime_str :=  smap_time(endtime)

  url := fmt.Sprintf("%sapi/data/uuid/%s?startime=%s&endtime=%s&limit=%d", conn.Url, uuid, starttime_str, endtime_str, limit)

  return pullData(url)
}

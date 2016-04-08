package gosMAP

import (
  "bytes"
  "fmt"
  "net/http"
  "encoding/json"
  "io/ioutil"
)

func (conn *sMAPConnection) Query(q string) ([]byte){
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

// We will not cache diddly
func (conn *sMAPConnection) Data_uuid(uuid string, starttime int, endtime int, limit int) ([]sMAPData, error){
  starttime_str := smap_time(starttime)
  d := make([]sMAPData,0)

  // endtime doesn't work
  if endtime == 0 {
    endtime = 2000000000000
  }
  endtime_str :=  smap_time(endtime)

  url := fmt.Sprintf("%sapi/data/uuid/%s?startime=%s&endtime=%s&limit=%d", conn.Url, uuid, starttime_str, endtime_str, limit)

  response, err := http.Get(url)
  if err != nil {
    return d, err
  }

  defer response.Body.Close()
  contents, err := ioutil.ReadAll(response.Body)
  if err != nil {
    return d, err
  }

  err = json.Unmarshal(contents, &d)
  if err != nil {
    return d, err
  }

  return d, nil
}

func smap_time(t int) string{
  if t == 0{
    return ""
  }else{
    // smap measures from microseconds since epoch so times it by 1000.
    return fmt.Sprintf("%d", t*1000);
  }
}

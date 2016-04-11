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

  // endtime doesn't work
  if endtime == 0 {
    endtime = 2000000000000
  }
  endtime_str :=  smap_time(endtime)

  url := fmt.Sprintf("%sapi/data/uuid/%s?startime=%s&endtime=%s&limit=%d", conn.Url, uuid, starttime_str, endtime_str, limit)

  return pullData(url)
}

func pullData(url string) ([]sMAPData, error){
  d := make([]rawsMAPData,0)
  r := make([]sMAPData,0)
  response, err := http.Get(url)
  if err != nil {
    return r, err
  }

  defer response.Body.Close()
  contents, err := ioutil.ReadAll(response.Body)
  if err != nil {
    return r, err
  }

  err = json.Unmarshal(contents, &d)
  if err != nil {
    return r, err
  }

  r = rawDataToClean(d)
  return r, nil
}

func rawDataToClean(dirty []rawsMAPData) []sMAPData{
  r := make([]sMAPData, len(dirty))

  for i,d := range dirty{
    r[i].Uuid = d.Uuid
    r[i].Readings = make([]readPair, len(d.Readings))

    for j,entry := range d.Readings{
      r[i].Readings[j].value,_ = entry[1].Float64()
    }
  }
  return r;
}
func smap_time(t int) string{
  if t == 0{
    return ""
  }else{
    // smap measures from microseconds since epoch so times it by 1000.
    return fmt.Sprintf("%d", t*1000);
  }
}

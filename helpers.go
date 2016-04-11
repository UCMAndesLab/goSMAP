package gosMAP

import (
  "fmt"
  "net/http"
  "encoding/json"
  "io/ioutil"
  "time"
)

func pullData(url string) ([]SMAPData, error){
  d := make([]rawsMAPData,0)
  r := make([]SMAPData,0)
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

func rawDataToClean(dirty []rawsMAPData) []SMAPData{
  r := make([]SMAPData, len(dirty))

  for i,d := range dirty{
    r[i].Uuid = d.Uuid
    r[i].Readings = make([]readPair, len(d.Readings))

    for j,entry := range d.Readings{
      r[i].Readings[j].Value,_ = entry[1].Float64()

      rawT,_ := entry[0].Float64()
      unixT := int64(rawT/1000)
      r[i].Readings[j].Time = time.Unix(unixT, 0)
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

package gosMAP

import (
  "fmt"
  "net/http"
  "encoding/json"
  "io/ioutil"
)

// We will not cache diddly
func (conn *sMAPConnection) Prev(uuid string) ([]sMAPData, error){
  d := make([]sMAPData,0)
  url := fmt.Sprintf("%sapi/prev/uuid/%s", conn.Url, uuid)

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

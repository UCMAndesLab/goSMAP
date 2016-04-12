package gosMAP

import (
  "bytes"
  "fmt"
  "encoding/json"
  "net/http"
  "io/ioutil"

)

func (conn *sMAPConnection) Post(data map[string]RawsMAPData) error{
  m,e := json.Marshal(data)
  if e != nil{
    return e
  }
  url := fmt.Sprintf("%sadd/%s", conn.Url, conn.APIkey)
//  fmt.Printf(string(m))
  req, err := http.NewRequest("POST", url, bytes.NewBuffer(m))
  if err != nil{
    return e
  }
  req.Header.Set("Content-Type", "application/json")
  client := &http.Client{}

  resp, err := client.Do(req)
  if err != nil{
    return e
  }
  defer resp.Body.Close()
  contents, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return nil
  }

  if resp.Status != "200 OK"{
    return fmt.Errorf("Status Error: %s\n Message:%s\n", resp.Status, contents)
  }
  return nil
}

package gosMAP

import (
    "fmt"
    "net/http"
    "encoding/json"
    "io/ioutil"
)

type sMAPConnection struct{
  Url string
  APIkey string
}

type rootPage struct{
  Contents []string
}

func Validateconnection(conn sMAPConnection)(error){
  // is url valid
  response, err := http.Get(conn.Url)
  if err != nil {
    return err
  }

  // Is the data in a json form we expect?
  defer response.Body.Close()
  contents, err := ioutil.ReadAll(response.Body)
  if err != nil {
    return err
  }

  var m rootPage
  err = json.Unmarshal(contents, &m)
  if err != nil {
    return err
  }

  if len(m.Contents)>0 && m.Contents[0] == "add" {
    return nil
  }else{
    return fmt.Errorf("Not Valid sMAP Archiver")
  }
}

func Connect(url string, key string)(sMAPConnection, error){
  conn := sMAPConnection{
    Url:url,
    APIkey:key,
  }
  err := Validateconnection(conn)
  return conn, err
}

package gosMAP

import (
    "fmt"
    "net/http"
    "encoding/json"
    "io/ioutil"
    "github.com/bradfitz/gomemcache/memcache"
)

type Connection struct{
  Url string
  APIkey string
  Mc *memcache.Client
}

type rootPage struct{
  Contents []string
}

func validateConnection(conn Connection)(error){
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

func Connect(url string, key string)(Connection, error){
  if url[len(url)-1] != '/'{
    url = url + "/"
  }
  conn := Connection{
    Url:url,
    APIkey:key,
  }
  err := validateConnection(conn)
  return conn, err
}

func (conn *Connection) ConnectMemcache(server string){
    conn.Mc = memcache.New(server)
}

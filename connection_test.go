package gosMAP

import (
    "fmt"
    "github.com/go-ini/ini"
    "testing"
)

var server string
var apiKey string
var uuid string
func init(){

  cfg, err := ini.Load("test.mine.ini")
  if err != nil{
    cfg, err = ini.Load("test.ini")
    if err != nil{
      panic(err)
    }
  }

  server = cfg.Section("MAIN").Key("server").String()
  apiKey = cfg.Section("MAIN").Key("apiKey").String()
  uuid   =  cfg.Section("MAIN").Key("uuid").String()
}

func TestConnection(t *testing.T) {
    _,err := Connect(server, apiKey)
    if err != nil{
      t.Error(err.Error())
    }
}

func TestDataCollection(t *testing.T){
    conn,err := Connect(server, apiKey)
    if err == nil{
      fmt.Printf("%s\n", conn.Url)
      _, err := conn.Data_uuid(uuid, 0, 0, 10)

      if err !=nil {
        t.Error(err.Error())
      }
    }else{
      t.Error(err.Error())
    }
}

func TestQuery(t *testing.T){
    conn,err := Connect(server, apiKey)
    if err != nil{
      t.Error(err.Error())
    }
    conn.query(fmt.Sprintf("select * where uuid='%s'", uuid))
}

func TestTags(t *testing.T){
    conn,err := Connect(server, apiKey)
    if err != nil{
      t.Error(err.Error())
    }
    d := conn.Tags(uuid)
    if (uuid != d[0].Uuid){
      t.Error("UUID Mismatch")
    }
    fmt.Printf("%s\n", string(d[0].Metadata["SourceName"].(string)))
}

func BenchmarkTagsCache(b *testing.B) {
    conn,_ := Connect(server, apiKey)
    for i := 0; i < b.N; i++ {
      conn.Tags(uuid)
    }
}

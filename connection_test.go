package gosMAP_test

import (
    "fmt"
    "github.com/go-ini/ini"
    "testing"
    "github.com/alexbeltran/gosMAP"
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
    _,err := gosMAP.Connect(server, apiKey)
    if err != nil{
      t.Error(err.Error())
    }
}

// There was a bug where a server without a trailing "/" would cause a crash.
// Lets check to see if it works both ways.
func TestConnectionTrailSlash(t *testing.T) {
    s := server

    // Without slash
    if s[len(s)-1] == '/'{
      s = s[:len(s)-1];
    }
    conn,err := gosMAP.Connect(s, apiKey)
    _, err = conn.Get(uuid, 0, 0, 10)
    if err != nil{
      t.Error(err.Error())
    }

    // With Slash
    s = s + "/"
    conn,err = gosMAP.Connect(s, apiKey)
    _, err = conn.Get(uuid, 0, 0, 10)
    if err != nil{
      t.Error(err.Error())
    }
}

func TestDataCollection(t *testing.T){
    conn,err := gosMAP.Connect(server, apiKey)
    if err == nil{
      fmt.Printf("%s\n", conn.Url)
      _, err := conn.Get(uuid, 0, 0, 10)

      if err !=nil {
        t.Error(err.Error())
      }
    }else{
      t.Error(err.Error())
    }
}

func TestQuery(t *testing.T){
    conn,err := gosMAP.Connect(server, apiKey)
    if err != nil{
      t.Error(err.Error())
    }
    conn.Query(fmt.Sprintf("select * where uuid='%s'", uuid))
}

func TestTags(t *testing.T){
    conn,err := gosMAP.Connect(server, apiKey)
    if err != nil{
      t.Error(err.Error())
    }
    d := conn.Tags(uuid)
    if (uuid != d[0].Uuid){
      t.Error("UUID Mismatch")
    }
    fmt.Printf("%s\n", string(d[0].Metadata["SourceName"].(string)))
}

func TestPrev(t *testing.T){
  conn,err := gosMAP.Connect(server, apiKey)
    if err != nil{
      t.Error(err.Error())
    }
    d,err := conn.Prev(uuid)
    if err != nil{
      t.Error(err.Error())
    }
    fmt.Printf("%f\n", d[0].Readings[0].Value)
    fmt.Printf("%s\n", d[0].Readings[0].Time)
}


func BenchmarkTagsCache(b *testing.B) {
    conn,_ := gosMAP.Connect(server, apiKey)
    for i := 0; i < b.N; i++ {
      conn.Tags(uuid)
    }
}

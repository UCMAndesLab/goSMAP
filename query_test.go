package gosMAP_test

import (
    "github.com/go-ini/ini"
    "testing"
    "github.com/alexbeltran/gosMAP"
)

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

// Function to test a gosMAP.QueryList()
func testQueryList(t *testing.T, query string){
    conn,err := gosMAP.Connect(server, apiKey)
    if err != nil{
      t.Error(err.Error())
    }

    d,err  := conn.QueryList(query)
    if err != nil{
        t.Error(err.Error())
    }
    if len(d) == 0{
        t.Error("No values retured on query %s", query)
    }
    t.Log(d)
}

// Function to test a gosMAP.Query()
// TODO: Combine testQueryList and testQuery
func testQuery(t *testing.T, query string){
    conn,err := gosMAP.Connect(server, apiKey)
    if err != nil{
      t.Error(err.Error())
    }

    d,err  := conn.Query(query)
    if err != nil{
        t.Error(err.Error())
    }
    if len(d) == 0{
        t.Error("No values retured on query %s", query)
    }
    t.Log(d[0])
}

func TestAllQuery(t *testing.T){
    testQuery(t, "select *")
}

func TestDistinct(t *testing.T){
    testQueryList(t, "select distinct")
}


func TestDistinctBuildings(t *testing.T){
    testQueryList(t, "select distinct Metadata/Location/Building")
}


func TestWeirdBuilding(t *testing.T){
    testQueryList(t, "select distinct where Metadata/SourceName = 'University of California, Merced'")
}

// FIXME: This returns a ton of null values. I have no idea why. This is likely to our data.
func TestBuildingCity(t *testing.T){
    testQuery(t, "select Metadata/Location/Building, Metadata/Location/City")
}

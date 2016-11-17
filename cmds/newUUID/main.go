package main

import (
    "github.com/alexbeltran/gosMAP"
    "encoding/json"
    "time"
    "fmt"
)

func generate(testUUID string) map[string]gosMAP.RawData{
  d := make(map[string]gosMAP.RawData)

  path := "/Kolligian Library/HVAC/West Wing/Third Floor/VAV-324 RM-336/Zone Temperature"

  // Generate Fake Data
  entry := [][]json.Number{}
  for i :=0; i < 1; i++{
    entry = append(entry,[]json.Number{gosMAP.TimeToNumber(time.Now().Add(time.Duration(i)*time.Second)), json.Number(fmt.Sprintf("%d",i))})
  }

  // Add Metadata
  meta := make(map[string]interface{})
  meta["SourceName"] = "University of California, Merced"
  d[path] = gosMAP.RawData{
    Uuid : testUUID,
    Readings: entry,
    Properties:&gosMAP.TagsProperties{
      Timezone:"America/Los_Angeles",
      UnitofMeasure:"Fahrenheit",
      ReadingType:"double",
    },
    Metadata:meta,
    }
  return d
}


func main(){
  conn,err := gosMAP.Connect("http://mercury:8079/", "9te21wWjfSZuq9aYqPqfwa3S8qBYAWP5zlav")
  testUUID :="0899f968-4941-4a5d-82d3-17254e107859"

  if err != nil{
    panic(err.Error())
  }
  d := generate(testUUID)

  err = conn.Post( d)
  if err != nil{
    panic(err.Error())
  }
}

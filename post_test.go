package gosMAP_test

import (
    "fmt"
    "time"
    "testing"
    "encoding/json"
    "../gosMAP"
)
// Test are ran in alphabetic order, and appended. We don't need to redeclare the server and apikey
func generateFakeSMAPData(testUUID string) map[string]gosMAP.RawData{
  d := make(map[string]gosMAP.RawData)

  path := "/pizza/alex"

  // Generate Fake Data
  entry := [][]json.Number{}
  for i :=0; i < 4; i++{
    entry = append(entry,[]json.Number{gosMAP.TimeToNumber(time.Now().Add(time.Duration(i)*time.Second)), json.Number(fmt.Sprintf("%d",i))})
  }

  // Add Metadata
  meta := gosMAP.Metadata{
      SourceName: "ThePizza",
  }

  d[path] = gosMAP.RawData{
    Uuid : testUUID,
    Readings: entry,
    Properties:&gosMAP.TagsProperties{
      Timezone:"America/Los_Angeles",
      UnitofMeasure:"Pizzas Eaten",
      ReadingType:"double",
    },
    Metadata:&meta,
    }
  return d
}


func TestPost(t *testing.T){
  conn,err := gosMAP.Connect(server, apiKey)
  testUUID :="fc17ecdc-135a-4b07-8ac7-555efb2df7d5"

  if err != nil{
    t.Error(err.Error())
  }
  d := generateFakeSMAPData(testUUID)

  err = conn.Post( d)
  if err != nil{
    t.Error(err.Error())
  }
  if ! conn.UUIDExists(testUUID){
    t.Error("Data was not posted correctly.")
  }
  conn.Delete(testUUID)

  if conn.UUIDExists(testUUID){
    t.Error("UUID was not delted correctly")
  }
}

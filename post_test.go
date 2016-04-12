package gosMAP

import (
    "fmt"
    "time"
    "testing"
    "encoding/json"
)
// Test are ran in alphabetic order, and appended. We don't need to redeclare the server and apikey
func generateFakeSMAPData(testUUID string) map[string]RawsMAPData{
  d := make(map[string]RawsMAPData)

  path := "/pizza/alex"

  // Generate Fake Data
  entry := [][]json.Number{}
  for i :=0; i < 4; i++{
    entry = append(entry,[]json.Number{TimeToNumber(time.Now().Add(time.Duration(i)*time.Second)), json.Number(fmt.Sprintf("%d",i))})
  }

  // Add Metadata
  meta := make(map[string]interface{})
  meta["SourceName"] = "ThePizza"
  d[path] = RawsMAPData{
    Uuid : testUUID,
    Readings: entry,
    Properties:SMAPTagsProperties{
      Timezone:"America/Los_Angeles",
      UnitofMeasure:"Pizzas Eaten",
      ReadingType:"double",
    },
    Metadata:meta,
    }
  return d
}


func TestPost(t *testing.T){
  conn,err := Connect(server, apiKey)
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

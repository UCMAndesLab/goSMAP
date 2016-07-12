package gosMAP
import(
  "encoding/json"
  "time"
)
// RawsMAPData is what is returned by a request from the archiver. To make this
// look cleaner, this value is typically converted to a type SMAPData as a return.
type RawData struct{
  Uuid string `json:"uuid"`
  Readings [][]json.Number  `json:"Readings"`
  Properties *TagsProperties   `json:",omitempty"`
  Metadata map[string]interface{}  `json:",omitempty"`
}

type Data struct{
  Uuid string `json:"uuid,omitempty"`
  Readings []ReadPair `json:"Readings,omitempty"`
  Properties *TagsProperties `json:",omitempty"`
  Metadata map[string]interface{} `json:",omitempty"`
}

// Each value returned by sMAP is a pair of time and float values.
type ReadPair struct{
  Time time.Time
  Value float64
}

type TagsProperties struct{
  Timezone string   `json:",omitempty"`
  UnitofMeasure string  `json:",omitempty"`
  ReadingType string   `json:",omitempty"`
}

// This is the bare minimium of what sMAP returns to you as
type Tags struct{
    Uuid string `json:"uuid"`
    Properties TagsProperties
    Path string
    Metadata map[string]interface{}
}

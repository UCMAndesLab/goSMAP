package gosMAP
import(
  "encoding/json"
  "time"
)
// RawsMAPData is what is returned by a request from the archiver. To make this
// look cleaner, this value is typically converted to a type SMAPData as a return.
type RawsMAPData struct{
  Uuid string `json:"uuid"`
  Readings [][]json.Number  `json:"Readings"`
  Properties SMAPTagsProperties
  Metadata map[string]interface{}
}

type SMAPData struct{
  Uuid string
  Readings []ReadPair
}

// Each value returned by sMAP is a pair of time and float values.
type ReadPair struct{
  Time time.Time
  Value float64
}

type SMAPTagsProperties struct{
  Timezone string
  UnitofMeasure string
  ReadingType string
}

// This is the bare minimium of what sMAP returns to you as
type SMAPTags struct{
    Uuid string `json:"uuid"`
    Properties SMAPTagsProperties
    Path string
    Metadata map[string]interface{}
}

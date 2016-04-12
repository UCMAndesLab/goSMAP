package gosMAP
import(
  "encoding/json"
  "time"
)
type RawsMAPData struct{
  Uuid string `json:"uuid"`
  Readings [][]json.Number  `json:"Readings"`
  Properties sMAPTagsProperties
  Metadata map[string]interface{}
}

type SMAPData struct{
  Uuid string
  Readings []ReadPair
}

type ReadPair struct{
  Time time.Time
  Value float64
}

type sMAPTagsProperties struct{
  Timezone string
  UnitofMeasure string
  ReadingType string
}

type SMAPTags struct{
    Uuid string `json:"uuid"`
    Properties sMAPTagsProperties
    Path string
    Metadata map[string]interface{}
}

package gosMAP
import(
  "encoding/json"
  "time"
)
type rawsMAPData struct{
  Uuid string `json:"uuid"`
  Readings [][]json.Number  `json:"Readings"`
}

type SMAPData struct{
  Uuid string
  Readings []readPair
}

type readPair struct{
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

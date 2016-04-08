package gosMAP
import(
  "encoding/json"
)
type sMAPData struct{
  Uuid string `json:"uuid"`
  Readings [][]json.Number  `json:"Readings"`
}

type sMAPTagsProperties struct{
  Timezone string
  UnitofMeasure string
  ReadingType string
}

type sMAPTags struct{
    Uuid string `json:"uuid"`
    Properties sMAPTagsProperties
    Path string
    Metadata map[string]interface{}
}

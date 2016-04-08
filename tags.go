package gosMAP

import (
  "fmt"
  "encoding/json"
  "github.com/bradfitz/gomemcache/memcache"
)

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

func (conn *sMAPConnection) Tags(uuid string) []sMAPTags{
  key := "tag_"+uuid
  mc := memcache.New("127.0.0.1:11211")
  item, err := mc.Get(key)

  var s []byte;
  if err == nil {
    // Cache Hit
    s = item.Value;
  }else{
    // Cache Miss
    q := fmt.Sprintf("select * where uuid='%s'", uuid)
    s = conn.Query(q)
    mc.Set(&memcache.Item{Key: key, Value: s, Expiration: 3600})
  }
  d := make([]sMAPTags,0)
  json.Unmarshal(s, &d)
  return d;
}

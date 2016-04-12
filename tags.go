package gosMAP

import (
  "fmt"
  "encoding/json"
  "github.com/bradfitz/gomemcache/memcache"
)

func tagKey(uuid string) string{
  return "tag_"+uuid
}

func (conn *SMAPConnection) Tags(uuid string) []SMAPTags{
  key := tagKey(uuid)
  item, err := conn.Mc.Get(key)

  var s []byte;
  if err == nil {
    // Cache Hit
    s = item.Value;
  }else{
    // Cache Miss
    q := fmt.Sprintf("select * where uuid='%s'", uuid)
    s = conn.Query(q)
    conn.Mc.Set(&memcache.Item{Key: key, Value: s, Expiration: 3600})
  }
  d := make([]SMAPTags,0)
  json.Unmarshal(s, &d)
  return d;
}

func (conn *SMAPConnection) UUIDExists(uuid string) bool{
  // Check to see if the data was put in
  // Remove tag from cache
  conn.Mc.Delete(tagKey(uuid))
  tags := conn.Tags(uuid)
  return len(tags) > 0 && len(tags[0].Path) > 0
}

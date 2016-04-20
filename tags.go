package gosMAP

import (
  "fmt"
  "encoding/json"
  "github.com/bradfitz/gomemcache/memcache"
)

func tagKey(uuid string) string{
  return "tag_"+uuid
}

func (conn *Connection) Tags(uuid string) []Tags{
  key := tagKey(uuid)
  item, err := conn.Mc.Get(key)

  var s []byte;
  if err == nil {
    // Cache Hit
    s = item.Value;
  }else{
    // Cache Miss
    q := fmt.Sprintf("select * where uuid='%s'", uuid)
    s = conn.query(q)
    conn.Mc.Set(&memcache.Item{Key: key, Value: s, Expiration: 3600})
  }
  d := make([]Tags,0)
  json.Unmarshal(s, &d)
  return d;
}

// Tag is similar to Tags, however only a single tag is returned
func (conn *Connection) Tag(uuid string) Tags{
  d := conn.Tags(uuid)
  return d[0];
}

func (conn *Connection) UUIDExists(uuid string) bool{
  // Check to see if the data was put in
  // Remove tag from cache
  conn.Mc.Delete(tagKey(uuid))
  tags := conn.Tags(uuid)
  return len(tags) > 0 && len(tags[0].Path) > 0
}

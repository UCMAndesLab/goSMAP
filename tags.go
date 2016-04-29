package gosMAP

import (
  "fmt"
  "encoding/json"
  "github.com/bradfitz/gomemcache/memcache"
)

func tagKey(uuid string) string{
  return "tag_"+uuid
}

func (conn *Connection) query_tags(uuid string) []byte{
    q := fmt.Sprintf("select * where uuid='%s'", uuid)
    return conn.query(q)
}

//  Use Memcache to get data
func (conn *Connection) memcache_tags(uuid string) ([]byte, error){
  var s []byte;
  if conn.Mc == nil{
    return s, fmt.Errorf("Memcache not connected")
  }

  key := tagKey(uuid)
  item, err := conn.Mc.Get(key)
  if err == nil {
    // Cache Hit
    s = item.Value;
  }else{
    // Cache Miss
    s = conn.query_tags(uuid)
    conn.Mc.Set(&memcache.Item{Key: key, Value: s, Expiration: 3600})
  }
  return s, nil
}

// Get all tags associate with uuid
func (conn *Connection) Tags(uuid string) []Tags{
  s, err := conn.memcache_tags(uuid);
  if err != nil{
    s = conn.query_tags(uuid);
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
  // Remove tag from cache just in case it was removed.
  if conn.Mc != nil{
    conn.Mc.Delete(tagKey(uuid))
  }
  tags := conn.Tags(uuid)
  return len(tags) > 0 && len(tags[0].Path) > 0
}

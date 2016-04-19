package gosMAP

import (
  "fmt"
)
func (conn *Connection) Delete(uuid string) error{
  q := fmt.Sprintf("delete where uuid = '%s'", uuid)
  s := conn.Query(q)
  conn.Mc.Delete(tagKey(uuid))
  fmt.Printf(string(s))
  return nil
}

package gosMAP

import (
  "fmt"
)
func (conn *sMAPConnection) Delete(uuid string) error{
  q := fmt.Sprintf("delete where uuid = '%s'", uuid)
  s := conn.Query(q)
  conn.mc.Delete(tagKey(uuid))
  fmt.Printf(string(s))
  return nil
}

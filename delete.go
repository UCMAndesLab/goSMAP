package gosMAP

import (
  "fmt"
)
func (conn *Connection) Delete(uuid string) error{
  q := fmt.Sprintf("delete where uuid = '%s'", uuid)
  s := conn.query(q)
  if conn.Mc != nil{
    conn.Mc.Delete(tagKey(uuid))
  }
  fmt.Printf(string(s))
  return nil
}

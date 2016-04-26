package gosMAP

import (
  "fmt"
)

// Return the last value from given uuid
func (conn *Connection) Prev(uuid string) ([]Data, error){
  url := fmt.Sprintf("%sapi/prev/uuid/%s?key=%s", conn.Url, uuid, conn.APIkey)
  return pullData(url)
}

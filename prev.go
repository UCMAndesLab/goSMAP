package gosMAP

import (
  "fmt"
)

// We will not cache diddly
func (conn *sMAPConnection) Prev(uuid string) ([]SMAPData, error){
  url := fmt.Sprintf("%sapi/prev/uuid/%s", conn.Url, uuid)
  return pullData(url)
}

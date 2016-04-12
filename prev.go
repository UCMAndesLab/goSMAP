package gosMAP

import (
  "fmt"
)

// Return the last value from given uuid 
func (conn *SMAPConnection) Prev(uuid string) ([]SMAPData, error){
  url := fmt.Sprintf("%sapi/prev/uuid/%s", conn.Url, uuid)
  return pullData(url)
}

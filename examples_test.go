package gosMAP_test

import (
  "fmt"
  "../goSMAP"
)

// This example gets the first 10 values of a uuid and print out all times and
// values.
func ExampleSMAPConnection_Get(){
  conn,e := gosMAP.Connect("server", "apikey")

  // Get the first 10 values for given uuid
  d, e := conn.Get("uuid", 0, 0, 10)
  if e != nil{
    panic(e)
  }

  // Print out all the time and values
  for _, r := range d[0].Readings{
    fmt.Printf("Time:%s\nValue:%.2f\n", r.Time, r.Value)
  }
}

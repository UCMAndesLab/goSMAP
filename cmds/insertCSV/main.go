package main

import (
    "os"
    "io"
    "fmt"
    "flag"
    "log"
    "encoding/csv"
    "encoding/json"
    "github.com/alexbeltran/gosMAP"
)

var uuid string
var inputFilename string
var apikey string
var server string

func Usage(){
    fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
    flag.PrintDefaults()
}


func init(){
    flag.StringVar(&uuid, "uuid", "", "sMAP UUID")
    flag.StringVar(&server, "server", "", "sMAP Server Address")
    flag.StringVar(&apikey, "api", "", "API Key with permission to data.")
    flag.StringVar(&inputFilename, "in", "in.csv", "CSV file that will be inputed into sMAP.")
    flag.Parse();

    if len(uuid) == 0 || len(apikey) == 0 || len(server) == 0{
        Usage()
        os.Exit(1)
    }
}

const (
    timeBase = 10
    timeBitSize = 64;
    floatBitSize = 64;
)

// Checks to see if an error has occur. All errors passed here are considered fatal.
func check(err error){
    if err != nil{
        log.Fatal(err)
    }
}

const maxReadings = 1000;

func main(){
    fmt.Printf("Reading From File: %s\nInserting Data at: %s\nWith Api Key: %s\n", inputFilename, uuid,apikey)

    // Get path information which is necessary for insertion. This is also a good
    // time to check to see if the uuid exists.
    conn := gosMAP.Connection{
        Url: server,
        APIkey: apikey,
    }
    tag, err := conn.Tag(uuid)
    check(err)
    sourceName := tag.Path


    // Open CSV File to read from
    f, err := os.Open(inputFilename);
    if err != nil{log.Fatal(err)}

    defer f.Close();

    r := csv.NewReader(f);
    r.Comment = '#'

    // Get smap structure ready.
    d := gosMAP.RawData{
            Uuid: uuid,
            Properties: &tag.Properties,
            Metadata: tag.Metadata,
            Readings: make([][]json.Number, maxReadings),
    }

    counter :=0
    for{
        // Start Reading
        record, err := r.Read()

         // End of file
        if err == io.EOF{break}
        check(err)

        // Put Data into reading structure
        d.Readings[counter] = make([]json.Number, 2, 2)
        d.Readings[counter][0] = json.Number(record[0])
        d.Readings[counter][1] = json.Number(record[1])
        counter++;

        // Max Readings hit, Send data and reset counter
        if counter == maxReadings{
            conn.Post(map[string]gosMAP.RawData{sourceName:d,})
            counter = 0
        }
    }

    // Remove additional data from the slice
    if counter > 0{
        d.Readings = d.Readings[0:(counter-1)]
        conn.Post(map[string]gosMAP.RawData{sourceName:d,})
    }
}

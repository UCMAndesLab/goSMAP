package gosMAP

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func pullData(url string) ([]RawData, error) {
	d := make([]RawData, 0)
	response, err := http.Get(url)
	if err != nil {
		return d, err
	}

	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return d, err
	}

	err = json.Unmarshal(contents, &d)
	if err != nil {
		return d, err
	}

	return d, nil
}

func smap_time(t int) string {
	if t == 0 {
		return ""
	} else {
		// smap measures from microseconds since epoch so times it by 1000.
		return fmt.Sprintf("%d", t*1000)
	}
}

// Converts a time.Time into a json.Number that is readable by an sMAP archiver
func TimeToNumber(t time.Time) json.Number {
	return json.Number(fmt.Sprintf("%d", t.Unix()*1000))
}

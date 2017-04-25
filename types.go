package gosMAP

import (
	"encoding/json"
	"fmt"
	"time"
)

type Row struct {
	Time  int64
	Value float64
}

// MarshalJSON is a custom type says we will be transmitting data in the form of [int float64]
func (r *Row) MarshalJSON() ([]byte, error) {
	arr := []interface{}{r.Time, r.Value}
	return json.Marshal(arr)
}

// UnmarshalJSON converts [int float] to a proper go type.
func (r *Row) UnmarshalJSON(bs []byte) error {
	arr := []interface{}{}
	json.Unmarshal(bs, &arr)

	switch t := arr[0].(type) {
	case int64:
		r.Time = t
	case float64:
		r.Time = int64(t)
	default:
		return fmt.Errorf("Element 0 is of type %T; expected a int64", arr[0])
	}

	v, ok := arr[1].(float64)
	if !ok {
		return fmt.Errorf("Element 1 is of type %T; expected a float64", arr[1])
	}
	r.Value = v
	return nil
}

// RawsMAPData is what is returned by a request from the archiver. To make this
// look cleaner, this value is typically converted to a type SMAPData as a return.
type RawData struct {
	Uuid       string          `json:"uuid"`
	Readings   []Row           `json:"Readings,string"`
	Properties *TagsProperties `json:",omitempty"`
	Path       string          `json:",omitempty"`
	Metadata   *Metadata       `json:",omitempty"`
}

type Data struct {
	Uuid       string          `json:"uuid,omitempty" bson:"uuid,omitempty"`
	Readings   []ReadPair      `json:"Readings,omitempty" bson:"Readings,omitempty"`
	Properties *TagsProperties `json:",omitempty" bson:"Properties,omitempty"`
	Path       string          `json:",omitempty" bson:"Path,omitempty"`
	Metadata   *Metadata       `json:",omitempty" bson:"Metadata,omitempty"`
}

// Each value returned by sMAP is a pair of time and float values.
type ReadPair struct {
	Time  time.Time
	Value float64
}

type TagsProperties struct {
	Timezone      string `json:",omitempty" bson:"TimeZone,omitempty"`
	UnitofMeasure string `json:",omitempty" bson:"UnitOfMeasure,omitempty"`
	ReadingType   string `json:",omitempty" bson:"ReadingType,omitempty"`
}

// This is the bare minimium of what sMAP returns to you as
type Tags struct {
	Uuid       string `json:"uuid"`
	Properties TagsProperties
	Path       string
	Metadata   *Metadata
}

func (d *RawData) String() string {
	b, err := json.MarshalIndent(d, "", "   ")

	// Return nothing
	if err != nil {
		return ""
	}
	return string(b)
}

// For metadata field
type Metadata struct {
	SourceName string   `json:",omitempty" bson:"SourceName,omitempty"`
	Location   Location `json:",omitempty" bson:"Location,omitempty"`
	Haystack   Haystack `json:",omitempty" bson:"Haystack,omitempty"`
	Extra      Extra    `json:",omitempty" bson:"Extra,omitempty"`
	HVAC       *HVAC    `json:",omitempty" bson:"HVAC,omitempty"`
}

type HVAC struct {
	// AHU is an identifer for grouping since some buildings can have multiple AHU
	AHU string
	// VAV is an indentifier for grouping rooms into zones
	VAV string
}

type Extra struct {
	Active bool `json:"Active,string"`
}

type Location struct {
	Building string `json:",omitempty" bson:"Building,omitempty"`
	Room     string `json:",omitempty" bson:"Room,omitempty"`
	Floor    string `json:",omitempty" bson:"Floor,omitempty"`
	City     string `json:",omitempty" bson:"City,omitempty"`
	State    string `json:",omitempty" bson:"State,omitempty"`
	Country  string `json:",omitempty" bson:"Country,omitempty"`
}

type Haystack struct {
	Tags string `json:",omitempty" bson:"Tags,omitempty"`
}

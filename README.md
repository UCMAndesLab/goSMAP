# gosMAP
--
    import "github.com/UCMAndesLab/goSMAP"

This is a go binding for sMAP archiver. It is currently in a very early beta and
is not ready for external use. Functions names, types, structures, and pretty
much everything in here is subject to change.

## Usage

#### func  TimeToNumber

```go
func TimeToNumber(t time.Time) json.Number
```
Converts a time.Time into a json.Number that is readable by an sMAP archiver

#### type Connection

```go
type Connection struct {
	Url    string
	APIkey string
	Mc     *memcache.Client
}
```


#### func  Connect

```go
func Connect(url string, key string) (Connection, error)
```

#### func (*Connection) ConnectMemcache

```go
func (conn *Connection) ConnectMemcache(server string)
```

#### func (*Connection) Delete

```go
func (conn *Connection) Delete(uuid string) error
```

#### func (*Connection) Get

```go
func (conn *Connection) Get(uuid string, starttime int, endtime int, limit int) ([]Data, error)
```
Get data given a uuid.

starttime and endtime are unix times in seconds based off of the epoch. A
starttime of 0 will get data starting from the first entry and a endtime of 0
will have no upper bound. Limit is the number of values to be retrieved. Set to
0 if you do not want a limit.

Although the return is an array of SMAPData, typically there should only be one
value with the given uuid.

#### func (*Connection) Post

```go
func (conn *Connection) Post(data map[string]RawData) error
```

#### func (*Connection) Prev

```go
func (conn *Connection) Prev(uuid string) ([]Data, error)
```
Return the last value from given uuid

#### func (*Connection) Query

```go
func (conn *Connection) Query(q string) ([]RawData, error)
```
Use sMAP querying language

See http://www.cs.berkeley.edu/~stevedh/smap2/archiver.html#archiverquery for
further documentation to retrieve data. The contents will be returned as json
text if success, and on some errors a text file

#### func (*Connection) QueryList

```go
func (conn *Connection) QueryList(q string) ([]string, error)
```

#### func (*Connection) Tag

```go
func (conn *Connection) Tag(uuid string) Tags
```
Tag is similar to Tags, however only a single tag is returned

#### func (*Connection) Tags

```go
func (conn *Connection) Tags(uuid string) []Tags
```
Get all tags associate with uuid

#### func (*Connection) UUIDExists

```go
func (conn *Connection) UUIDExists(uuid string) bool
```

#### type Data

```go
type Data struct {
	Uuid     string
	Readings []ReadPair
}
```


#### type RawData

```go
type RawData struct {
	Uuid       string          `json:"uuid"`
	Readings   [][]json.Number `json:"Readings"`
	Properties TagsProperties
	Metadata   map[string]interface{}
}
```

RawsMAPData is what is returned by a request from the archiver. To make this
look cleaner, this value is typically converted to a type SMAPData as a return.

#### type ReadPair

```go
type ReadPair struct {
	Time  time.Time
	Value float64
}
```

Each value returned by sMAP is a pair of time and float values.

#### type Tags

```go
type Tags struct {
	Uuid       string `json:"uuid"`
	Properties TagsProperties
	Path       string
	Metadata   map[string]interface{}
}
```

This is the bare minimium of what sMAP returns to you as

#### type TagsProperties

```go
type TagsProperties struct {
	Timezone      string
	UnitofMeasure string
	ReadingType   string
}
```

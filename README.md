# MultiFormatTime
MultiFormatTime is a Go library designed to facilitate the handling of the time.Time type. Typically, marshaling and unmarshaling the time.Time type is restricted to the RFC3339 format. However, by using this library, you can handle the time.Time type in a variety of formats. It significantly expands the versatility and usability of time data management in Go applications, making it ideal for projects that require flexibility in time formatting and parsing.

# Example
```go
package main

import (
	"encoding/json"
	"fmt"

	"github.com/KantaHasegawa/multi_format_time"
)

type Example struct {
	ExTime multi_format_time.MultiFormatTime `json:"ex_time"`
}

func main() {
	// DateOnly format
	fmt.Println("`````````````DateOnly format``````````````````")
	dateOnlyJSON := `{"ex_time":"2024-12-25"}`
	var dateOnlyEx Example
	json.Unmarshal([]byte(dateOnlyJSON), &dateOnlyEx)
	fmt.Println("time is ", dateOnlyEx.ExTime.Time)

	var dateOnlyMarshaled []byte
	dateOnlyMarshaled, _ = json.Marshal(dateOnlyEx)
	// Marshals in the format used during unmarshaling
	fmt.Println(string(dateOnlyMarshaled))

	// ANSIC format
	fmt.Println("`````````````ANSIC format``````````````````")
	ANSICJSON := `{"ex_time":"Wed Dec 25 00:00:00 2024"}`
	var ANSICEx Example
	json.Unmarshal([]byte(ANSICJSON), &ANSICEx)
	fmt.Println("time is ", ANSICEx.ExTime.Time)

	var ANSICMarshaled []byte
	ANSICMarshaled, _ = json.Marshal(ANSICEx)
	// Marshals in the format used during unmarshaling
	fmt.Println(string(ANSICMarshaled))
}

```
```
`````````````DateOnly format``````````````````
time is  2024-12-25 00:00:00 +0000 UTC
{"ex_time":"2024-12-25"}
`````````````ANSIC format``````````````````
time is  2024-12-25 00:00:00 +0000 UTC
{"ex_time":"Wed Dec 25 00:00:00 2024"}
```

playground  
https://go.dev/play/p/7zKzbntXxun

# Supported Formats
MultiFormatTime supports all formats defined in the Go 'time' package.
| Formats |                              |
|--------------------|-------------------------------------|
| Layout             | 01/02 03:04:05PM '06 -0700          |
| ANSIC              | Mon Jan _2 15:04:05 2006            |
| UnixDate           | Mon Jan _2 15:04:05 MST 2006        |
| RubyDate           | Mon Jan 02 15:04:05 -0700 2006      |
| RFC822             | 02 Jan 06 15:04 MST                 |
| RFC822Z            | 02 Jan 06 15:04 -0700               |
| RFC850             | Monday, 02-Jan-06 15:04:05 MST      |
| RFC1123            | Mon, 02 Jan 2006 15:04:05 MST       |
| RFC1123Z           | Mon, 02 Jan 2006 15:04:05 -0700     |
| RFC3339            | 2006-01-02T15:04:05Z07:00           |
| RFC3339Nano        | 2006-01-02T15:04:05.999999999Z07:00 |
| Kitchen            | 3:04PM                              |
| Stamp              | Jan _2 15:04:05                     |
| StampMilli         | Jan _2 15:04:05.000                 |
| StampMicro         | Jan _2 15:04:05.000000              |
| StampNano          | Jan _2 15:04:05.000000000           |
| DateTime           | 2006-01-02 15:04:05                 |
| DateOnly           | 2006-01-02                          |
| TimeOnly           | 15:04:05                            |

# Contribute
Welcome!

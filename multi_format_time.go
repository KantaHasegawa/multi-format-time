package multi_format_time

import (
	"encoding/json"
	"strings"
	"time"
)

// NOTE: Formatフィールドを持つことでUnMarshal時にどのフォーマットでパースされたかを保持し、それを元にMarshal時に元のフォーマットで出力する
type MultiFormatTime struct {
	time.Time
	Format string
}

// NOTE: 一致するフォーマットが存在しない場合はこのフォーマットで出力する
const DEFAULT_FORMAT = time.RFC3339

var formats = []string{
	time.Layout,
	time.ANSIC,
	time.RubyDate,
	time.UnixDate,
	time.RFC822Z,
	time.RFC822,
	time.RFC850,
	time.RFC1123Z,
	time.RFC1123,
	time.RFC3339,
	time.RFC3339Nano,
	time.Kitchen,
	time.StampNano,
	time.StampMicro,
	time.StampMilli,
	time.Stamp,
	time.DateTime,
	time.DateOnly,
	time.TimeOnly,
}

func (mt *MultiFormatTime) UnmarshalJSON(b []byte) error {
	s := string(b)
	var err error
	var t time.Time
	for _, format := range formats {
		t, err = time.Parse(`"`+format+`"`, s)
		if err == nil {
			// NOTE: RFC3339とRFC3339Nanoは同じフォーマットでパースできるため、ナノ秒が含まれているかどうかで判定する
			if format == time.RFC3339 {
				if strings.Contains(s, ".") {
					t, err = time.Parse(`"`+time.RFC3339Nano+`"`, s)
					if err != nil {
						return err
					}
					format = time.RFC3339Nano
				}
			}
			mt.Time = t
			mt.Format = format
			return nil
		}
	}

	return err
}

func (mt MultiFormatTime) MarshalJSON() ([]byte, error) {
	if mt.Format != "" {
		return json.Marshal(mt.Time.Format(mt.Format))
	}
	return json.Marshal(mt.Time.Format(DEFAULT_FORMAT))
}

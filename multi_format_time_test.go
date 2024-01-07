package multi_format_time

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMultiFormatTime_UnmarshalJSON(t *testing.T) {
	type wants struct {
		multiFormatTime MultiFormatTime
		isErr           bool
	}

	tests := []struct {
		name  string
		args  string
		wants wants
	}{
		{
			name: "ANSIC format",
			args: `"Mon Jan 01 00:00:00 2024"`,
			wants: wants{
				multiFormatTime: MultiFormatTime{Time: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC), Format: time.ANSIC},
			},
		},
		{
			name: "RubyDate format",
			args: `"Mon Jan 01 00:00:00 +0000 2024"`,
			wants: wants{
				multiFormatTime: MultiFormatTime{Time: time.Date(2024, 1, 1, 0, 0, 0, 0, time.FixedZone("", 0)), Format: time.RubyDate},
			},
		},
		{
			name: "UnixDate format",
			args: `"Mon Jan 01 00:00:00 UTC 2024"`,
			wants: wants{
				multiFormatTime: MultiFormatTime{Time: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC), Format: time.UnixDate},
			},
		},
		{
			name: "RFC822Z format",
			args: `"01 Jan 24 00:00 +0000"`,
			wants: wants{
				multiFormatTime: MultiFormatTime{Time: time.Date(2024, 1, 1, 0, 0, 0, 0, time.FixedZone("", 0)), Format: time.RFC822Z},
			},
		},
		{
			name: "RFC822 format",
			args: `"01 Jan 24 00:00 UTC"`,
			wants: wants{
				multiFormatTime: MultiFormatTime{Time: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC), Format: time.RFC822},
			},
		},
		{
			name: "RFC850 format",
			args: `"Monday, 01-Jan-24 00:00:00 UTC"`,
			wants: wants{
				multiFormatTime: MultiFormatTime{Time: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC), Format: time.RFC850},
			},
		},
		{
			name: "RFC1123Z format",
			args: `"Mon, 01 Jan 2024 00:00:00 +0000"`,
			wants: wants{
				multiFormatTime: MultiFormatTime{Time: time.Date(2024, 1, 1, 0, 0, 0, 0, time.FixedZone("", 0)), Format: time.RFC1123Z},
			},
		},
		{
			name: "RFC1123 format",
			args: `"Mon, 01 Jan 2024 00:00:00 UTC"`,
			wants: wants{
				multiFormatTime: MultiFormatTime{Time: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC), Format: time.RFC1123},
			},
		},
		{
			name: "RFC3339Nano format",
			args: `"2024-01-01T00:00:00.123456789Z"`,
			wants: wants{
				multiFormatTime: MultiFormatTime{Time: time.Date(2024, 1, 1, 0, 0, 0, 123456789, time.UTC), Format: time.RFC3339Nano},
			},
		},
		{
			name: "RFC3339 format",
			args: `"2024-01-01T00:00:00Z"`,
			wants: wants{
				multiFormatTime: MultiFormatTime{Time: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC), Format: time.RFC3339},
			},
		},
		{
			name: "Kitchen format",
			args: `"12:00AM"`,
			wants: wants{
				multiFormatTime: MultiFormatTime{Time: time.Date(0, 1, 1, 0, 0, 0, 0, time.UTC), Format: time.Kitchen},
			},
		},
		{
			name: "Stamp format",
			args: `"Jan 01 00:00:00"`,
			wants: wants{
				multiFormatTime: MultiFormatTime{Time: time.Date(0, 1, 1, 0, 0, 0, 0, time.UTC), Format: time.Stamp},
			},
		},
		{
			name: "StampMilli format",
			args: `"Jan 01 00:00:00.123"`,
			wants: wants{
				multiFormatTime: MultiFormatTime{Time: time.Date(0, 1, 1, 0, 0, 0, 123000000, time.UTC), Format: time.StampMilli},
			},
		},
		{
			name: "StampMicro format",
			args: `"Jan 01 00:00:00.123456"`,
			wants: wants{
				multiFormatTime: MultiFormatTime{Time: time.Date(0, 1, 1, 0, 0, 0, 123456000, time.UTC), Format: time.StampMicro},
			},
		},
		{
			name: "StampNano format",
			args: `"Jan 01 00:00:00.123456789"`,
			wants: wants{
				multiFormatTime: MultiFormatTime{Time: time.Date(0, 1, 1, 0, 0, 0, 123456789, time.UTC), Format: time.StampNano},
			},
		},
		{
			name: "DateTime format",
			args: `"2024-01-01 00:00:00"`,
			wants: wants{
				multiFormatTime: MultiFormatTime{Time: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC), Format: time.DateTime},
			},
		},
		{
			name: "DateOnly format",
			args: `"2024-01-01"`,
			wants: wants{
				multiFormatTime: MultiFormatTime{Time: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC), Format: time.DateOnly},
			},
		},
		{
			name: "TimeOnly format",
			args: `"00:00:00"`,
			wants: wants{
				multiFormatTime: MultiFormatTime{Time: time.Date(0, 1, 1, 0, 0, 0, 0, time.UTC), Format: time.TimeOnly},
			},
		},
		{
			name: "Invalid format",
			args: `"2023:01:01"`,
			wants: wants{
				multiFormatTime: MultiFormatTime{},
				isErr:           true,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var mt MultiFormatTime
			err := json.Unmarshal([]byte(test.args), &mt)
			if test.wants.isErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, test.wants.multiFormatTime, mt)
			}
		})
	}
}

func TestMultiFormatTime_MarshalJSON(t *testing.T) {
	tests := []struct {
		name string
		args MultiFormatTime
		want string
	}{
		{
			name: "ANSIC format",
			args: MultiFormatTime{Time: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC), Format: time.ANSIC},
			want: `"Mon Jan  1 00:00:00 2024"`,
		},
		{
			name: "RubyDate format",
			args: MultiFormatTime{Time: time.Date(2024, 1, 1, 0, 0, 0, 0, time.FixedZone("", 0)), Format: time.RubyDate},
			want: `"Mon Jan 01 00:00:00 +0000 2024"`,
		},
		{
			name: "UnixDate format",
			args: MultiFormatTime{Time: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC), Format: time.UnixDate},
			want: `"Mon Jan  1 00:00:00 UTC 2024"`,
		},
		{
			name: "RFC822Z format",
			args: MultiFormatTime{Time: time.Date(2024, 1, 1, 0, 0, 0, 0, time.FixedZone("", 0)), Format: time.RFC822Z},
			want: `"01 Jan 24 00:00 +0000"`,
		},
		{
			name: "RFC822 format",
			args: MultiFormatTime{Time: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC), Format: time.RFC822},
			want: `"01 Jan 24 00:00 UTC"`,
		},
		{
			name: "RFC850 format",
			args: MultiFormatTime{Time: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC), Format: time.RFC850},
			want: `"Monday, 01-Jan-24 00:00:00 UTC"`,
		},
		{
			name: "RFC1123Z format",
			args: MultiFormatTime{Time: time.Date(2024, 1, 1, 0, 0, 0, 0, time.FixedZone("", 0)), Format: time.RFC1123Z},
			want: `"Mon, 01 Jan 2024 00:00:00 +0000"`,
		},
		{
			name: "RFC1123 format",
			args: MultiFormatTime{Time: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC), Format: time.RFC1123},
			want: `"Mon, 01 Jan 2024 00:00:00 UTC"`,
		},
		{
			name: "RFC3339Nano format",
			args: MultiFormatTime{Time: time.Date(2024, 1, 1, 0, 0, 0, 123456789, time.UTC), Format: time.RFC3339Nano},
			want: `"2024-01-01T00:00:00.123456789Z"`,
		},
		{
			name: "RFC3339 format",
			args: MultiFormatTime{Time: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC), Format: time.RFC3339},
			want: `"2024-01-01T00:00:00Z"`,
		},
		{
			name: "Kitchen format",
			args: MultiFormatTime{Time: time.Date(0, 1, 1, 0, 0, 0, 0, time.UTC), Format: time.Kitchen},
			want: `"12:00AM"`,
		},
		{
			name: "Stamp format",
			args: MultiFormatTime{Time: time.Date(0, 1, 1, 0, 0, 0, 0, time.UTC), Format: time.Stamp},
			want: `"Jan  1 00:00:00"`,
		},
		{
			name: "StampMilli format",
			args: MultiFormatTime{Time: time.Date(0, 1, 1, 0, 0, 0, 123000000, time.UTC), Format: time.StampMilli},
			want: `"Jan  1 00:00:00.123"`,
		},
		{
			name: "StampMicro format",
			args: MultiFormatTime{Time: time.Date(0, 1, 1, 0, 0, 0, 123456000, time.UTC), Format: time.StampMicro},
			want: `"Jan  1 00:00:00.123456"`,
		},
		{
			name: "StampNano format",
			args: MultiFormatTime{Time: time.Date(0, 1, 1, 0, 0, 0, 123456789, time.UTC), Format: time.StampNano},
			want: `"Jan  1 00:00:00.123456789"`,
		},
		{
			name: "DateTime format",
			args: MultiFormatTime{Time: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC), Format: time.DateTime},
			want: `"2024-01-01 00:00:00"`,
		},
		{
			name: "DateOnly format",
			args: MultiFormatTime{Time: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC), Format: time.DateOnly},
			want: `"2024-01-01"`,
		},
		{
			name: "TimeOnly format",
			args: MultiFormatTime{Time: time.Date(0, 1, 1, 0, 0, 0, 0, time.UTC), Format: time.TimeOnly},
			want: `"00:00:00"`,
		},
		{
			name: "Default format",
			args: MultiFormatTime{Time: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)},
			want: `"2024-01-01T00:00:00Z"`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			marshaled, err := json.Marshal(test.args)
			assert.NoError(t, err)
			assert.JSONEq(t, test.want, string(marshaled))
		})
	}
}

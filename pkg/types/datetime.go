package types

import (
	"time"

	"github.com/spf13/cast"
)

const DefaultDateLayout = "2006-01-02 15:04:05.000Z"

func NowDateTime() DateTime {
	return DateTime{t: time.Now()}
}

type DateTime struct {
	t time.Time
}

func (d DateTime) Time() time.Time {
	return d.t
}

func (d DateTime) IsZero() bool {
	return d.Time().IsZero()
}

func (d DateTime) String() string {
	if d.IsZero() {
		return ""
	}
	return d.Time().UTC().Format(DefaultDateLayout)
}

func (d *DateTime) Scan(value any) error {
	switch v := value.(type) {
	case DateTime:
		d.t = v.Time()
	case time.Time:
		d.t = v
	case int:
		d.t = cast.ToTime(v)
	case string:
		if v == "" {
			d.t = time.Time{}
		} else {
			t, err := time.Parse(DefaultDateLayout, v)
			if err != nil {
				// check for other common date layouts
				t = cast.ToTime(v)
			}
			d.t = t
		}
	default:
		str := cast.ToString(v)
		if str == "" {
			d.t = time.Time{}
		} else {
			d.t = cast.ToTime(str)
		}
	}

	return nil
}

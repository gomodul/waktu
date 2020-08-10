package waktu

import (
	"time"
)

// Time struct
type Time struct {
	time.Time
}

// LastDay func
func LastDay() Time {
	return Now().LastDay()
}

// Now func
func Now() Time {
	loc, _ := time.LoadLocation("Asia/Jakarta")
	return Time{
		time.Now().In(loc),
	}
}

// Parse func
func Parse(layout, value string) (Time, error) {
	t, err := time.Parse(layout, value)
	return Time{t}, err
}

// Date func
func Date(year int, month Month, day, hour, min, sec, nsec int, loc *time.Location) Time {
	if loc == nil {
		loc = time.UTC
	}
	return Time{time.Date(year, time.Month(month), day, hour, min, sec, nsec, loc)}
}

// AddDate func
func (t Time) AddDate(years int, months int, days int) Time {
	return Time{t.Time.AddDate(years, months, days)}
}

// Add func
func (t Time) Add(d time.Duration) Time {
	return Time{t.Time.Add(d)}
}

// ResetTime func
func (t *Time) ResetTime() Time {
	return Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.UTC).In(t.Location())
}

// SetDate func
func (t *Time) SetDate(date interface{}) Time {
	return Date(t.Year(), t.Month(), date.(int), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), t.Location())
}

// In func
func (t Time) In(loc *time.Location) Time {
	return Time{
		t.Time.In(loc),
	}
}

// LastDay func
func (t Time) LastDay() Time {
	day := 24 * time.Hour
	return t.SetDate(1).AddDate(0, 1, 0).Add(-day - time.Nanosecond)
}

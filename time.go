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
func (t Time) ResetTime() Time {
	return Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.UTC).In(t.Location())
}

// SetDate func
func (t *Time) SetDate(date interface{}) Time {
	return Date(t.Year(), t.Month(), date.(int), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), t.Location())
}

// GetDate ex: 930910
func (t *Time) GetDate(format ...string) string {
	layout := YYMMDD
	if len(format) > 0 {
		layout = format[0]
	}
	return t.Format(layout)
}

// SetHour func
func (t *Time) SetHour(hour interface{}) Time {
	return Date(t.Year(), t.Month(), t.Day(), hour.(int), t.Minute(), t.Second(), t.Nanosecond(), t.Location())
}

// SetMinute func
func (t *Time) SetMinute(minute interface{}) Time {
	return Date(t.Year(), t.Month(), t.Day(), t.Hour(), minute.(int), t.Second(), t.Nanosecond(), t.Location())
}

// SetSecond func
func (t *Time) SetSecond(second interface{}) Time {
	return Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), second.(int), t.Nanosecond(), t.Location())
}

// SetNanosecond func
func (t *Time) SetNanosecond(nanosecond interface{}) Time {
	return Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), nanosecond.(int), t.Location())
}

// In func
func (t Time) In(loc *time.Location) Time {
	return Time{
		t.Time.In(loc),
	}
}

// GetTime 133700
func (t *Time) GetTime(format ...string) string {
	layout := HHMMSS
	if len(format) > 0 {
		layout = format[0]
	}
	return t.Format(layout)
}

// LastDay func
func (t Time) LastDay() Time {
	day := 24 * time.Hour
	return t.SetDate(1).AddDate(0, 1, 0).Add(-day - time.Nanosecond)
}

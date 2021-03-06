package waktu

import (
	"time"
)

// EndOfDay ex: 1993-09-10 23:59:59.999999999 +0700 GMT+7.
func (t Time) EndOfDay() Time {
	return t.StartOfDay().AddDate(0, 0, 1).Add(-time.Nanosecond)
}

// EndOfWeek ex: 1993-09-11 23:59:59.999999999 +0700 GMT+7.
func (t Time) EndOfWeek(startWeek ...Weekday) Time {
	weekday := 6
	if len(startWeek) > 0 {
		weekday += int(startWeek[0])
	}

	return t.AddDate(0, 0, weekday-int(t.Weekday())).EndOfDay()
}

// EndOfMonth ex: 1993-09-30 23:59:59.999999999 +0700 GMT+7.
func (t Time) EndOfMonth() Time {
	return t.StartOfMonth().AddDate(0, 1, 0).Add(-time.Nanosecond)
}

// EndOfYear ex: 1993-12-31 23:59:59.999999999 +0700 GMT+7.
func (t Time) EndOfYear() Time {
	month := 12
	return t.SetMonth(month).EndOfMonth()
}

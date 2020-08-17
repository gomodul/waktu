package waktu

// StartOfDay ex: 1993-09-10 00:00:00 +0700 GMT+7.
func (t Time) StartOfDay() Time {
	return t.ResetTimeLocal()
}

// StartOfWeek ex: 1993-09-05 00:00:00 +0700 GMT+7.
func (t Time) StartOfWeek(startWeek ...Weekday) Time {
	weekday := int(t.Weekday())
	if len(startWeek) > 0 {
		weekday -= int(startWeek[0])
	}

	return t.AddDate(0, 0, -weekday).StartOfDay()
}

// StartOfMonth ex: 1993-09-01 00:00:00 +0700 GMT+7.
func (t Time) StartOfMonth() Time {
	return Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location())
}

// StartOfYear ex: 1993-01-01 00:00:00 +0700 GMT+7.
func (t Time) StartOfYear() Time {
	return Date(t.Year(), 1, 1, 0, 0, 0, 0, t.Location())
}

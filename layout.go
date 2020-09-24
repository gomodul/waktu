package waktu

const (
	// DEFAULT ex: 1993-09-10 13:37:11.000000011 +0700 GMT+7
	DEFAULT = "2006-01-02 15:04:05.999999999 -0700 MST"

	// ANSIC ex: Fri Sep 10 13:37:00 1993
	ANSIC = "Mon Jan _2 15:04:05 2006"

	// UnixDate ex: Fri Sep 10 13:37:00 GMT+7 1993
	UnixDate = "Mon Jan _2 15:04:05 MST 2006"

	// RubyDate ex: Fri Sep 10 13:37:00 +0700 1993
	RubyDate = "Mon Jan 02 15:04:05 -0700 2006"

	// RFC822 ex: 10 Sep 93 13:37 GMT+7
	RFC822 = "02 Jan 06 15:04 MST"

	// RFC822Z ex: 10 Sep 93 13:37 +0700
	RFC822Z = "02 Jan 06 15:04 -0700" // RFC822 with numeric zone

	// RFC850 ex: Friday, 10-Sep-93 13:37:00 GMT+7
	RFC850 = "Monday, 02-Jan-06 15:04:05 MST"

	// RFC1123 ex: Fri, 10 Sep 1993 13:37:00 GMT+7
	RFC1123 = "Mon, 02 Jan 2006 15:04:05 MST"

	// RFC1123Z ex: Fri, 10 Sep 1993 13:37:00 +0700
	RFC1123Z = "Mon, 02 Jan 2006 15:04:05 -0700" // RFC1123 with numeric zone

	// RFC3339 ex: 1993-09-10T13:37:00+07:00
	RFC3339 = "2006-01-02T15:04:05Z07:00"

	// RFC3339Nano ex: 1993-09-10T13:37:00+07:00
	RFC3339Nano = "2006-01-02T15:04:05.999999999Z07:00"

	// Kitchen ex: 13:37AM
	Kitchen = "3:04PM"

	// Stamp ex: Sep 10 13:37:00
	Stamp = "Jan _2 15:04:05"

	// StampMilli ex: Sep 10 13:37:00.000
	StampMilli = "Jan _2 15:04:05.000"

	// StampMicro ex: Sep 10 13:37:00.000000
	StampMicro = "Jan _2 15:04:05.000000"

	// StampNano ex: Sep 10 13:37:00.000000000
	StampNano = "Jan _2 15:04:05.000000000"

	// ISO8601 ex: 1993-09-10T13:37:11Z
	ISO8601 = "2006-01-02T15:04:05.999Z"

	// YYMM ex: 9309
	YYMM = "0601"

	// MMDD ex: 0910
	MMDD = "0102"

	// YYMMDD ex: 930910
	YYMMDD = "060102"

	// MMDDHHMMSS ex: 0910133711
	MMDDHHMMSS = "0102150405"

	// HHMMSS ex: 133700
	HHMMSS = "150405"
)

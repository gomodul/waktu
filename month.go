package waktu

// Month int
type Month int

const (
	// Januari = 1
	Januari Month = 1 + iota

	// Februari = 2
	Februari

	// Maret = 3
	Maret

	// April = 4
	April

	// Mei = 5
	Mei

	// Juni = 6
	Juni

	// Juli = 7
	Juli

	// Agustus = 8
	Agustus

	// September = 9
	September

	// Oktober = 10
	Oktober

	// November = 11
	November

	// Desember = 12
	Desember
)

// months var
var months = [...]string{
	"Januari",
	"Februari",
	"Maret",
	"April",
	"Mei",
	"Juni",
	"Juli",
	"Agustus",
	"September",
	"Oktober",
	"November",
	"Desember",
}

// Month returns the month of the year specified by t.
func (t Time) Month() Month {
	return Month(t.Time.Month())
}

// SetMonth func
func (t *Time) SetMonth(month int) Time {
	return Date(t.Year(), Month(month), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), t.Location())
}

// String func
func (m Month) String() string {
	if Januari <= m && m <= Desember {
		return months[m-1]
	}
	buf := make([]byte, 20)
	n := fmtInt(buf, uint64(m))
	return "%!Month(" + string(buf[n:]) + ")"
}

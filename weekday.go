package waktu

// Weekday int
type Weekday int

const (
	// Senin = 0
	Senin Weekday = iota

	// Selasa = 1
	Selasa

	// Rabu = 2
	Rabu

	// Kamis = 3
	Kamis

	// Jumat = 4
	Jumat

	// Sabtu = 5
	Sabtu

	// Minggu = 6
	Minggu
)

var days = [...]string{
	"Senin",
	"Selasa",
	"Rabu",
	"Kamis",
	"Jumat",
	"Sabtu",
	"Minggu",
}

// Weekday returns the day of the week specified by t.
func (t Time) Weekday() Weekday {
	return Weekday(t.Time.Weekday())
}

// String returns the English name of the day ("Senin", "Selasa", ...).
func (d Weekday) String() string {
	if Minggu <= d && d <= Sabtu {
		return days[d]
	}
	buf := make([]byte, 20)
	n := fmtInt(buf, uint64(d))
	return "%!Weekday(" + string(buf[n:]) + ")"
}

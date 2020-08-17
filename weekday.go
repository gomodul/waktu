package waktu

// Weekday int.
type Weekday int

const (
	// Minggu = 0
	Minggu Weekday = iota

	// Senin = 1
	Senin

	// Selasa = 2
	Selasa

	// Rabu = 3
	Rabu

	// Kamis = 4
	Kamis

	// Jumat = 5
	Jumat

	// Sabtu = 6
	Sabtu
)

var days = [...]string{ //nolint:gochecknoglobals
	"Minggu",
	"Senin",
	"Selasa",
	"Rabu",
	"Kamis",
	"Jumat",
	"Sabtu",
}

// Weekday returns the day of the week specified by t.
func (t Time) Weekday() Weekday {
	return Weekday(t.Time.Weekday())
}

// String returns the English name of the day ("Minggu", "Senin", ...).
func (d Weekday) String() string {
	if Minggu <= d && d <= Sabtu {
		return days[d]
	}

	buf := make([]byte, 20)
	n := fmtInt(buf, uint64(d))

	return "%!Weekday(" + string(buf[n:]) + ")"
}

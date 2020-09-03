package waktu_test

import (
	"testing"

	. "github.com/gomodul/waktu"
)

func TestTime_Format_DEFAULT(t *testing.T) {
	if now.Format(DEFAULT) != now.Time.Format(DEFAULT) {
		t.Fatal("invalid result")
	}
}

func TestTime_Format_ANSIC(t *testing.T) {
	if now.Format(ANSIC) != now.Time.Format(ANSIC) {
		t.Fatal("invalid result")
	}
}

func TestTime_Format_UnixDate(t *testing.T) {
	if now.Format(UnixDate) != now.Time.Format(UnixDate) {
		t.Fatal("invalid result")
	}
}

func TestTime_Format_RubyDate(t *testing.T) {
	if now.Format(RubyDate) != now.Time.Format(RubyDate) {
		t.Fatal("invalid result")
	}
}

func TestTime_Format_RFC822(t *testing.T) {
	if now.Format(RFC822) != now.Time.Format(RFC822) {
		t.Fatal("invalid result")
	}
}

func TestTime_Format_RFC822Z(t *testing.T) {
	if now.Format(RFC822Z) != now.Time.Format(RFC822Z) {
		t.Fatal("invalid result")
	}
}

func TestTime_Format_RFC850(t *testing.T) {
	if now.Format(RFC850) != now.Time.Format(RFC850) {
		t.Fatal("invalid result")
	}
}

func TestTime_Format_RFC1123(t *testing.T) {
	if now.Format(RFC1123) != now.Time.Format(RFC1123) {
		t.Fatal("invalid result")
	}
}

func TestTime_Format_RFC1123Z(t *testing.T) {
	if now.Format(RFC1123Z) != now.Time.Format(RFC1123Z) {
		t.Fatal("invalid result")
	}
}

func TestTime_Format_RFC3339(t *testing.T) {
	if now.Format(RFC3339) != now.Time.Format(RFC3339) {
		t.Fatal("invalid result")
	}
}

func TestTime_Format_RFC3339Nano(t *testing.T) {
	if now.Format(RFC3339Nano) != now.Time.Format(RFC3339Nano) {
		t.Fatal("invalid result")
	}
}

func TestTime_Format_Kitchen(t *testing.T) {
	if now.Format(Kitchen) != now.Time.Format(Kitchen) {
		t.Fatal("invalid result")
	}
}

func TestTime_Format_Stamp(t *testing.T) {
	if now.Format(Stamp) != now.Time.Format(Stamp) {
		t.Fatal("invalid result")
	}
}

func TestTime_Format_StampMilli(t *testing.T) {
	if now.Format(StampMilli) != now.Time.Format(StampMilli) {
		t.Fatal("invalid result")
	}
}

func TestTime_Format_StampMicro(t *testing.T) {
	if now.Format(StampMicro) != now.Time.Format(StampMicro) {
		t.Fatal("invalid result")
	}
}

func TestTime_Format_StampNano(t *testing.T) {
	if now.Format(StampNano) != now.Time.Format(StampNano) {
		t.Fatal("invalid result")
	}
}

func TestTime_Format_ISO8601(t *testing.T) {
	if now.Format(ISO8601) != now.Time.Format(ISO8601) {
		t.Fatal("invalid result")
	}
}

func TestTime_Format_YYMM(t *testing.T) {
	if now.Format(YYMM) != now.Time.Format(YYMM) {
		t.Fatal("invalid result")
	}
}

func TestTime_Format_YYMMDD(t *testing.T) {
	if now.Format(YYMMDD) != now.Time.Format(YYMMDD) {
		t.Fatal("invalid result")
	}
}

func TestTime_Format_YMMDDHHMMSS(t *testing.T) {
	if now.Format(MMDDHHMMSS) != now.Time.Format(MMDDHHMMSS) {
		t.Fatal("invalid result")
	}
}

func TestTime_Format_HHMMSS(t *testing.T) {
	if now.Format(HHMMSS) != now.Time.Format(HHMMSS) {
		t.Fatal("invalid result")
	}
}

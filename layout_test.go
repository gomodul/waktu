package waktu

import (
	"fmt"
	"testing"
)

var now = Date(1993, 9, 10, 11, 37, 0, 0, GetLocationByUTCOffset(7))

func TestTime_Format_ANSIC(t *testing.T) {
	if now.Format(ANSIC) != "Fri Sep 10 11:37:00 1993" {
		t.Fatal("invalid result")
	}
}

func TestTime_Format_UnixDate(t *testing.T) {
	if now.Format(UnixDate) != "Fri Sep 10 11:37:00 GMT+7 1993" {
		t.Fatal("invalid result")
	}
}

func TestTime_Format_RubyDate(t *testing.T) {
	fmt.Println(now.Format(RubyDate))
}

func TestTime_Format_RFC822(t *testing.T) {
	fmt.Println(now.Format(RFC822))
}

func TestTime_Format_RFC822Z(t *testing.T) {
	fmt.Println(now.Format(RFC822Z))
}

func TestTime_Format_RFC850(t *testing.T) {
	fmt.Println(now.Format(RFC850))
}

func TestTime_Format_RFC1123(t *testing.T) {
	fmt.Println(now.Format(RFC1123))
}

func TestTime_Format_RFC1123Z(t *testing.T) {
	fmt.Println(now.Format(RFC1123Z))
}

func TestTime_Format_RFC3339(t *testing.T) {
	fmt.Println(now.Format(RFC3339))
}

func TestTime_Format_RFC3339Nano(t *testing.T) {
	fmt.Println(now.Format(RFC3339Nano))
}

func TestTime_Format_Kitchen(t *testing.T) {
	fmt.Println(now.Format(Kitchen))
}

func TestTime_Format_Stamp(t *testing.T) {
	fmt.Println(now.Format(Stamp))
}

func TestTime_Format_StampMilli(t *testing.T) {
	fmt.Println(now.Format(StampMilli))
}

func TestTime_Format_StampMicro(t *testing.T) {
	fmt.Println(now.Format(StampMicro))
}

func TestTime_Format_StampNano(t *testing.T) {
	fmt.Println(now.Format(StampNano))
}

func TestTime_Format_ISO8601(t *testing.T) {
	fmt.Println(now.Format(ISO8601))
}

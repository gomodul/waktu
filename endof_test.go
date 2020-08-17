package waktu_test

import (
	"fmt"
	"testing"

	. "github.com/gomodul/waktu"
)

func TestTime_EndOfDay(t *testing.T) {
	fmt.Println(now.EndOfDay())
}

func TestTime_EndOfWeek(t *testing.T) {
	endOfWeek := now.EndOfWeek()
	if int(endOfWeek.Weekday()) != int(Sabtu) {
		t.Fatal("invalid value")
	}

	endOfWeek = now.EndOfWeek(Senin)
	if int(endOfWeek.Weekday()) != int(Minggu) {
		t.Fatal("invalid value")
	}
}

func TestTime_EndOfMonth(t *testing.T) {
	fmt.Println(now.EndOfMonth())
}

func TestTime_EndOfYear(t *testing.T) {
	fmt.Println(now.EndOfYear())
}

package waktu

import (
	"fmt"
	"testing"
)

func TestNow(t *testing.T) {
	fmt.Println(now)
}

func TestTime_LastDay(t *testing.T) {
	fmt.Println(now.LastDay())
}

func TestMonth_LastDay(t *testing.T) {
	fmt.Println(now.SetMonth(9).LastDay())
}

func TestLastDay(t *testing.T) {
	fmt.Println(LastDay())
}

func TestDate(t *testing.T) {
	now := Date(2020, Now().Month(), 1, 0, 0, 0, 0, nil)
	fmt.Println(now)
}

func TestTime_GetDate(t *testing.T) {
	if now.GetDate() != now.Format(YYMMDD) {
		t.Fatal("invalid value")
	}
}

func TestTime_GetTime(t *testing.T) {
	if now.GetTime() != now.Format(HHMMSS) {
		t.Fatal("invalid value")
	}
}

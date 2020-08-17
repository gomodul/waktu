package waktu

import (
	"fmt"
	"testing"
	"time"
)

func TestNow(t *testing.T) {
	fmt.Println(now)
}

func TestParse(t *testing.T) {
	value := "9309"
	resParse, errParse := Parse(YYMM, value)
	if errParse != nil {
		t.Fatal(errParse)
	}

	resTimeParse, errTimeParse := time.Parse(YYMM, value)
	if errTimeParse != nil {
		t.Fatal(errParse)
	}

	if resParse.String() != resTimeParse.String() {
		t.Fatal("invalid value")
	}
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

func TestTime_GetDateUTC(t *testing.T) {
	if now.GetDateUTC(time.RFC3339Nano) != now.In(time.UTC).Format(time.RFC3339Nano) {
		t.Fatal("invalid value")
	}
}

func TestTime_GetTime(t *testing.T) {
	if now.GetTime() != now.Format(HHMMSS) {
		t.Fatal("invalid value")
	}
}

func TestTime_ResetTime(t *testing.T) {
	fmt.Println(now.ResetTime())
}

func TestTime_ResetTimeLocal(t *testing.T) {
	fmt.Println(now.ResetTimeLocal())
}

func TestTime_ResetTimeSetDate(t *testing.T) {
	fmt.Println(now.ResetTime().SetDate(2))
}

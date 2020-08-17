package waktu

import (
	"fmt"
	"testing"
)

func TestTime_StartOfDay(t *testing.T) {
	fmt.Println(now.StartOfDay())
}

func TestTime_StartOfWeek(t *testing.T) {
	startOfWeek := now.StartOfWeek()
	if int(startOfWeek.Weekday()) != int(Minggu) {
		t.Fatal("invalid value")
	}
}

func TestTime_StartOfMonth(t *testing.T) {
	fmt.Println(now.StartOfMonth())
}

func TestTime_StartOfYear(t *testing.T) {
	fmt.Println(now.StartOfYear())
}

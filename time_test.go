package waktu

import (
	"fmt"
	"testing"
)

func TestNow(t *testing.T) {
	now := Now()
	fmt.Println(now)
}

func TestTime_LastDay(t *testing.T) {
	now := Now()
	fmt.Println(now.LastDay())
}

func TestLastDay(t *testing.T) {
	fmt.Println(LastDay())
}

func TestDate(t *testing.T) {
	now := Date(2020, Now().Month(), 1, 0, 0, 0, 0, nil)
	fmt.Println(now)
}

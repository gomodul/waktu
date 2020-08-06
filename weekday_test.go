package waktu

import (
	"fmt"
	"testing"
)

func TestWeekday_String(t *testing.T) {
	now := Now()
	fmt.Println(now.Weekday().String())
}

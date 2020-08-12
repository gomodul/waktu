package waktu

import (
	"fmt"
	"testing"
)

func TestWeekday_String(t *testing.T) {
	fmt.Println(now.Weekday().String())
}

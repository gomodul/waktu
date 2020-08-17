package waktu_test

import (
	"fmt"
	"testing"
)

func TestWeekday_String(t *testing.T) {
	fmt.Println(now.Weekday().String())
}

package waktu

import (
	"fmt"
	"testing"
)

func TestMonth_String(t *testing.T) {
	now := Now()
	fmt.Println(now.Month().String())
}

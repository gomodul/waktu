package waktu

import (
	"fmt"
	"testing"
)

func TestTime_Format(t *testing.T) {
	now := Now()
	fmt.Println(now.Format(RFC822))
}

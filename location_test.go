package waktu

import (
	"testing"
	"time"
)

func TestGetLocationByUTCOffset(t *testing.T) {
	if GetLocationByUTCOffset(0) != time.UTC {
		t.Fatal("invalid value")
	}
}

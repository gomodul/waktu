package waktu_test

import (
	"testing"
	"time"

	. "github.com/gomodul/waktu"
)

func TestGetLocationByUTCOffset_Offset0(t *testing.T) {
	if GetLocationByUTCOffset(0) != time.UTC {
		t.Fatal("invalid value")
	}
}

func TestGetLocationByUTCOffset_Offset1(t *testing.T) {
	if GetLocationByUTCOffset(1) == time.UTC {
		t.Fatal("invalid value")
	}
}

package waktu_test

import (
	. "github.com/gomodul/waktu"
)

var now = Date(1993, 9, 10, 13, 37, 11, 11, GetLocationByUTCOffset(7)) //nolint:gochecknoglobals,nolintlint

package waktu

import (
	"fmt"
	"time"
)

// GetLocationByUTCOffset func(offset int) *time.Location.
func GetLocationByUTCOffset(offset int) *time.Location {
	if offset == 0 {
		return time.UTC
	}

	name := fmt.Sprintf("GMT%v", offset)
	if offset > 0 {
		name = fmt.Sprintf("GMT+%v", offset)
	}

	const sixty = 60

	return time.FixedZone(name, offset*sixty*sixty)
}

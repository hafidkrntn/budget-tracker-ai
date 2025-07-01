package time

import (
	"backend-go/constants"
	"time"
)

func Times() string {
	now := time.Now()
	format := constants.DateFromStd
	time := now.Format(format)

	return time
}

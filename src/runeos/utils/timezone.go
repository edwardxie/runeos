package utils

import (
	"time"
)

func Time(t time.Time, loc, layout string) string {
	l, _ := time.LoadLocation(loc)
	return time.Now().In(l).Format(layout)
}

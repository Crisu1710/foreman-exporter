package parser

import (
	"time"
)

func ConvertTime(lastReport string) (float64, error) {
	layout := "2006-01-02 15:04:05 UTC"
	t, err := time.Parse(layout, lastReport)
	toFloat64 := float64(t.Unix())

	return toFloat64, err
}

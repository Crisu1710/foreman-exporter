package parser

import (
	"fmt"
	"time"
)

func ConvertTime(lastReport string) float64 {
	layout := "2006-01-02 15:04:05 UTC"
	t, err := time.Parse(layout, lastReport)
	if err != nil {
		fmt.Println(err)
	}
	toFloat64 := float64(t.Unix())

	return toFloat64
}

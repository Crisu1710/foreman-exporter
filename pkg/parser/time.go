package parser

import (
	"fmt"
	"time"
)

func ConvertTime(newtime string) float64 {
	layout := "2006-01-02 15:04:05 UTC"
	t, err := time.Parse(layout, newtime)
	if err != nil {
		fmt.Println(err)
	}
	tofloat64 := float64(t.Unix())

	return tofloat64
}

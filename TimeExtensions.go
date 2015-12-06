package main

import (
	"fmt"
	"time"
)

func formatDate(t time.Time) string {
	str := fmt.Sprintf("%d-%02d-%02d", t.Year(), t.Month(), t.Day())
	//str := fmt.Sprintf("%d", 10)
	return str
}

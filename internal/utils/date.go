package utils

import (
	"time"
)

var defaultDateLayout = "02/01/2006"

// StringToDate parses the string into date format data
func StringToDate(dateString string) (time.Time, error) {
	return time.Parse(defaultDateLayout, dateString)
}

// DateToString parses the date data into a string
func DateToString(date time.Time) string {
	return date.Format(defaultDateLayout)
}

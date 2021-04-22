// Package timefmt can translate
package timefmt

import (
	"fmt"
	"regexp"
	"time"
)

type timeReplacement struct {
	Input    string
	GoFormat string
}

const (
	defaultTimeFormat  = "HH:mm:ss"
	filenameTimeFormat = "YYYY-MM-DD HH:mm:ss"
)

var timeReplacements = []timeReplacement{
	{"MMMM", "January"},
	{"MMM", "Jan"},
	{"MM", "01"},
	{"M", "1"},
	{"dddd", "Monday"},
	{"ddd", "Mon"},
	{"YYYY", "2006"},
	{"YY", "06"},
	{"DD", "02"},
	{"D", "2"},
	{"HH", "15"},
	{"hha", "03PM"},
	{"hh", "03"},
	{"ha", "3PM"},
	{"h", "3"},
	{"mm A", "04 PM"},
	{"mm", "04"},
	{"m", "4"},
	{"ssa", "05PM"},
	{"ss", "05"},
	{"s", "5"},
	{"ZZ", "-07:00"},
	{"Z", "MST"},
}

// Format returns time in passed timezone and timeFormat.
// Available timeFormat check here https://github.com/vjeantet/jodaTime#format
// It panics when can not find timezone in default timezones of Go
func Format(goTime time.Time, timeFormat, timezone string) string {
	location, err := time.LoadLocation(timezone)
	if err != nil {
		panic(err)
	}
	timeInLocation := goTime.In(location)

	formatStr := timeFormat
	for _, v := range timeReplacements {
		var pattern = regexp.MustCompile(v.Input)
		formatStr = pattern.ReplaceAllString(formatStr, v.GoFormat)
	}
	return timeInLocation.Format(formatStr)
}

// FormatWithTime is the same as Format but appends " HH:mm:ss" to passed dateFormat
func FormatWithTime(goTime time.Time, dateFormat, timezone string) string {
	return Format(goTime, fmt.Sprintf("%s %s", dateFormat, defaultTimeFormat), timezone)
}

// FormatFilenameWithTime changes time to passed timezone with "filenameTimeFormat" foramt
func FormatFilenameWithTime(goTime time.Time, timezone string) string {
	return Format(goTime, filenameTimeFormat, timezone)
}

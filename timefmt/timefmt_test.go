package timefmt

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var exampleTime, err = time.Parse(time.RFC3339, "2018-02-03T05:07:09Z")
var formatCases = map[string]string{
	"YYYY": "2018",
	"YY":   "18",
	"MMMM": "February",
	"MMM":  "Feb",
	"MM":   "02",
	"M":    "2",
	"DD":   "03",
	"D":    "3",
	"dddd": "Saturday",
	"ddd":  "Sat",
	"HH":   "05",
	"hha":  "05AM",
	"hh":   "05",
	"ha":   "5AM",
	"mm":   "07",
	"m":    "7",
	"ss":   "09",
	"s":    "9",
	"ZZ":   "+00:00",
	"Z":    "UTC",
	"YYYY-DD-MM HH:mm:ss": "2018-03-02 05:07:09",
}

func TestUTCFormat(t *testing.T) {
	for format, expected := range formatCases {
		actual := Format(exampleTime, format, "UTC")
		assert.Equal(t, expected, actual, format)
	}
}

func TestFormatBadTimezone(t *testing.T) {
	assert.Panics(t, func() {
		Format(exampleTime, "YYYY-DD-MM HH:mm:ss", "BAD TIMEZONE")
	})
}

func TestUTCFormatWithTime(t *testing.T) {
	actual := FormatWithTime(exampleTime, "YY/D/M", "UTC")
	assert.Equal(t, "18/3/2 05:07:09", actual)
}

func TestZoneFormat(t *testing.T) {
	actual := Format(exampleTime, "YYYY-DD-MMTHH:mm:ss", "America/Noronha")
	assert.Equal(t, "2018-03-02T03:07:09", actual)
}

func TestFormatFilenameWithTime(t *testing.T) {
	actual := FormatFilenameWithTime(exampleTime, "America/Noronha")
	assert.Equal(t, "2018-02-03 03:07:09", actual)
}

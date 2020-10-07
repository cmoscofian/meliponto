package util

import (
	"errors"
	"fmt"
	"log"
	"math"
	"time"

	"github.com/cmoscofian/meliponto/src/cli/context"
	"github.com/cmoscofian/meliponto/src/shared/util/constant"
)

var holidays []time.Time
var today time.Time

func init() {
	ctx := context.New()
	for _, d := range ctx.Holidays {
		day, err := time.Parse(constant.DateInputLayout, d)
		if err != nil {
			log.Fatalln(err)
		}
		holidays = append(holidays, day)
	}

	d := 24 * time.Hour
	today = time.Now().Truncate(d)
}

// IsSameDay reports whether day1 and day2 have same
// day, month and year. It does not check for hours, minutes
// and seconds.
func IsSameDay(day1, day2 time.Time) bool {
	y1, m1, d1 := day1.Date()
	y2, m2, d2 := day2.Date()
	return y1 == y2 && m1 == m2 && d1 == d2
}

// IsWeekday reports whether date is a valid weekday
// neither saturday nor sunday.
func IsWeekday(date time.Time) bool {
	return !IsSaturday(date) && !IsSunday(date)
}

// IsSaturday reports whether date is a saturday
func IsSaturday(date time.Time) bool {
	return date.Weekday() == time.Saturday
}

// IsSunday reports whether date is a saturday
func IsSunday(date time.Time) bool {
	return date.Weekday() == time.Sunday
}

// IsHoliday reports whether date is a holiday
// as provided in the config file.
func IsHoliday(date time.Time) bool {
	y1, m1, d1 := date.Date()

	for _, h := range holidays {
		if y2, m2, d2 := h.Date(); d1 == d2 && m1 == m2 && y1 == y2 {
			return true
		}
	}

	return false
}

// IsRegularDay reports whether date is NOT a holiday
// as provided by the config file and is also a weekday
// neither saturday nor sunday.
func IsRegularDay(date time.Time) bool {
	return IsWeekday(date) && !IsHoliday(date)
}

// IsWeekHoliday reports whether date is a holiday
// as provided by the config file and is also a weekday
// neither saturday nor sunday.
func IsWeekHoliday(date time.Time) bool {
	return IsWeekday(date) && IsHoliday(date)
}

// ParseInputDate returns a Time and an error. It will attempt to parse the date
// given that it is in a valid format ('d-m-yy') and return it as a time.Time format.
func ParseInputDate(date string) (time.Time, error) {
	data, err := time.Parse(constant.DateInputLayout, date)

	if err != nil {
		return data, errors.New(constant.InvalidDateError)
	}

	if data.After(today) || data.Equal(today) {
		return data, errors.New(constant.PastDateError)
	}

	return data, nil
}

// ParseInputDateTime returns a Time and an error. It will attempt to parse the date
// and hour given that they are both in valid formats (date: 'd-m-yy', hour: 'HH:mm')
// and return it as a time.Time format.
func ParseInputDateTime(date, hour string) (time.Time, error) {
	layout := fmt.Sprintf("%s %s", constant.DateInputLayout, constant.TimeLayout)
	data, err := time.Parse(layout, fmt.Sprintf("%s %s", date, hour))
	if data.After(time.Now()) {
		return data, errors.New(constant.PastDateTimeError)
	}

	return data, err
}

// RangeBetweenDatesInDays returns start and end as time.Time formats an int
// representing the range size in days and an error.
// It will attempt to parse start and end date given that they are in a valid format
// ('d-m-yy') to time.Time format and calculate the range in days between them.
func RangeBetweenDatesInDays(start, end string) (time.Time, time.Time, int, error) {
	d1, err := ParseInputDate(start)
	if err != nil {
		return d1, time.Now(), -1, err
	}
	d2, err := ParseInputDate(end)
	if err != nil {
		return d1, d2, -1, err
	}
	if d1.After(d2) {
		return d1, d2, -1, errors.New(constant.EndAfterBeginDateError)
	}
	duration := d2.Sub(d1)

	return d1, d2, int(math.Ceil(duration.Hours()/24)) + 1, nil
}

package util

import (
	"errors"
	"fmt"
	"log"
	"math"
	"time"

	"github.com/cmoscofian/meliponto/src/context"
	"github.com/cmoscofian/meliponto/src/util/constants"
)

var holidays []time.Time
var today time.Time

func init() {
	ctx := context.Create()
	for _, d := range ctx.Holidays {
		day, err := time.Parse(constants.DateInputLayout, d)
		if err != nil {
			log.Fatalln(err)
		}
		holidays = append(holidays, day)
	}

	d := 24 * time.Hour
	today = time.Now().Truncate(d)
}

func IsSameDay(day1, day2 time.Time) bool {
	y1, m1, d1 := day1.Date()
	y2, m2, d2 := day2.Date()
	return y1 == y2 && m1 == m2 && d1 == d2
}

func IsWeekday(date time.Time) bool {
	return date.Weekday() != time.Saturday && date.Weekday() != time.Sunday
}

func IsSaturday(date time.Time) bool {
	return date.Weekday() == time.Saturday
}

func IsSunday(date time.Time) bool {
	return date.Weekday() == time.Sunday
}

func IsHoliday(date time.Time) bool {
	y1, m1, d1 := date.Date()

	for _, h := range holidays {
		if y2, m2, d2 := h.Date(); d1 == d2 && m1 == m2 && y1 == y2 {
			return true
		}
	}

	return false
}

func IsRegularDay(date time.Time) bool {
	return IsWeekday(date) && !IsHoliday(date)
}

func IsWeekHoliday(date time.Time) bool {
	return IsWeekday(date) && IsHoliday(date)
}

func ParseFlagDate(date string) (time.Time, error) {
	data, err := time.Parse(constants.DateInputLayout, date)

	if err != nil {
		return data, errors.New(constants.InvalidDateError)
	}

	if data.After(today) || data.Equal(today) {
		return data, errors.New(constants.PastDateError)
	}

	return data, nil
}

func ParseFlagDatetime(date, hour string) (time.Time, error) {
	layout := fmt.Sprintf("%s %s", constants.DateInputLayout, constants.TimeLayout)
	data, err := time.Parse(layout, fmt.Sprintf("%s %s", date, hour))
	if data.After(time.Now()) {
		return data, errors.New(constants.PastDateTimeError)
	}

	return data, err
}

func RangeBetweenDatesInDays(start, end string) (time.Time, time.Time, int, error) {
	d1, err := ParseFlagDate(start)
	if err != nil {
		return d1, time.Now(), -1, err
	}
	d2, err := ParseFlagDate(end)
	if err != nil {
		return d1, d2, -1, err
	}
	if d1.After(d2) {
		return d1, d2, -1, errors.New(constants.EndAfterBeginDateError)
	}
	duration := d2.Sub(d1)

	return d1, d2, int(math.Ceil(duration.Hours()/24)) + 1, nil
}

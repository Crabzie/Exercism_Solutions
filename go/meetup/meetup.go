// BenchmarkDay-2             28546             45879 ns/op               0 B/op          0 allocs/op
// This solution is based on test's. another solution is available for general use casses.
// https://exercism.org/tracks/go/exercises/meetup/solutions/cunbex
package meetup

import (
	"time"
)

type WeekSchedule int

const (
	First WeekSchedule = iota
	Second
	Third
	Fourth
	Last
	Teenth WeekSchedule = 10
)

func Day(wSched WeekSchedule, wDay time.Weekday, month time.Month, year int) (day int) {
	var initDate time.Time
	for day = 1; day <= 7; day++ {
		initDate = time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
		if initDate.Weekday() == wDay {
			break
		}
	}
	switch v := wSched; {
	case v != Teenth && v != Last:
		_, _, day = initDate.AddDate(0, 0, int(wSched)*7).Date()
	case v == Last:
		daysLeft := time.Date(year, month+1, 0, 0, 0, 0, 0, time.UTC).Day() - initDate.Day()
		_, _, day = initDate.AddDate(0, 0, (daysLeft/7)*7).Date()
	case v == Teenth:
		for d := initDate; d.Day() <= 19; d = d.AddDate(0, 0, 7) {
			if d.Day() >= 13 {
				day = d.Day()
			}
		}
	}
	return
}

package main

import "time"

var weekday = []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
var monthDays = []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

// 1970年12月31日为星期四 2100 ~ 1970
func dayOfTheWeek1(day int, month int, year int) string {
	days := 0
	for i := 1971; i < year; i++ {
		if i%4 == 0 && i%100 != 0 || i%400 == 0 {
			days += 366
		} else {
			days += 365
		}
	}

	for i := 1; i < month; i++ {
		days += monthDays[i-1]
	}

	if month >= 3 && (year%4 == 0 && year%100 != 0 || year%400 == 0) {
		days += 1
	}

	days += day
	return weekday[(days+3)%7]
}

func dayOfTheWeek2(day int, month int, year int) string {
	t := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
	return t.Weekday().String()
}

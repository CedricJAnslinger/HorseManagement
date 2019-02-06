package calendar

import (
	"time"
)

// getDaysOfAWeek returns the days of a week ready to render them in the week-template
// @parameter - year: int >> The year number to which week belongs to.
// @parameter - week: int >> The week number we want to get the days from.
// @return - *[]Day >> All days of the week, each wrapped in a Day-struct.
func getDaysOfAWeek(year int, week int) []Day {
	days := make([]Day, 7)
	monday := getFirstDayOfAWeek(year, week)

	// TODO: Check isActive
	for i, _ := range days {
		newDate := monday.AddDate(0, 0, i)
		days[i] = Day{DateDay: newDate.Day(), Month: newDate.Month().String(), IsActive: false}
	}

	return days
}

// getFirstDayOfAWeek returns the first day of a week >> used to generate a full week later
// @param - year: int >> The year of the week we want to get the first day from
// @param - week: int >> The week we want to get the firs day from
// @return - t: time.Time >> The first monday for the given week in the given year
func getFirstDayOfAWeek(year int, week int) time.Time {
	// Start calculating from the middle of the year to reduce calculations
	t := time.Date(year, 7, 1, 0, 0, 0, 0, time.UTC)

	// Get the monday of this date (monday cause week starts on monday
	// Special case needed fro Sunday cause int(wd) for sunday == 0
	if wd := t.Weekday(); wd == time.Sunday {
		t = t.AddDate(0, 0, -6)
	} else {
		t = t.AddDate(0, 0, -int(wd)+1)
	}

	// Subtract amount of days from t to get the right monday
	_, w := t.ISOWeek()
	t = t.AddDate(0, 0, (week-w)*7)

	return t
}
package calendar

import (
	"errors"
	"html/template"
	"strconv"
	"time"
)

// getDaysOfAWeek returns the days of a week ready to render them in the week-template
// @parameter - year: int >> The year number to which week belongs to.
// @parameter - week: int >> The week number we want to get the days from.
// @return - *[]Day >> All days of the week, each wrapped in a Day-struct.
func getDaysOfAWeek(year int, week int) []Day {
	days := make([]Day, 7)
	monday := getFirstDayOfAWeek(year, week)

	for i, _ := range days {
		newDate := monday.AddDate(0, 0, i)
		day := Day{DateDay: newDate.Day(), Month: newDate.Month().String(), Year: newDate.Year(), IsActive: false}
		// Set day active if its today's date
		if time.Now().Day() == day.DateDay && time.Now().Month().String() == day.Month && time.Now().Year() == newDate.Year(){
			day.IsActive = true
		}
		days[i] = day
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

// getCalendarWeekTemplate returns the template and the data to render calendar_week for a given week in a given year
// @param - year: int >> The year of the week we want to get the first day from
// @param - week: int >> The week we want to get the firs day from
// @return - tmpl: *template.Template(Pointer) >> The combined template of layout and calendar_week_template to display the data
// @return - data: interface >> The data needed to display the week
// @return - err: error >> Potential error that could occur
func getCalendarWeekTemplate(year int, week int) (tmpl *template.Template, data interface{}, err error) {
	days := getDaysOfAWeek(year, week)

	// Generate PrevLink
	prevLink := generatePrevLinkWeek(year, week)
	// Generate NextLink
	nextLink := generateNextLinkWeek(year, week)

	// Generate data for this page
	data = CalendarWeekPage{
		PageTitle: "Project Horse-Management | Week Calendar",
		Year: strconv.Itoa(year),
		Month: days[0].Month,
		KWWeek: "KW " + strconv.Itoa(week) + ", " + strconv.Itoa(year),
		Days: days,
		PrevLink: prevLink,
		NextLink: nextLink,
	}

	tmpl, err = template.ParseFiles("website/templates/calendar_week_template.html", "website/templates/layout.html")
	if err != nil {
		return nil,nil, errors.New("Couldn't Parse files, error: " + err.Error())
	}

	return tmpl, data, nil
}

// generatePrevLinkWeek returns the PrevLink for the calendar_week template
// @param - year: int >> The year of the week we want to get the first day from
// @param - week: int >> The week we want to get the firs day from
// @return - string >> The actual link
func generatePrevLinkWeek(year int, week int) string {
	if week == 1 {
		// Decide whether leap-year or not >> Leap-year if time to next 01.01 == 366
		t1 := time.Date(year - 1, 1, 1, 0, 0, 0, 0, time.UTC)
		t2 := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)
		amountOfDays := t2.Sub(t1).Hours() / 24
		if amountOfDays == 366 {
			// Leap year >> 53 weeks
			return "/calendar_week/" + strconv.Itoa(year-1) + "/53"
		}

		return "/calendar_week/" + strconv.Itoa(year-1) + "/52"
	}

	return "/calendar_week/" + strconv.Itoa(year) + "/" + strconv.Itoa(week-1)
}

// generateNextLinkWeek returns the NextLink for the calendar_week template
// @param - year: int >> The year of the week we want to get the first day from
// @param - week: int >> The week we want to get the firs day from
// @return - string >> The actual link
func generateNextLinkWeek(year int, week int) string {
	if week == 52 || week == 53 {
		// Decide whether leap-year or not >> Leap-year if time to next 01.01 == 366
		t1 := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)
		t2 := time.Date(year+1, 1, 1, 0, 0, 0, 0, time.UTC)
		amountOfDays := t2.Sub(t1).Hours() / 24
		if amountOfDays == 366 {
			// Leap year
			if week == 53 {
				return "/calendar_week/" + strconv.Itoa(year+1) + "/1"
			}

			return "/calendar_week/" + strconv.Itoa(year) + "/53"
		} else {
			// No leap year >> no 53.th week exists
			return "/calendar_week/" + strconv.Itoa(year+1) + "/1"
		}
	}

	return "/calendar_week/" + strconv.Itoa(year) + "/" + strconv.Itoa(week+1)
}

// getDaysOfAMonth returns the days of a month ready to render them in the month-template
// @parameter - year: int >> The year number to which month belongs to.
// @parameter - month: int >> The month number we want to get the days from.
// @return - *[]Day >> All days of the month, each wrapped in a Day-struct.
func getDaysOfAMonth(year int, month int) []Day {
	days := make([]Day, 35) // 35 = max amount of days in one calendar view

	newDate:= time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)
	newDate = newDate.AddDate(0, 0, -int(newDate.Weekday()))		// Subtract one too much cause we add 1 in the loop

	for i, _ := range days {
		newDate = newDate.AddDate(0, 0, 1)
		day := Day{DateDay: newDate.Day(), Month: newDate.Month().String(), Year: newDate.Year(), IsActive: false}
		// Set day active if its today's date
		if time.Now().Day() == day.DateDay && time.Now().Month().String() == day.Month  && time.Now().Year() == day.Year{
			day.IsActive = true
		}
		days[i] = day
	}

	return days
}

// getCalendarMonthTemplate returns the template and the data to render calendar_month for a given week in a given year
// @param - year: int >> The year of the month
// @param - month: int >> The month we want to display
// @return - tmpl: *template.Template(Pointer) >> The combined template of layout and calendar_month_template to display the data
// @return - data: interface >> The data needed to display the month
// @return - err: error >> Potential error that could occur
func getCalendarMonthTemplate(year int, month int) (tmpl *template.Template, data interface{}, err error) {
	days := getDaysOfAMonth(year, month)

	// Generate PrevLink
	prevLink := generatePrevLinkMonth(year, month)
	// Generate NextLink
	nextLink :=generateNextLinkMonth(year, month)

	// Generate data for this page
	data = CalendarMonthPage{
		PageTitle: "Project Horse-Management | Month Calendar",
		Year: strconv.Itoa(year),
		Month: days[8].Month,
		Days: days,
		PrevLink: prevLink,
		NextLink: nextLink,
	}

	tmpl, err = template.ParseFiles("website/templates/calendar_month_template.html", "website/templates/layout.html")
	if err != nil {
		return nil,nil, errors.New("Couldn't Parse files, error: " + err.Error())
	}

	return tmpl, data, nil
}

// generatePrevLinkMonth returns the PrevLink for the calendar_month template
// @param - year: int >> The year of the month
// @param - month: int >> The month we want to display
// @return - string >> The actual link
func generatePrevLinkMonth(year int, month int) string {
	if month == 1 {
		return "/calendar_month/" + strconv.Itoa(year-1) + "/12"
	}
	return "/calendar_month/" + strconv.Itoa(year) + "/" + strconv.Itoa(month-1)
}

// generatePrevLinkMonth returns the NextLink for the calendar_month template
// @param - year: int >> The year of the month
// @param - month: int >> The month we want to display
// @return - string >> The actual link
func generateNextLinkMonth(year int, month int) string {
	if month == 12 {
		return "/calendar_month/" + strconv.Itoa(year+1) + "/1"
	}
	return "/calendar_month/" + strconv.Itoa(year) + "/" + strconv.Itoa(month+1)
}
package calendar

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"
)

// WeekDefaultController handles a request to display the week calendar for the current week.
// @parameter - w: http.ResponseWriter >> Interface used by an HTTP handler to construct an HTTP response.
// @parameter - r: http.Request(Pointer) >> Received HTTP request.
func WeekDefaultController(w http.ResponseWriter, r *http.Request) {
	// Simply generate data for current date and forward it to WeekController
	today := time.Now()
	yearToday, weekToday := today.ISOWeek()
	WeekController(w, r, weekToday, yearToday)
}

// WeekController handles a request to display the week calendar for a given week in a given year.
// @parameter - w: http.ResponseWriter >> Interface used by an HTTP handler to construct an HTTP response.
// @parameter - r: http.Request(Pointer) >> Received HTTP request.
// @parameter - week: int >> The week number we want to display.
// @parameter - year: int >> The year number to which week belongs to.
func WeekController(w http.ResponseWriter, r *http.Request, week int, year int) {
	days := getDaysOfAWeek(time.Now().Year(), week)

	// Generate data for this page
	data := CalendarWeekPage{
		PageTitle: "Project Horse-Management | Week Calendar",
		Year: strconv.Itoa(year),
		Month: days[0].Month,
		KWWeek: "KW " + strconv.Itoa(week) + ", " + strconv.Itoa(year),
		Days: days,
	}

	tmpl, err := template.ParseFiles("website/templates/calendar_week_template.html", "website/templates/layout.html")
	if err != nil {
		log.Fatal(err)
	}

	tmpl.ExecuteTemplate(w, "layout", data)
}
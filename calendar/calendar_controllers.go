package calendar

import (
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
	// WeekController(w, r, weekToday, yearToday)
	tmpl, data, err := getCalendarWeekTemplate(yearToday, weekToday)
	if err != nil {
		// TODO: Write to log
		log.Println("Error on WeekDefaultController: ", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	tmpl.ExecuteTemplate(w, "layout", data)
}

// WeekController handles a request to display the week calendar for a given week in a given year.
// @parameter - w: http.ResponseWriter >> Interface used by an HTTP handler to construct an HTTP response.
// @parameter - r: http.Request(Pointer) >> Received HTTP request.
func WeekController(w http.ResponseWriter, r *http.Request) {
	year, err := strconv.Atoi(r.Form.Get("year"))
	week, err := strconv.Atoi(r.Form.Get("week"))
	if err != nil {
		log.Println("Error on WeekController: ", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	tmpl, data, err := getCalendarWeekTemplate(year, week)
	if err != nil {
		log.Println("Error on WeekController: ", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	tmpl.ExecuteTemplate(w, "layout", data)
}

// MonthDefaultController handles a request to display the month calendar for the current month.
// @parameter - w: http.ResponseWriter >> Interface used by an HTTP handler to construct an HTTP response.
// @parameter - r: http.Request(Pointer) >> Received HTTP request.
func MonthDefaultController(w http.ResponseWriter, r *http.Request) {
	// Simply generate data for current date and forward it to WeekController
	today := time.Now()
	year := today.Year()
	month := today.Month()
	// WeekController(w, r, weekToday, yearToday)
	tmpl, data, err := getCalendarMonthTemplate(year, int(month))
	if err != nil {
		log.Println("Error on MonthController: ", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	tmpl.ExecuteTemplate(w, "layout", data)
}

// MonthController handles a request to display the month calendar for a given month in a given year.
// @parameter - w: http.ResponseWriter >> Interface used by an HTTP handler to construct an HTTP response.
// @parameter - r: http.Request(Pointer) >> Received HTTP request.
func MonthController(w http.ResponseWriter, r *http.Request) {
	year, err := strconv.Atoi(r.Form.Get("year"))
	month, err := strconv.Atoi(r.Form.Get("month"))
	if err != nil {
		log.Println("Error on MonthController: ", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	tmpl, data, err := getCalendarMonthTemplate(year, month)
	if err != nil {
		log.Println("Error on MonthController: ", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	tmpl.ExecuteTemplate(w, "layout", data)
}

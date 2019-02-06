package calendar

import (
	"html/template"
	"log"
	"net/http"
)

func WeekController(w http.ResponseWriter, r *http.Request) {
	// Get the data for this week
	data := CalendarWeekPage{
		PageTitle: "Project Horse-Management | Week Calendar",
		Year: "2019",
		Month: "January",
		KWWeek: "KW 6, 2018",
		Days: []Day{
			{5, false},
			{6, false},
			{7, true},
			{8, false},
			{9, false},
			{10, false},
			{11, false},
		},
	}

	tmpl, err := template.ParseFiles("website/templates/calendar_week_template.html", "website/templates/layout.html")
	if err != nil {
		log.Fatal(err)
	}

	tmpl.ExecuteTemplate(w, "layout", data)
}
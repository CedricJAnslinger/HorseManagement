package calendar

type Day struct {
	DateDay 	int	 // The number of this day >> lke 22.
	IsActive 	bool // Whether the day shall be highlighted because its today's day
}

type CalendarWeekPage struct {
	PageTitle 	string	// Title of the page
	Year		string	// The year the selected week is in
	Month 		string 	// The month the selected week is in
	KWWeek		string	// The week of the year
	Days     	[]Day	// All days of this week
}
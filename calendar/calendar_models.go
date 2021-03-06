package calendar

type Day struct {
	DateDay 	int	   // The number of this day >> lke 22.
	Month 		string // The month of this day
	Year		int	   // The year of this day
	IsActive 	bool   // Whether the day shall be highlighted because its today's day
}

type CalendarWeekPage struct {
	PageTitle 	string	// Title of the page
	Year		string	// The year the selected week is in
	Month 		string 	// The month the selected week is in
	KWWeek		string	// The week of the year
	Days     	[]Day	// All days of this week
	PrevLink	string  // Link to the previous week
	NextLink	string  // Link to the next week
}

type CalendarMonthPage struct {
	PageTitle 	string	// Title of the page
	Year		string	// The year the selected month is in
	Month 		string 	// The selected month
	Days     	[]Day	// All days of this week
	PrevLink	string  // Link to the previous week
	NextLink	string  // Link to the next week
}
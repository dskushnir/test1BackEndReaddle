package main

var daysHolidays []DayPublicHoliday

type DayPublicHoliday struct {
	Date string `json:"date"`
	Name string `json:"name"`
}

var daysLongHolidays []DayLongHoliday

type DayLongHoliday struct {
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
	DayCount  int    `json:"dayCount"`
}

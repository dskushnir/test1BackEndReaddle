package main

import (
	"fmt"
	"testing"
	"time"
)

func TestTodayLongHolidayPrintHolidayInfo(t *testing.T) {
	var result string
	today = time.Date(2020, 8, 24, 0, 0, 0, 0, time.Local)
	urlYear := today.Format("2006")
	urlToHoliday := fmt.Sprintf("%s%s%s", urlHoliday, urlYear, country)
	urlToWeekend := fmt.Sprintf("%s%s%s", urlWeekend, urlYear, country)

	result = PrintHolidayInfo(urlToHoliday, urlToWeekend)

	if result != "Today is a Independence Day, August 24, and the weekend will last 3 day: August 22 - August 24" {
		t.Error("Expected : Today is a Independence Day, August 24, and the weekend will last 3 day: August 22 - August 24, got ", result)
	}

}

func TestNextLongHolidayPrintHolidayInfo(t *testing.T) {
	var result string

	today = time.Date(2020, 12, 20, 0, 0, 0, 0, time.Local)
	urlYear := today.Format("2006")
	urlToHoliday := fmt.Sprintf("%s%s%s", urlHoliday, urlYear, country)
	urlToWeekend := fmt.Sprintf("%s%s%s", urlWeekend, urlYear, country)

	result = PrintHolidayInfo(urlToHoliday, urlToWeekend)

	if result != "Next holiday is a (Gregorian and Revised Julian) Christmas, December 25, and the weekend will last 3 day: December 25 - December 27" {
		t.Error("Expected : Next holiday is a (Gregorian and Revised Julian) Christmas, December 25, and the weekend will last 3 day: December 25 - December 27, got ", result)
	}
}
func TestTodayHolidayPrintHolidayInfo(t *testing.T) {
	var result string

	today = time.Date(2020, 3, 8, 0, 0, 0, 0, time.Local)
	urlYear := today.Format("2006")
	urlToHoliday := fmt.Sprintf("%s%s%s", urlHoliday, urlYear, country)
	urlToWeekend := fmt.Sprintf("%s%s%s", urlWeekend, urlYear, country)

	result = PrintHolidayInfo(urlToHoliday, urlToWeekend)

	if result != "Today is a International Women's Day, March 8" {
		t.Error("Expected : Today is a International Women's Day, March 8, got ", result)
	}

}

func TestNextHolidayPrintHolidayInfo(t *testing.T) {
	var result string

	today = time.Date(2020, 3, 7, 0, 0, 0, 0, time.Local)
	urlYear := today.Format("2006")
	urlToHoliday := fmt.Sprintf("%s%s%s", urlHoliday, urlYear, country)
	urlToWeekend := fmt.Sprintf("%s%s%s", urlWeekend, urlYear, country)

	result = PrintHolidayInfo(urlToHoliday, urlToWeekend)

	if result != "Next holiday is a International Women's Day, March 8" {
		t.Error("Expected : Next holiday is a International Women's Day, March 8, got ", result)
	}

}
func TestNotFoundHolidayPrintHolidayInfo(t *testing.T) {
	var result string

	today = time.Date(2020, 12, 31, 0, 0, 0, 0, time.Local)
	urlYear := today.Format("2006")
	urlToHoliday := fmt.Sprintf("%s%s%s", urlHoliday, urlYear, country)
	urlToWeekend := fmt.Sprintf("%s%s%s", urlWeekend, urlYear, country)

	result = PrintHolidayInfo(urlToHoliday, urlToWeekend)

	if result != "There are no more holidays this year" {
		t.Error("Expected : There are no more holidays this year ", result)
	}

}

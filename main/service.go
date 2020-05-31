package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func ConvertDate(time2 time.Time) time.Time {
	todayString := time2.Format("2006-01-02")
	todayDate, _ := time.Parse("2006-01-02", todayString)
	return todayDate
}

func MakeRequest(url string) []byte {
	response, err := http.Get(url)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	result, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	return result
}

func FindDayPublicHoliday(url string) *DayPublicHoliday {

	err := json.Unmarshal(MakeRequest(url), &daysHolidays)
	if err != nil {
		fmt.Println("error", err)
	}
	for _, dayPublicHoliday := range daysHolidays {
		date = dayPublicHoliday.Date
		t, _ := time.Parse("2006-01-02", date)
		if err != nil {
			fmt.Println("error", err)
		}
		if t.Equal(ConvertDate(today)) || t.After(ConvertDate(today)) {
			return &dayPublicHoliday
		}
	}
	return nil
}

func FindLongHoliday(url1 string, url2 string) *DayLongHoliday {
	day := FindDayPublicHoliday(url1)
	if day != nil {
		date = day.Date
		t, _ := time.Parse("2006-01-02", date)
		err := json.Unmarshal(MakeRequest(url2), &daysLongHolidays)
		if err != nil {
			fmt.Println("error", err)
		}
		for _, dayLongHoliday := range daysLongHolidays {
			dateStart = dayLongHoliday.StartDate
			tS, _ := time.Parse("2006-01-02", dateStart)
			dateEnd = dayLongHoliday.EndDate
			tE, _ := time.Parse("2006-01-02", dateEnd)
			if t.Equal(tS) || t.Equal(tE) || (t.After(tS) && t.Before(tE)) {
				return &dayLongHoliday
			}
		}
	}
	return nil

}
func PrintHolidayInfo(url1 string, url2 string) string {
	dayH := FindDayPublicHoliday(url1)
	if dayH != nil {
		date = dayH.Date
		tH, _ := time.Parse("2006-01-02", date)
		dayL := FindLongHoliday(url1, url2)
		stringToday := "Today is a "
		stringNext := "Next holiday is a "
		stringWeekend := "and the weekend will last "
		stringDay := " day"
		stringComa := ", "
		stringDash := " - "
		stringColon := ": "

		if dayL != nil {
			dateStartWeekend = dayL.StartDate
			tStart, _ := time.Parse("2006-01-02", dateStartWeekend)
			dateEndWeekend = dayL.EndDate
			tEnd, _ := time.Parse("2006-01-02", dateEndWeekend)
			if tH.Equal(ConvertDate(today)) {
				return fmt.Sprint(stringToday, dayH.Name, stringComa, tH.Month(), tH.Day(), stringComa, stringWeekend, dayL.DayCount, stringDay, stringColon, tStart.Month(), tStart.Day(), stringDash, tEnd.Month(), tEnd.Day())
			}
			if tH.After(ConvertDate(today)) {
				return fmt.Sprint(stringNext, dayH.Name, stringComa, tH.Month(), tH.Day(), stringComa, stringWeekend, dayL.DayCount, stringDay, stringColon, tStart.Month(), tStart.Day(), stringDash, tEnd.Month(), tEnd.Day())
			}
		} else if tH.Equal(ConvertDate(today)) {
			return fmt.Sprint(stringToday, dayH.Name, stringComa, tH.Month(), tH.Day())
		}
		return fmt.Sprint(stringNext, dayH.Name, stringComa, tH.Month(), tH.Day())
	}

	return fmt.Sprint("There are no more holidays this year")

}

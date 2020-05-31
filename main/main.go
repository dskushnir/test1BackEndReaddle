package main

import (
	"fmt"
	"time"
)

var urlHoliday = "https://date.nager.at/Api/v2/PublicHolidays/"
var country = "/UA"
var urlWeekend = "https://date.nager.at/Api/v2/LongWeekend/"
var today time.Time
var date string
var dateStart string
var dateEnd string
var dateStartWeekend string
var dateEndWeekend string

func main() {
	today = time.Now()
	urlYear := today.Format("2006")
	urlToHoliday := fmt.Sprintf("%s%s%s", urlHoliday, urlYear, country)
	urlToWeekend := fmt.Sprintf("%s%s%s", urlWeekend, urlYear, country)

	fmt.Println(PrintHolidayInfo(urlToHoliday, urlToWeekend))

}

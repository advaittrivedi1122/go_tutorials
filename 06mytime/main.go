package main

import (
	"fmt"
	"strings"
	"time"
)

func myDateFormat(date time.Time, formatString string) string {
	const DD string = "02"
	const MM string = "01"
	const YYYY string = "2006"
	const YY string = "06"
	const HH string = "15"
	const MI string = "04"
	const SS string = "05"
	const MS string = "000000000"
	const ST string = "MST"
	const HALFDAY string = "Mon"
	const FULLDAY string = "Monday"

	formatString = strings.Replace(formatString, "dd", DD, 1)
	formatString = strings.Replace(formatString, "mm", MM, 1)
	formatString = strings.Replace(formatString, "yyyy", YYYY, 1)
	formatString = strings.Replace(formatString, "yy", YY, 1)
	formatString = strings.Replace(formatString, "hh", HH, 1)
	formatString = strings.Replace(formatString, "mm", MI, 1)
	formatString = strings.Replace(formatString, "mi", MI, 1)
	formatString = strings.Replace(formatString, "ss", SS, 1)
	formatString = strings.Replace(formatString, "ms", MS, 1)
	formatString = strings.Replace(formatString, "st", ST, 1)
	formatString = strings.Replace(formatString, "fday", FULLDAY, 1)
	formatString = strings.Replace(formatString, "hday", HALFDAY, 1)
	return date.Format(formatString)
}

func main() {
	fmt.Println("Welcome to Time study of Golang!")

	presentTime := time.Now()
	fmt.Println(presentTime)

	// always use this date values as a standard for formatting
	/*
		STANDARD CONSTANT VALUES TO GET DATE Formatted as we want in Go
		Date : 02 | 2
		Month : 01 | 1
		Year : 2006 | 06
		Day : Monday | Mon
		Hour : 15
		Minute : 04 | 4
		Second : 05 | 5
		MilliSecond : 000000000 (number of zeros decides number of precision for millisecond)
		TimeZone : -070000 (number of zeros decides number of precision for timeZone)
		Month By Name : Jan | January
		Standard Time Name : MST


		Date : 02-01-2006 (dd-mm-yyyy)
		Time : 15:04:05.000000000 (hh:mm:ss:ms)
		TimeZone Difference : -07:00:00 (+05:30:00)
	*/
	fmt.Println(presentTime.Format("2006-1-2 15:04:05.000000000 -0700 MST Mon"))
	fmt.Println(presentTime.Format("2-1-2006 15:04:05.000000000 -07:00:00 MST Monday"))

	// my custom date formatter
	fmt.Println()
	fmt.Println(myDateFormat(presentTime, "dd-mm-yyyy hh:mm:ss.ms"))
	fmt.Println(myDateFormat(presentTime, "hh:mi:ss st"))
	fmt.Println(myDateFormat(presentTime, "dd/mm/yy, hh:mm:ss st hday"))
	fmt.Println(myDateFormat(presentTime, "dd/mm/yy, hh:mm:ss st fday"))

	createdDate := time.Date(2020, time.August, 6, 13, 2, 12, 0, time.UTC)
	fmt.Println(myDateFormat(createdDate, "dd-mm-yyyy fday"))
}

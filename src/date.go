// Returns dates that are 30 days apart
// Date sets will be used
package main

import (
	"fmt"
	"time"
)

// Variables to customize this file
var increment = 30
var startDate = "2018-02-01"

type DatePair struct {
	Start string
	End   string
}

// Here's how to run these
// GetDateGroups(startDate, increment)
// GetRecentDates(increment)

// Returns a slice of pairs of dates from the startDate thru the present
// This function is useful for getting historical fitbit data
func GetDateGroups(startDate string, increment int) []DatePair {
	// Create the variable to save all the date pairs in
	var dates []DatePair

	// Parse the starting date
	t, err := time.Parse("2006-01-02", startDate)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
	// Create a new variable x days away from the start date
	t_new := t.AddDate(0, 0, increment)

	// As long as our start date is prior to today's date, do the following
	// for each start date
	for t.Before(time.Now().UTC()) {
		// Append dates to the dates array
		var datePair DatePair
		datePair.Start = t.Format("2006-01-02")
		datePair.End = t_new.Format("2006-01-02")
		dates = append(dates, datePair)

		// set the end date as the start date, and create a new end date
		t = t_new
		t_new = t_new.AddDate(0, 0, increment)

		// print the start date and the day x days in the future
		// fmt.Printf("%v --> %v\n", datePair.Start, datePair.End)
	}

	// Return a slice of dates
	return dates
}

// This function returns a group of dates prior to todays date
// This function will be used to repeatedly get recent fitbit data
func GetRecentDates(increment int) DatePair {
	t := time.Now()
	t_new := t.AddDate(0, 0, -increment)
	fmt.Printf("%v --> %v\n", t_new.Format("2006-01-02"), t.Format("2006-01-02"))

	var datePair DatePair
	datePair.Start = t.Format("2006-01-02")
	datePair.End = t_new.Format("2006-01-02")

	return datePair
}

package main

import (
	"fmt"
	"time"
)

func main() {
	// Program header
	fmt.Println("Time Package")

	// Get current local time
	currentTime := time.Now()
	fmt.Println("Now():", currentTime)

	// Format date using a layout string (month-day-year)
	// layout: "01-02-2006" -> month-day-year
	fmt.Println("Date (MM-DD-YYYY):")
	layoutStr := "01-02-2006"
	formattedDate := currentTime.Format(layoutStr)
	fmt.Println(formattedDate)

	// Format time using 12-hour clock with AM/PM
	fmt.Println("Time (12-hour):")
	layoutTime12 := "03:04:05 PM" // 03:04:05 PM represents 12-hour clock with AM/PM
	formattedTime12 := currentTime.Format(layoutTime12)
	fmt.Println(formattedTime12)

	// Day of the week
	fmt.Println("Day of week:")
	layoutDay := "Monday"
	day := currentTime.Format(layoutDay)
	fmt.Println(day)

	// Converting a date string into a time.Time value
	layoutStr3 := "2006-01-02"
	datestr := "2026-08-24"
	formatted_date, _ := time.Parse(layoutStr3, datestr)
	fmt.Println(formatted_date)

	// Add one day to the current time
	new_day := currentTime.Add(24 * time.Hour)
	// Add 24 minutes to the current time
	new_Min := currentTime.Add(24 * time.Minute)
	// Format the new time values for display
	formatted_newDay := new_day.Format("2006-01-02")
	formatted_newMin := new_Min.Format("03:04:05")
	fmt.Println(formatted_newDay)
	fmt.Println(formatted_newMin)
	fmt.Println("Added 1 day")

	// Individual components
	fmt.Println("Date components:")
	year, month, dayNum := currentTime.Date()
	fmt.Println("Year:", year, "Month:", month, "Day:", dayNum)
	// currentTime.Clock() returns three values (hour, minute, second).
	// Assign them to variables before passing to fmt.Println.
	hour, minute, second := currentTime.Clock()
	fmt.Println("Clock (hour, min, sec):", hour, minute, second)
	fmt.Println("Hour:", currentTime.Hour())
	fmt.Println("Minute:", currentTime.Minute())
	fmt.Println("Second:", currentTime.Second())
}

// ================================================================================
// TIME PACKAGE - FORMATS & IMPORTANT FUNCTIONS
// ================================================================================

// Layout reference (Go uses a reference time to define formats):
// 01   : month (01)
// 02   : day (02)
// 2006 : year (2006)
// 03   : hour (12-hour clock)
// 04   : minute
// 05   : second
// 15   : hour (24-hour clock)
// PM   : AM/PM marker

// Now : (returns the current local time)
// Inputs: none
// Returns: time.Time
// Example:
// t := time.Now()

// Format : (formats a Time according to a layout string)
// Inputs: a time.Time receiver and a layout string; returns a string
// Example:
// s := t.Format("01-02-2006") // "MM-DD-YYYY"

// Parse : (parses a formatted string according to layout and returns time)
// Inputs: a layout string and a value string; returns (time.Time, error)
// Example:
// t, err := time.Parse("01-02-2006", "06-01-2023")

// Unix : (returns seconds since January 1, 1970 UTC)
// Inputs: a time.Time receiver; returns int64 (seconds)
// Example:
// secs := t.Unix()

// Add : (adds a duration to a time and returns the new time)
// Inputs: a time.Duration; returns time.Time
// Example:
// later := t.Add(2 * time.Hour)

// Sub : (returns the duration between two times)
// Inputs: another time.Time; returns time.Duration
// Example:
// dur := t2.Sub(t1)

// Since : (time since a given time)
// Inputs: a time.Time; returns time.Duration
// Example:
// d := time.Since(t0) // equivalent to time.Now().Sub(t0)

// Until : (duration until a given time)
// Inputs: a time.Time; returns time.Duration
// Example:
// d := time.Until(tFuture)

// Sleep : (pauses execution for the specified duration)
// Inputs: a time.Duration; returns nothing
// Example:
// time.Sleep(500 * time.Millisecond)

// ParseDuration : (parses a duration string like "300ms", "1h2m")
// Inputs: a string; returns (time.Duration, error)
// Example:
// d, err := time.ParseDuration("2h45m")

// After : (returns a channel that will receive the current time after the duration)
// Inputs: a time.Duration; returns <-chan time.Time
// Example:
// <-time.After(2 * time.Second)

// NewTimer : (creates a Timer that will send the time on its channel after duration)
// Inputs: a time.Duration; returns *time.Timer
// Example:
// timer := time.NewTimer(2 * time.Second)

// NewTicker : (creates a Ticker that sends the time on its channel at regular intervals)
// Inputs: a time.Duration; returns *time.Ticker
// Example:
// ticker := time.NewTicker(1 * time.Second)

// UTC / Local : (convert time to UTC or local timezone)
// Inputs: a time.Time receiver; returns time.Time
// Example:
// utc := t.UTC()

// Truncate : (rounds down to a multiple of the provided duration)
// Inputs: a time.Duration; returns time.Time
// Example:
// t2 := t.Truncate(time.Minute)

// Format examples recap:
// t.Format("01-02-2006") => "MM-DD-YYYY"
// t.Format("2006-01-02") => "YYYY-MM-DD"
// t.Format("15:04:05") => "HH:MM:SS (24-hour)"
// t.Format("03:04:05 PM") => "HH:MM:SS AM/PM (12-hour)"

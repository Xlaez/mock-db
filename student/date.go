package student

import (
	"fmt"
	"time"
)

type Date struct {
	year  int
	month int
	day   int
}

// create a setter method that sets field values and stores them
func (d *Date) SetYear(year int) error {
	var currentTime time.Time = time.Now()
	currentYear := currentTime.Year()

	if year < 1 || year > currentYear {
		return fmt.Errorf("invalid year")
	}

	d.year = year
	return nil
}

func (d *Date) SetMonth(month int) error {
	if month < 1 || month > 12 {
		return fmt.Errorf("invalid month")
	}
	d.month = month
	return nil
}

func (d *Date) SetDay(day int) error {
	if day < 1 || day > 31 {
		return fmt.Errorf("invalid day")
	}
	d.day = day
	return nil
}

// create a getter method to get values from struct

func (d *Date) Day() int {
	return d.day
}

func (d *Date) Month() int {
	return d.month
}

func (d *Date) Year() int {
	return d.year
}

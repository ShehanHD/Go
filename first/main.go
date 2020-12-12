package main

import "fmt"

type Date struct{
	Day, Month, Year int
}

func main() {
	d := Date{}
	d2 := Date{}

	FirstDate(&d)
	SecondDate(&d2)

	DatesToDays(d, d2)
}

func FirstDate (d *Date){
	for{
		fmt.Scanln(&d.Day)
		fmt.Scanln(&d.Month)
		fmt.Scanln(&d.Year)
		if validDate(*d) { break }
		fmt.Printf("Invalid date, Please try again\n")
	}
}

func SecondDate (d *Date){
	for{
		fmt.Scanln(&d.Day)
		fmt.Scanln(&d.Month)
		fmt.Scanln(&d.Year)
		if validDate(*d) { break }
		fmt.Printf("Invalid date, Please try again\n")
	}
}

func DatesToDays(d Date, d2 Date){

}

func LeapYear(year int) bool{
	if (year % 4 == 0 && year % 100 != 0) || year % 400 == 0 {
		return true
	}
	return false
}

func validDate(d Date) bool{
	var isValid bool = false

	if d.Month>0 &&  d.Month<13 {
		if LeapYear(d.Year) {
			if d.Month == 2 && (d.Month <= 29 || d.Day >= 1) { isValid = true }
		}
		if d.Month == 2 && (d.Month <= 28 || d.Day >= 1) { isValid = true }
		if d.Month % 2 == 0 && (d.Day <= 30 || d.Day >= 1) { isValid = true }
		if d.Month % 2 == 1 && (d.Day <= 31 || d.Day >= 1) { isValid = true }
	}

	return isValid
}
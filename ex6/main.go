package main

import (
	"fmt"
	"strconv"
)

func main() {
	welcomeMsm := welcome("Paulo")
	fmt.Println(welcomeMsm)

	happyMsm := happyBirthday("Rick", 40)
	fmt.Println(happyMsm)

	assugn := assingnTable("Rick", "Morty", "on the left", 27, 23.785555)
	fmt.Println(assugn)
}

func welcome(name string) string {
	return fmt.Sprintf("Welcome to my party, %s", name)
}

func happyBirthday(name string, age int) string {
	return fmt.Sprintf("Happy birthday %s! You are now %d years old!", name, age)
}

func assingnTable(name, seatmate, directTable string, tableNumber int, distanceTable float64) string {
	welcome := welcome(name)

	var numberString string
	if tableNumber < 100 {
		numberString = strconv.Itoa(tableNumber)
		numberString = fmt.Sprintf("0%s", numberString)
	}

	row2 := fmt.Sprintf("You have been assigned to table %s. Your table is %s, exactly %.1f meters from here",
		numberString, directTable, distanceTable)

	row3 := fmt.Sprintf("you will be sitting nect to %s.", seatmate)

	return fmt.Sprintf("%s \n %s \n %s", welcome, row2, row3)
}

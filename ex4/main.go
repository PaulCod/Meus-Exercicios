package main

import (
	"fmt"
	"strings"
)

func main() {
	message := welcomeMessage("Judy")
	fmt.Println(message)

	messagePremium := addBorder(message, len(message))
	fmt.Println(messagePremium)

	messageOld := `
		**************************
		*    BUY NOW, SAVE 10%   *
		**************************
		`

	newMessage := cleanUpMessage(messageOld)
	fmt.Println(newMessage)
}

func welcomeMessage(name string) string {
	upName := strings.ToUpper(name)
	message := "Welcome to the Tech Palace, " + upName
	return message
}

func addBorder(message string, stars int) string {
	var slice []string
	for i := 0; i < stars; i++ {
		slice = append(slice, "*")
	}

	sliceJoin := strings.Join(slice, "")
	messagePremium := fmt.Sprintf("\n %s \n %s \n %s", sliceJoin, message, sliceJoin)
	return messagePremium
}

func cleanUpMessage(oldMessage string) string {
	newMessage := strings.Replace(oldMessage, "*", "", -1)
	return newMessage
}

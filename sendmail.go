package main

import (
	"log"
	"net/smtp"
)

func main() {
	// Set up authentication information.
	auth := smtp.PlainAuth(
		"",
		"bokujunjikubo@gmail.com", // foo@gmail.com
		"lniazcozfreoamyw",
		"smtp.gmail.com",
	)
	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	err := smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		"bokujunjikubo+golang@gmail.com", //foo@gmail.com
		[]string{"junjikubo@outlook.com"},
		[]byte("testmail"),
	)
	if err != nil {
		log.Fatal(err)
	}
}

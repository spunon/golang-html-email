package main

import (
	"fmt"
	"net/smtp"
)

func main() {

	bodyVar := "I'm a variable!"

	body := `<html>
        <h1>This is a heading</h1>
        <p>This is some paragraph text</p>
        <p><i>This is a sincere salutation</i></p>
        <p><strong>P.S. This is a variable: ` + bodyVar + `</strong>`

	send("recipient@email.com", "Some Subject", body)
}

func send(to string, subject string, body string) {
	// Example with gmail smtp
	hostname := "smtp.gmail.com"
	port := "587"
	from := "youremail@gmail.com"
	pass := "yourpassword"

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: " + subject + "\n" +
		"MIME-Version: 1.0\n" +
		"Content-Type: text/html; charset=ISO-8859-1\n\n" +
		body

	err := smtp.SendMail(hostname+":"+port,
		smtp.PlainAuth("", from, pass, hostname),
		from, []string{to}, []byte(msg))

	if err != nil {
		fmt.Printf("SMTP Error: %s", err)
		return
	}

	fmt.Print("Email Sent Successfully to " + to)
}

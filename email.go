package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/mail"

	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap/client"
	"github.com/jung-kurt/gofpdf"
)
​
func main() {
	var password string
	fmt.Println("Enter Your Password:: ")
	fmt.Println("\033[8m") // Hide input
	fmt.Scan(&password)
	fmt.Println("\033[28m") // Show input
	log.Println("Connecting to server...")
​
	// Connect to server
	c, err := client.DialTLS("imap.gmail.com:993", nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected")
​
	// Don't forget to logout
	defer c.Logout()
​
	// Login
​
	if err := c.Login("vidyakailasamp@gmail.com", "Vidyakpa"); err != nil {
		log.Fatal(err)
	}
	log.Println("Logged in")
​
	// Select INBOX
	mbox, err := c.Select("INBOX", false)
	if err != nil {
		log.Fatal(err)
	}
​
	// Get the last message
	if mbox.Messages == 0 {
		log.Fatal("No message in mailbox")
	}
	seqset := new(imap.SeqSet)
	seqset.AddRange(mbox.Messages, mbox.Messages)
​
	// Get the whole message body
	section := &imap.BodySectionName{}
	items := []imap.FetchItem{section.FetchItem()}
​
	messages := make(chan *imap.Message, 1)
	done := make(chan error, 1)
	go func() {
		done <- c.Fetch(seqset, items, messages)
	}()
​
	log.Println("Last message:")
	msg := <-messages
	r := msg.GetBody(section)
	if r == nil {
		log.Fatal("Server didn't returned message body")
	}
​
	if err := <-done; err != nil {
		log.Fatal(err)
	}
​
	m, err := mail.ReadMessage(r)
	if err != nil {
		log.Fatal(err)
	}
​
	header := m.Header
	date := header.Get("Date")
	from := header.Get("From")
	to := header.Get("To")
	sub := header.Get("Subject")
	log.Println("Date:", header.Get("Date"))
	log.Println("From:", header.Get("From"))
	log.Println("To:", header.Get("To"))
	log.Println("Subject:", header.Get("Subject"))
	// log.Println("Body:", header.Get("Body"))
	// log.Println("Attachments:", header.Get("attachment"))
​
	body, err := ioutil.ReadAll(m.Body)
	if err != nil {
		log.Fatal(err)
	}
​
	log.Println(string(body)[:1000])
​
	// Generating PDF.
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	// pdf.Cell(100, 10, date)
	// pdf.Cell(100, 10, from)
	// pdf.Cell(100, 10, to)
	// pdf.Cell(100, 10, sub)
	pdf.CellFormat(100, 10, date, "0", 1, "CM", false, 0, "L")
	pdf.CellFormat(100, 10, from, "0", 1, "CM", false, 0, "L")
	pdf.CellFormat(100, 10, to, "0", 1, "CM", false, 0, "L")
	pdf.CellFormat(100, 10, sub, "0", 1, "CM", false, 0, "L")
	pdf.CellFormat(100, 10, string(body), "0", 1, "CM", false, 0, "L")
​
	pdf.OutputFileAndClose("email.pdf")
	log.Println("Done!")
​
}
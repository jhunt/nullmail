package main

import (
	"bytes"
	"fmt"
	"net"
	"net/mail"
	"os"

	"github.com/mhale/smtpd"
)

func handler(origin net.Addr, from string, to []string, data []byte) {
	fmt.Printf("...\n")
	msg, _ := mail.ReadMessage(bytes.NewReader(data))
	subject := msg.Header.Get("Subject")
	fmt.Printf(">> received mail from %s for %s with subject: %s\n", from, to[0], subject)
	fmt.Printf("%s\n\n", string(data))
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "2525"
	}
	fmt.Printf("listening for inbound SMTP on :%s\n", port)
	smtpd.ListenAndServe(fmt.Sprintf(":%s", port), handler, "nullmail", "")
}

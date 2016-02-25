package web

import (
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"net/mail"
	"net/smtp"
)

// SSL/TLS Email Example, doesn't work ....
func SendEmail() {
	from := mail.Address{"John Lau", "soulmanjohn@163.com"}
	to := mail.Address{"John Lau", "soulmanjohn@163.com"}
	subj := "This is the email subject"
	body := "This is an example body.\n With two lines."

	// Setup headers
	headers := make(map[string]string)
	headers["From"] = from.String()
	headers["To"] = to.String()
	headers["Subject"] = subj

	// Setup message
	message := ""
	//	for k, v := range headers {
	//		message += fmt.Sprintf("%s: %s\r\n", k, v)
	//	}
	message += "\r\n" + body

	// Connect to the SMTP Server
	servername := "smtp.163.com:25"

	host, _, _ := net.SplitHostPort(servername)

	auth := smtp.PlainAuth("", "soulmanjohn@163.com", "passwd", host)

	// TLS config
	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         host,
	}

	// Here is the key, you need to call tls.Dial instead of smtp.Dial
	// for smtp servers running on 465 that require an ssl connection
	// from the very beginning (no starttls)
	conn, err := tls.Dial("tcp", servername, tlsconfig)
	if err != nil {
		log.Panic(err)
	}

	c, err := smtp.NewClient(conn, host)
	if err != nil {
		log.Panic(err)
	}

	// Auth
	if err = c.Auth(auth); err != nil {
		log.Panic(err)
	}

	// To && From
	if err = c.Mail(from.Address); err != nil {
		log.Panic(err)
	}

	if err = c.Rcpt(to.Address); err != nil {
		log.Panic(err)
	}

	// Data
	w, err := c.Data()
	if err != nil {
		log.Panic(err)
	}

	_, err = w.Write([]byte(message))
	if err != nil {
		log.Panic(err)
	}

	err = w.Close()
	if err != nil {
		log.Panic(err)
	}

	c.Quit()

	fmt.Printf("send [%s]email successfully!", message)
}

//It works.
func SendShortEmail(to string, message string, subject string) {
	//TODO ....(password), Set up authentication information.
	auth := smtp.PlainAuth("", "soulmanjohn@163.com", "password...TODO", "smtp.163.com")

	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	_to := []string{to}
	msg := []byte("To: " + to + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"\r\n" +
		"" + message + "\r\n")
	err := smtp.SendMail("smtp.163.com:25", auth, "soulmanjohn@163.com", _to, msg)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("----------------------------------------------\n send short email successfully!")
}

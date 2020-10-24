package helper

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"html/template"
	"log"
	"net/mail"
	"net/smtp"

	"github.com/Mynor2397/virtual-parish-office/internal/lib"
	"github.com/Mynor2397/virtual-parish-office/internal/models"
)

// Dest contiene la información del usuario a presentar en el template
type Dest struct {
	Firstname       string
	Secondname     string
	Thirdname      string
	Lastname       string
	Secondlastname string
}

func checkError(err error) error {
	if err != nil {
		return err
	}
	return nil
}

// SendEmail envia email a cada usuario que se registra
func SendEmail(person models.Person) error {
	from := mail.Address{Name: "Parroquia Inmaculada Concepción", Address: "mynor2397cas@gmail.com"}
	to := mail.Address{Name: lib.Config().NAMEMAIL, Address: lib.Config().ADDRESSMAIL}

	subject := "Solicitud de búsqueda"
	dest := Dest{
		Firstname: person.Firstname,
		Secondname: person.Secondname,
		Thirdname: person.Thirdname,
		Lastname: person.Lastname,
		Secondlastname: person.Secondlastname,
	}

	log.Println(dest)
	headers := make(map[string]string)
	headers["From"] = from.String()
	headers["To"] = to.String()
	headers["Subject"] = subject
	headers["Content-Type"] = `text/html; charset="UTF-8"`

	message := ""
	for k, v := range headers {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}

	t, err := template.ParseFiles("./internal/templates/email.html")
	checkError(err)

	buf := new(bytes.Buffer)
	err = t.Execute(buf, dest)
	checkError(err)

	message += buf.String()
	servername := "smtp.gmail.com:465"
	host := "smtp:gmail.com"

	auth := smtp.PlainAuth("", "mynor2397cas@gmail.com", "Miprincesa1009", host)

	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         host,
	}

	conn, err := tls.Dial("tcp", servername, tlsConfig)
	checkError(err)

	client, err := smtp.NewClient(conn, host)
	checkError(err)

	err = client.Auth(auth)
	checkError(err)

	err = client.Mail(from.Address)
	checkError(err)

	err = client.Rcpt(to.Address)
	checkError(err)

	w, err := client.Data()
	checkError(err)

	_, err = w.Write([]byte(message))
	checkError(err)

	err = w.Close()
	checkError(err)

	client.Quit()
	return nil
}

package send

import (
	"izihrm/utils"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

// Mail ...
func Mail(content string, name string, email string) string {
	from := mail.NewEmail("IZIHRM", "izi.hrm.2020@gmail.com")
	subject := "Activate your account on IZIHRM"
	to := mail.NewEmail(name, email)
	plainTextContent := content
	htmlContent := content
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(utils.ViperEnvVariable("SENDGRID_API_KEY"))
	response, _ := client.Send(message)

	return response.Body
}

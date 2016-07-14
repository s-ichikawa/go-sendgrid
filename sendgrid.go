package main

import (
    "github.com/sendgrid/sendgrid-go/helpers/mail"
    "github.com/sendgrid/sendgrid-go"
    "os"
    "fmt"
)

func main()  {
    from := mail.NewEmail("Example User", "ichikawa.shingo.0829@gmail.com")
    subject := "Hello World from the SendGrid Go Library"
    to := mail.NewEmail("Example User", "ichikawa.shingo.0829@gmail.com")
    content := mail.NewContent("text/plain", "some text here")
    m := mail.NewV3MailInit(from, subject, to, content)

    request := sendgrid.GetRequest(os.Getenv("SENDGRID_API_KEY"), "/v3/mail/send", "https://api.sendgrid.com")
    request.Method = "POST"
    request.Body = mail.GetRequestBody(m)
    response, err := sendgrid.API(request)
    if err != nil {
       fmt.Println(err)
    } else {
        fmt.Println(response.StatusCode)
        fmt.Println(response.Body)
        fmt.Println(response.Headers)
    }
}

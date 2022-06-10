/* RZFeeser | Alta3 Research
   Sending an Email (SMTP) message   */
   
package main

import (
    "log"
    "net/smtp"
)

func main() {

    // Configuration
    from := "my_email@gmail.com"            // update this to reflect your value
    password := "super_secret_password"     // update this to reflect your value
    to := []string{"recipient@email.com"}   // update this to reflect your value
    smtpHost := "smtp.gmail.com"
    smtpPort := "587"

    // update this to be your message body 
    message := []byte("From: my_email@gmail.com\r\n" + "To: recipient@email.com\r\n" + "Subject: Testing email from go\r\n\n" + "My super secret message.")

    // Create authentication
    auth := smtp.PlainAuth("", from, password, smtpHost)

    // Send actual message
    err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
    if err != nil {
        log.Fatal(err)
    }
}
/*
022/06/10 14:20:49 535 5.7.8 Username and Password not accepted. Learn more at
5.7.8  https://support.google.com/mail/?p=BadCredentials h2-20020a05620a244200b006a6d3fa430csm10626294qkn.58 - gsmtp
exit status 1
*/

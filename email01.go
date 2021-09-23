/* RZFeeser | Alta3 Research
   Sending an Email (SMTP) message   */
   
package main

import (
    "fmt"	
    "log"
    "net/smtp"
)

func main() {

    // Configuration
    from := "tbGoBurner@gmail.com"            // update this to reflect your value
    password := "Goburnerfr3ak$how"     // update this to reflect your value
    to := []string{"todd.butcher71@gmail.com"}   // update this to reflect your value
    smtpHost := "smtp.gmail.com"
    smtpPort := "587"

    // update this to be your message body 
    message := []byte("from:tbGoBurner@gmail.com\r\n" +"To:recipient@gmail.com\r\n" + "Subject: Message from go\n" + " Coming to you live from Ubuntu.")

    // Create authentication
    auth := smtp.PlainAuth("", from, password, smtpHost)

    // Send actual message
    err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("email sent")
}


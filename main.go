package main

import (
    "fmt"
    "log"
    "github.com/xuri/excelize/v2"
    "gopkg.in/gomail.v2"
)

// Function to read emails from an Excel file
func readEmailsFromExcel(filename string) ([]string, error) {
    f, err := excelize.OpenFile(filename)
    if err != nil {
        return nil, err
    }

    sheetName := f.GetSheetName(1) // Assumes emails are in the first sheet
    rows, err := f.GetRows(sheetName)
    if err != nil {
        return nil, err
    }

    var emails []string
    for _, row := range rows {
        if len(row) > 0 {
            emails = append(emails, row[0]) // Assumes emails are in the first column
        }
    }
    return emails, nil
}

// Function to send a welcome email
func sendWelcomeEmail(host string, port int, user, pass, toEmail string) error {
    m := gomail.NewMessage()
    m.SetHeader("From", user)
    m.SetHeader("To", toEmail)
    m.SetHeader("Subject", "Welcome!")
    m.SetBody("text/plain", "Hello! Welcome to our service.")

    d := gomail.NewDialer(host, port, user, pass)
    return d.DialAndSend(m)
}

// Main function to send emails
func main() {
    // Configuration of the Excel file and reading emails
    excelFile := "data/emails.xlsx"
    emails, err := readEmailsFromExcel(excelFile)
    if err != nil {
        log.Fatalf("Error reading the Excel file: %v", err)
    }

    // SMTP server configuration
    smtpHost := "smtp.example.com" // Replace with your SMTP server
    smtpPort := 587                // Replace with your SMTP server port
    smtpUser := "your-email@example.com" // Replace with your email
    smtpPass := "your-email-password"    // Replace with your email password

    // Sending emails
    for _, email := range emails {
        err := sendWelcomeEmail(smtpHost, smtpPort, smtpUser, smtpPass, email)
        if err != nil {
            log.Printf("Error sending email to %s: %v", email, err)
        } else {
            fmt.Printf("Email successfully sent to %s\n", email)
        }
    }
}

package email

import (
	"fmt"
	"log"
	"net/smtp"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

func SendOTPEmail(email, otp string) {
	rootPath, err := filepath.Abs("../")
	if err != nil {
		log.Fatal("Failed to get root directory:", err)
	}

	envPath := filepath.Join(rootPath, ".env")
	err = godotenv.Load(envPath)
	if err != nil {
		log.Fatal("Error loading .env file from:", envPath)
	}

	from := os.Getenv("EMAIL_ADDRESS")
	password := os.Getenv("EMAIL_PASSWORD")

	if from == "" || password == "" {
		log.Fatal("Email credentials not set in environment variable")
	}

	to := []string{email}
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	message := []byte(fmt.Sprintf(
		"Subject: Your OTP Code\nMIME-Version: 1.0\nContent-Type: text/html; charset=\"UTF-8\"\n\n"+
			"<html><body>"+
			"<h2 style='color: #2d89ef;'>Hello,</h2>"+
			"<p>Your One-Time Password (OTP) for verification is:</p>"+
			"<h1 style='color: #d9534f;'>%s</h1>"+
			"<p>Please use this OTP within the next 10 minutes. Do not share it with anyone.</p>"+
			"<p>If you didn't request this OTP, please ignore this email.</p>"+
			"<br>"+
			"<p>Best regards,</p>"+
			"<p><strong>Zyra Moments Team</strong></p>"+
			"</body></html>", otp))

	auth := smtp.PlainAuth("", from, password, smtpHost)
	err = smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		log.Printf("Failed to send email: %v", err)
	} else {
		log.Printf("OTP email sent to %s", email)
	}
}

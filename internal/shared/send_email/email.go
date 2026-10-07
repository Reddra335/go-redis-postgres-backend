package sendemail

import (
	"fmt"
	"os"
	"strconv"

	"gopkg.in/gomail.v2"
)

func SendEmail(emailUser string, otp string) error {

	msg := gomail.NewMessage()
	msg.SetHeader("FROM", os.Getenv("SMTP_EMAIL"))
	msg.SetHeader("To", emailUser)
	msg.SetHeader("Subject", "Verifikasi OTP")
	msg.SetBody("text/plain", "kode OTP anda berlaku 5 menit: "+otp)

	htmlBody := fmt.Sprintf(`
		<div style="font-family: Arial, sans-serif; padding: 20px; background-color: #f4f4f4;">
			<div style="max-width: 500px; margin: 0 auto; background: #ffffff; padding: 20px; border-radius: 8px;">
				<h2 style="color: #333333; text-align: center;">Verifikasi Akun</h2>
				<p>Gunakan kode OTP berikut untuk melanjutkan proses verifikasi:</p>
				<div style="text-align: center; margin: 20px 0;">
					<span style="font-size: 32px; font-weight: bold; letter-spacing: 5px; color: #2563eb; background: #eff6ff; padding: 10px 20px; border-radius: 6px;">%s</span>
				</div>
				<p style="color: #666666; font-size: 12px;">Kode ini hanya berlaku selama 5 menit. Jangan berikan kode ini kepada siapapun.</p>
			</div>
		</div>
	`, otp)

	msg.AddAlternative("text/html", htmlBody)
	port, err := strconv.Atoi(os.Getenv("SMTP_PORT"))
	if err != nil {
		return err
	}

	dialer := gomail.NewDialer(
		os.Getenv("SMTP_HOST"),
		port,
		os.Getenv("SMTP_EMAIL"),
		os.Getenv("SMTP_PASSWORD"),
	)

	return dialer.DialAndSend(msg)
}

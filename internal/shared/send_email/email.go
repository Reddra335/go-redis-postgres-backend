package sendemail

import (
	"fmt"
	"newProjek/internal/shared/config"
	"newProjek/internal/shared/logger"
	"strconv"

	"github.com/sirupsen/logrus"
	"gopkg.in/gomail.v2"
)

// Interface dengan penamaan yang lebih idiomatic Go
type EmailSender interface {
	SendOTP(emailUser string, otp string) error
}

type implnSendEmail struct {
	dialer *gomail.Dialer
	from   string
}

// Inisialisasi koneksi dialer
func NewImplnSendEmail(config *config.Config) EmailSender {
	send := config.SendEmail
	port, err := strconv.Atoi(send.Port)
	log := logger.New()
	if err != nil {
		log.WithFields(logrus.Fields{

			"operation": "send email",
			"layer":     "send email",
		}).WithError(err).Error("gagal mengkomper port")
	}
	//Host,Port,Email,Password
	dialer := gomail.NewDialer(send.Host, port, send.Email, send.Password)
	return &implnSendEmail{
		dialer: dialer,
		from:   send.Email,
	}
}

func (g *implnSendEmail) SendOTP(emailUser string, otp string) error {
	msg := gomail.NewMessage()
	msg.SetHeader("From", g.from)
	msg.SetHeader("To", emailUser)
	msg.SetHeader("Subject", "Verifikasi OTP")
	msg.SetBody("text/plain", "Kode OTP Anda berlaku 5 menit: "+otp)

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

	// Tinggal panggil dialer yang sudah siap dari constructor
	return g.dialer.DialAndSend(msg)
}

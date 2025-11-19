package libs

import (
	"fmt"
	"os"
	"strings"

	"gopkg.in/gomail.v2"
)

func SendEmail(to, subject string, pin any) error {
	m := gomail.NewMessage()
	body := strings.Replace(htmlBody, "PIN", fmt.Sprint(pin), 1)

	from := os.Getenv("SMTP_EMAIL")
	pass := os.Getenv("SMTP_PASSWORD")

	m.SetHeader("From", from)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	d := gomail.NewDialer(
		"smtp.gmail.com",
		587,
		from,
		pass,
	)

	return d.DialAndSend(m)
}

var htmlBody = `<!DOCTYPE html>
<html lang="id">
<head>
	<meta charset="UTF-8" />
	<meta name="viewport" content="width=device-width, initial-scale=1.0" />
	<title>Kode Verifikasi</title>
</head>
<body style="background-color: #f4f4f7; margin: 0; padding: 0; font-family: Arial, sans-serif;">
	<table role="presentation" cellspacing="0" cellpadding="0" border="0" width="100%">
		<tr>
			<td align="center" style="padding: 20px 0;">
				<table role="presentation" cellspacing="0" cellpadding="0" border="0" width="100%" style="max-width: 600px; background: #ffffff; border-radius: 10px; overflow: hidden;">
					<!-- Header -->
					<tr>
						<td style="background: #111827; padding: 20px; text-align: center;">
							<h2 style="color: #ffffff; margin: 0; font-size: 24px;">DEAR COFFEE</h2>
						</td>
					</tr>

					<!-- Content -->
					<tr>
						<td style="padding: 30px;">
							<h3 style="margin-top: 0; color: #111827; font-size: 20px;">Kode Verifikasi Anda</h3>
							<p style="margin: 0 0 20px; color: #4b5563; font-size: 16px; line-height: 1.5;">
								Kami menerima permintaan untuk reset password akun Anda.  
								Gunakan kode berikut untuk melanjutkan proses verifikasi:
							</p>

							<div style="
								text-align: center;
								font-size: 32px;
								font-weight: bold;
								color: #111827;
								background: #f3f4f6;
								padding: 15px 0;
								border-radius: 8px;
								letter-spacing: 4px;
								margin-bottom: 30px;
							">
								PIN
							</div>

							<p style="color: #4b5563; font-size: 14px; line-height: 1.5;">
								Kode ini hanya berlaku selama <strong>10 menit</strong>.  
								Jika Anda tidak merasa meminta reset password, Anda dapat mengabaikan email ini.
							</p>
						</td>
					</tr>

					<tr>
						<td style="background: #f9fafb; padding: 20px; text-align: center;">
							<p style="margin: 0; font-size: 12px; color: #9ca3af;">
								&copy; 2025 Dear Coffee Production. All rights reserved.
							</p>
						</td>
					</tr>
				</table>
			</td>
		</tr>
	</table>
</body>
</html>
`

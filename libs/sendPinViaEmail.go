package libs

import (
	"fmt"
	"os"
	"strings"

	"gopkg.in/gomail.v2"
)

func SendEmail(to, subject string, pin any, link string) error {
	m := gomail.NewMessage()
	body := strings.Replace(htmlBody, "LINK", link, 1)
	body = strings.Replace(body, "PIN", fmt.Sprint(pin), 1)

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

var htmlBody = `
<!DOCTYPE html>
<html lang="id">

<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Kode Verifikasi - Dear Coffee</title>
    <link href="https://fonts.googleapis.com/css2?family=Poppins:wght@400;500;600;700&display=swap" rel="stylesheet">
    <style>
        /* Reset CSS untuk kompatibilitas email */
        .container {
            max-width: 600px;
            margin: 0 auto;
            background-color: #ffffff;
            border-radius: 16px;
            overflow: hidden;
            box-shadow: 0 10px 30px rgba(0, 0, 0, 0.08);
        }

        /* Media query untuk responsivitas */
        @media only screen and (max-width: 600px) {
            .container {
                width: 100% !important;
                border-radius: 0 !important;
            }

            .header h1 {
                font-size: 24px !important;
            }

            .verification-code {
                font-size: 28px !important;
                letter-spacing: 6px !important;
            }

            .content {
                padding: 30px 20px !important;
            }
        }
    </style>
</head>

<body
    style="background: linear-gradient(135deg, #f5f7fa 0%, #c3cfe2 100%); margin: 0; padding: 0; font-family: 'Poppins', Arial, sans-serif;">
    <table role="presentation" cellspacing="0" cellpadding="0" border="0" width="100%">
        <tr>
            <td align="center" style="padding: 40px 20px;">

                <div class="container">
                    <table role="presentation" cellspacing="0" cellpadding="0" border="0" width="100%">
                        <tr>
                            <td class="header"
                                style="background: linear-gradient(135deg, #8B4513 0%, #D2691E 100%); padding: 30px 20px; text-align: center;">
                                <h1 style="color: #ffffff; margin: 0; font-size: 28px; font-weight: 700;">DEAR COFFEE
                                </h1>
                                <p style="color: rgba(255,255,255,0.8); margin: 8px 0 0; font-size: 16px;">Kode
                                    Verifikasi Anda</p>
                            </td>
                        </tr>
                    </table>

                    <table role="presentation" cellspacing="0" cellpadding="0" border="0" width="100%">
                        <tr>
                            <td class="content" style="padding: 40px 30px;">
                                <p style="margin: 0 0 20px; color: #4b5563; font-size: 16px; line-height: 1.6;">
                                    Halo Pelanggan Dear Coffee,
                                </p>
                                <p style="margin: 0 0 25px; color: #4b5563; font-size: 16px; line-height: 1.6;">
                                    Kami menerima permintaan untuk reset password akun Anda. Gunakan kode berikut untuk
                                    melanjutkan proses verifikasi:
                                </p>

                                <table role="presentation" cellspacing="0" cellpadding="0" border="0" width="100%"
                                    style="margin: 30px 0;">
                                    <tr>
                                        <td align="center">
                                            <div style="
                                                background: linear-gradient(135deg, #f8f9fa 0%, #e9ecef 100%);
                                                border: 2px dashed #8B4513;
                                                border-radius: 12px;
                                                padding: 20px;
                                                display: inline-block;
                                            ">
                                                <span class="verification-code" style="
                                                    font-size: 36px;
                                                    font-weight: 700;
                                                    color: #8B4513;
                                                    letter-spacing: 8px;
                                                    text-shadow: 0 2px 4px rgba(0,0,0,0.1);
                                                ">PIN</span>
                                            </div>
                                        </td>
                                    </tr>
                                </table>

                                <p style="color: #4b5563; font-size: 14px; line-height: 1.6; margin-bottom: 30px;">
                                    Kode ini hanya berlaku selama <strong style="color: #8B4513;">10 menit</strong>.
                                    Jika Anda tidak merasa meminta reset password, Anda dapat mengabaikan email ini.
                                </p>

                                <table role="presentation" cellspacing="0" cellpadding="0" border="0" width="100%">
                                    <tr>
                                        <td align="center">
                                            <a href="LINK" style="
                                                background: linear-gradient(135deg, #8B4513 0%, #D2691E 100%);
                                                color: #ffffff;
                                                text-decoration: none;
                                                padding: 14px 30px;
                                                border-radius: 50px;
                                                font-weight: 600;
                                                font-size: 16px;
                                                display: inline-block;
                                                box-shadow: 0 4px 15px rgba(139, 69, 19, 0.3);
                                            ">Reset Password Sekarang</a>
                                        </td>
                                    </tr>
                                </table>
                            </td>
                        </tr>
                    </table>

                    <table role="presentation" cellspacing="0" cellpadding="0" border="0" width="100%">
                        <tr>
                            <td
                                style="background: #f9fafb; padding: 25px; text-align: center; border-top: 1px solid #e5e7eb;">
                                <p style="margin: 0 0 10px; font-size: 14px; color: #6b7280;">
                                    Butuh bantuan? <a href="mailto:support@dearcoffee.com"
                                        style="color: #8B4513; text-decoration: none;">Hubungi Tim Support Kami</a>
                                </p>
                                <p style="margin: 0; font-size: 12px; color: #9ca3af;">
                                    &copy; 2025 Dear Coffee Production. All rights reserved.
                                </p>
                                <p style="margin: 15px 0 0; font-size: 11px; color: #9ca3af;">
                                    Jl. Kopi Aroma No. 123, Jakarta Selatan
                                </p>
                            </td>
                        </tr>
                    </table>
                </div>

                <table role="presentation" cellspacing="0" cellpadding="0" border="0" width="100%"
                    style="max-width: 600px; margin-top: 20px;">
                    <tr>
                        <td align="center">
                            <p style="margin: 0; font-size: 12px; color: #6b7280;">
                                <strong>Keamanan:</strong> Jangan bagikan kode ini kepada siapapun. Tim Dear Coffee
                                tidak akan pernah meminta kode verifikasi Anda.
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

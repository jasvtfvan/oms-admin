package crypto

import (
	"encoding/base64"

	"github.com/jasvtfvan/oms-admin/server/utils/crypto/internal"
)

var privateKey = []byte(`
-----BEGIN RSA PRIVATE KEY-----
MIIEpAIBAAKCAQEA7TQDqigwjCPWURt/BuAS37NMQ3U2uBGpf5gBzNlyZjp96lkU
vF93BqItt9UAAtNUVh4rjOgmIWX3xcCnu6ETB8bho6kfbBy8qoTb9lTekz6Y2WHX
3eydLozFxC8h1exZ6V3kYoiztUkChz6v+kwNs6kGospt8r0rax81Eav3RgajJmpd
4OXMOpVX4lNdRLOgzK9jUOmlaVQXJDHRbyzWoC1ooNguKIOpGWQ5doKD+tznUwhi
sjTxmuW/tNbWDd2nPl/pTjYIY2RE9xigtOovAHfY3AUY7NN+mGZn9UxTI2jjdzOQ
TB9un7izQvOB+otEF8Wb67dtHSFlDB7W8TbQVQIDAQABAoIBAA+za7Ktqlj8XklM
GqJn3pf0FE46yf5xHNkXRLc8hXgC0ybZ8qdtYkGMJp6OeMu3FVQF9zgCfdOkHjx2
viOLS+kt3u2oWi4b0NkwpiauA3WXpSJueY11Bgp4wvZzcDfqxyDNWDq1db/AL2yo
V5mnwxhrTcckwxZYMzGKBUdALMPzt+FG4pYDz/dI5a4I0rchyJHx/QcBk9asFspw
xDukLcnSchc8AyVQwpTwOhY1vEFV4sgaHMoUnlWo1ytkk808cJXl4yiJfvo2h+lE
WfhQmxR8fIsrLdaWidAgs/CnxWAJvQ+NV2fBUhU6TAjKolhrjoZbK2XM62dKxI9X
NLKUM4ECgYEA/mKNsGoLbigJacbrKdNfTFIJpdf7h5DITWsx3SxLsxwlw1sxzBoV
B3zT0QtKn4npjxuFsw/HGg6L7jDnxZH/lqj0EyLCuxeLJ6a51Q5KV6J8ztYXEEWd
ZkLg8T7F+PHHAoBsGf01g2+yNNTstsulxAeSRnsMXZM5Ze3VIv+zNiUCgYEA7rWJ
H2zMFl2PiDHro3PlhT16A0kX91q720cDGeNA5F8S3P/j2j9M/+7H8B+xFA+R9zMX
OqiZ9aDvqdg4SdsZy/1+L7obKiQ4WtFbp2CZvVY3aHcEC54Z3O1kLRooDBJKHrvc
Lag3lwCOYt2FuS5P55PvuA2gA16Uj3VnzCbPInECgYEAxxejBRs57vDuzRaeHpIL
19OtMVskxSkPW2g2EoAEjx3MgGTzSGZxZvbPYKCRuuNZJGPJ9Ca5ES+pXLZx7zMg
8m0w+XkPJxZ6FoJqltEkZgoJ3Ge6jUWutsZI/wa+MuQneVHBSWXfaAsXUjoDOd2Q
0yeJ2Bedye7b0WaelVHClPECgYACR7WhmTZx2D7wvBlWHFtK5IVv1pjmAfXdaFY9
PxB1nfreJYuVoBkqMKu0PXlBicyJIfHM26Ns1zay1p/jBLbAXhGAfzSXOHVZWLqZ
ZLDTQCmTU9+0BLSWiaX1UFSlmN8gYAcAYKT3SkgR5a/LTwfwXFdj2K14msSsgiCV
sKE14QKBgQDV1Jc8iVda1/AVizhDIZbus+Es8ANUcI7JV2QDZdLPkLb1g8jietrX
QBN8JMuAtMILo0DOCuulUN75pf/NCaoeelEHz7gmLsFNrvAn3bXD38qTtsh6K+Eu
M+9lY/9QvBUuDGWIkeWE+yzLH/3kLaWBGmywRxXQ2zXVewHlpjlRWw==
-----END RSA PRIVATE KEY-----
`)

var publicKey = []byte(`
-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA7TQDqigwjCPWURt/BuAS
37NMQ3U2uBGpf5gBzNlyZjp96lkUvF93BqItt9UAAtNUVh4rjOgmIWX3xcCnu6ET
B8bho6kfbBy8qoTb9lTekz6Y2WHX3eydLozFxC8h1exZ6V3kYoiztUkChz6v+kwN
s6kGospt8r0rax81Eav3RgajJmpd4OXMOpVX4lNdRLOgzK9jUOmlaVQXJDHRbyzW
oC1ooNguKIOpGWQ5doKD+tznUwhisjTxmuW/tNbWDd2nPl/pTjYIY2RE9xigtOov
AHfY3AUY7NN+mGZn9UxTI2jjdzOQTB9un7izQvOB+otEF8Wb67dtHSFlDB7W8TbQ
VQIDAQAB
-----END PUBLIC KEY-----
`)

func GenerateKeyPair() (string, string, error) {
	p1, p2, err := internal.GetRsaKeyPair()
	return p1, p2, err
}

func RsaEncrypt(orig string) string {
	enBytes, _ := internal.RsaEncrypt([]byte(orig), publicKey)
	return base64.StdEncoding.EncodeToString(enBytes)
}

func RsaDecrypt(encrypted string) string {
	enBytes, _ := base64.StdEncoding.DecodeString(encrypted)
	deBytes, _ := internal.RsaDecrypt(enBytes, privateKey)
	return string(deBytes)
}

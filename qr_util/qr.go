package qr_util

import (
	qrCodeEn "github.com/skip2/go-qrcode"
	qrCodeDe "github.com/tuotoo/qrcode"
	"os"
)

func Generate(content string, size int) (png []byte, err error) {
	png, err = qrCodeEn.Encode(content, qrCodeEn.Medium, size)
	return png, err
}

func GenerateToFile(content string, size int, filename string) (err error) {
	err = qrCodeEn.WriteFile(content, qrCodeEn.Medium, size, filename)
	return err
}

func Decode(file *os.File) (content string, err error) {
	qrMatrix, err := qrCodeDe.Decode(file)
	if err != nil {
		return "", err
	}
	return qrMatrix.Content, nil
}

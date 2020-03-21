package qr_util

import (
	"os"
	"testing"
)

func TestGenerateToFile(t *testing.T) {
	err := GenerateToFile("https://www.baidu.com", 256, "qrCode.png")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("ok")
}

func TestDecode(t *testing.T) {
	file, err := os.Open("qrCode.png")
	if err != nil {
		t.Error(err)
		return
	}
	content, err := Decode(file)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(content, "ok")
}

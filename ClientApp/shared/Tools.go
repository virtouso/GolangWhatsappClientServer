package shared

import (
	"fmt"
	"github.com/mdp/qrterminal/v3"
	"github.com/skip2/go-qrcode"
	"os"
)

func GenerateQRCode(text string, filename string, size int) error {

	err := qrcode.WriteFile(text, qrcode.Medium, size, filename)
	if err != nil {
		return fmt.Errorf("failed to generate QR code: %v", err)
	}
	return nil
}

func RenderQRCodeInTerminal(text string) {
	qrterminal.GenerateHalfBlock(text, qrterminal.L, os.Stdout)
}

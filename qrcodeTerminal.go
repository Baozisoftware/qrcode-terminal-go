package qrcodeTerminal

import (
	"fmt"

	"github.com/skip2/go-qrcode"
	"github.com/mattn/go-colorable"
)

type consoleColor string
type consoleColors struct {
	NormalBlack   consoleColor
	NormalRed     consoleColor
	NormalGreen   consoleColor
	NormalYellow  consoleColor
	NormalBlue    consoleColor
	NormalMagenta consoleColor
	NormalCyan    consoleColor
	NormalWhite   consoleColor
	BrightBlack   consoleColor
	BrightRed     consoleColor
	BrightGreen   consoleColor
	BrightYellow  consoleColor
	BrightBlue    consoleColor
	BrightMagenta consoleColor
	BrightCyan    consoleColor
	BrightWhite   consoleColor
}
type qrcodeRecoveryLevel qrcode.RecoveryLevel
type qrcodeRecoveryLevels struct {
	Low     qrcodeRecoveryLevel
	Medium  qrcodeRecoveryLevel
	High    qrcodeRecoveryLevel
	Highest qrcodeRecoveryLevel
}

var (
	ConsoleColors consoleColors = consoleColors{
		NormalBlack:   "\033[38;5;0m  \033[0m",
		NormalRed:     "\033[38;5;1m  \033[0m",
		NormalGreen:   "\033[38;5;2m  \033[0m",
		NormalYellow:  "\033[38;5;3m  \033[0m",
		NormalBlue:    "\033[38;5;4m  \033[0m",
		NormalMagenta: "\033[38;5;5m  \033[0m",
		NormalCyan:    "\033[38;5;6m  \033[0m",
		NormalWhite:   "\033[38;5;7m  \033[0m",
		BrightBlack:   "\033[48;5;0m  \033[0m",
		BrightRed:     "\033[48;5;1m  \033[0m",
		BrightGreen:   "\033[48;5;2m  \033[0m",
		BrightYellow:  "\033[48;5;3m  \033[0m",
		BrightBlue:    "\033[48;5;4m  \033[0m",
		BrightMagenta: "\033[48;5;5m  \033[0m",
		BrightCyan:    "\033[48;5;6m  \033[0m",
		BrightWhite:   "\033[48;5;7m  \033[0m"}
	QRCodeRecoveryLevels = qrcodeRecoveryLevels{
		Low:     qrcodeRecoveryLevel(qrcode.Low),
		Medium:  qrcodeRecoveryLevel(qrcode.Medium),
		High:    qrcodeRecoveryLevel(qrcode.High),
		Highest: qrcodeRecoveryLevel(qrcode.Highest)}
)

type QRCodeString string

func (v *QRCodeString) Print() {
	fmt.Fprint(outer, *v)
}

type qrcodeTerminal struct {
	front consoleColor
	back  consoleColor
	level qrcodeRecoveryLevel
}

func (v *qrcodeTerminal) Get(content interface{}) (result *QRCodeString) {
	var qr *qrcode.QRCode
	var err error
	if t, ok := content.(string); ok {
		qr, err = qrcode.New(t, qrcode.RecoveryLevel(v.level))
	} else if t, ok := content.([]byte); ok {
		qr, err = qrcode.New(string(t), qrcode.RecoveryLevel(v.level))
	}
	str := ""
	if qr != nil && err == nil {
		for ir, row := range qr.Bitmap() {
			lr := len(row)
			if ir == 0 || ir == 1 || ir == 2 ||
				ir == lr-1 || ir == lr-2 || ir == lr-3 {
				continue
			}
			for ic, col := range row {
				lc := len(qr.Bitmap())
				if ic == 0 || ic == 1 || ic == 2 ||
					ic == lc-1 || ic == lc-2 || ic == lc-3 {
					continue
				}
				if col {
					str += fmt.Sprint(v.front)
				} else {
					str += fmt.Sprint(v.back)
				}
			}
			str += fmt.Sprintln()
		}
		obj := QRCodeString(str)
		result = &obj
	}
	return
}

func New2(front, back consoleColor, level qrcodeRecoveryLevel) *qrcodeTerminal {
	obj := qrcodeTerminal{front: front, back: back, level: level}
	return &obj
}

func New() *qrcodeTerminal {
	front, back, level := ConsoleColors.BrightBlack, ConsoleColors.BrightWhite, QRCodeRecoveryLevels.Medium
	return New2(front, back, level)
}

var outer = colorable.NewColorableStdout()
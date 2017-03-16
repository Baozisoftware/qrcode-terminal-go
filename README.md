# qrcode-terminal-go
QRCode terminal for golang.

# Example
```go
package main

import "github.com/Baozisoftware/qrcode-terminal-go"

func main() {
	Test1()
	Test2()
}

func Test1(){
	content := "Hello, 世界"
	obj := qrcodeTerminal.New()
	obj.Get([]byte(content)).Print()
}

func Test2(){
	content := "https://github.com/Baozisoftware/qrcode-terminal-go"
	obj := qrcodeTerminal.New2(qrcodeTerminal.ConsoleColors.BrightBlue,qrcodeTerminal.ConsoleColors.BrightGreen,qrcodeTerminal.QRCodeRecoveryLevels.Low)
	obj.Get([]byte(content)).Print()
}
```

## Screenshot

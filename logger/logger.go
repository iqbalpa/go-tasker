package logger

import "fmt"

func Info(msg string) {
	str := fmt.Sprintf("[INFO] %s", msg)
	fmt.Println(str)
}

func Warn(msg string) {
	str := fmt.Sprintf("[WARN] %s", msg)
	fmt.Println(str)
}

func Error(msg string) {
	str := fmt.Sprintf("[ERROR] %s", msg)
	fmt.Println(str)
}

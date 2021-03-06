package logger

import (
	"fmt"
)

const (
	InfoColor        = "\033[1;0m%s\033[0m"
	InfoSuccessColor = "\033[1;32m%s\033[0m"
	WarningColor     = "\033[1;33m%s\033[0m"
	ErrorColor       = "\033[1;31m%s\033[0m"
)

func Info(msg interface{}) {
	s := fmt.Sprintf(InfoColor, msg)
	fmt.Print(s)
}

func Infoln(msg interface{}) {
	s := fmt.Sprintf(InfoColor, msg)
	fmt.Println(s)
}

func InfoSuccess(msg interface{}) {
	s := fmt.Sprintf(InfoSuccessColor, msg)
	fmt.Print(s)
}

func InfoSuccessln(msg interface{}) {
	s := fmt.Sprintf(InfoSuccessColor, msg)
	fmt.Println(s)
}

func Successf(msg interface{}) string {
	return fmt.Sprintf(InfoSuccessColor, msg)
}

func Error(msg interface{}) {
	s := fmt.Sprintf(ErrorColor, msg)
	fmt.Print(s)
}

func Errorln(msg interface{}) {
	s := fmt.Sprintf(ErrorColor, msg)
	fmt.Println(s)
}

func Errorf(msg interface{}) string {
	return fmt.Sprintf(ErrorColor, msg)
}

func Warn(msg interface{}) {
	s := fmt.Sprintf(WarningColor, msg)
	fmt.Print(s)
}

func Warnln(msg interface{}) {
	s := fmt.Sprintf(WarningColor, msg)
	fmt.Println(s)
}

func Warnf(msg interface{}) string {
	return fmt.Sprintf(WarningColor, msg)
}

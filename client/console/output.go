package console

import (
	"fmt"
)

var (
	split = "###############################################"
)

func PrintSeperatorLine() {
	fmt.Println(split)
}

func ListeAuf[T fmt.Stringer](items []T) {
	for i, item := range items {
		fmt.Println(fmt.Sprintf("%d. %s", i+1, item.String()))
	}
}

func Info(info string) {
	fmt.Println(info)
}

func Error(err string) {
	Info("Error: " + err)
}

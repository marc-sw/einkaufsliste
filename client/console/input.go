package console

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
)

func LeseString(info string) string {
	fmt.Print(info + ": ")
	scanner.Scan()
	return scanner.Text()
}

func LeseZahl(info string) int {
	for {
		var eingabe = LeseString(info)
		var zahl, err = strconv.Atoi(eingabe)
		if err == nil {
			return zahl
		}
		fmt.Println("Eingabe war keine Zahl.")
	}
}

func LeseBestimmteZahl(info string, min int, max int) int {
	for {
		var zahl = LeseZahl(info)

		if zahl >= min && zahl <= max {
			return zahl
		}
		fmt.Println(fmt.Sprintf("Eingabe muss zwischen %d und %d liegen.", min, max))
	}
}

func LeseBool(info string) bool {
	for {
		var eingabe = strings.ToLower(LeseString(info))
		if strings.HasPrefix(eingabe, "y") || strings.HasPrefix(eingabe, "j") || eingabe == "1" || eingabe == "true" {
			return true
		} else if strings.HasPrefix(eingabe, "n") || eingabe == "0" || eingabe == "false" {
			return false
		}
		fmt.Println("Eingabe muss ja oder nein sein.")
	}
}

func LeseOptionalString(info string, fallback string) string {
	var eingabe = LeseString(info)
	if len(eingabe) == 0 {
		eingabe = fallback
	}
	return eingabe
}

func LeseOptionalZahl(info string, fallback int) int {
	for {
		var eingabe = LeseString(info)
		if len(eingabe) == 0 {
			return fallback
		}
		var zahl, err = strconv.Atoi(eingabe)
		if err == nil {
			return zahl
		}
		fmt.Println("Eingabe war keine Zahl.")
	}
}

func LeseOptionalBool(info string, fallback bool) bool {
	for {
		var eingabe = strings.ToLower(LeseString(info))
		if len(eingabe) == 0 {
			return fallback
		}
		if strings.HasPrefix(eingabe, "y") || strings.HasPrefix(eingabe, "j") || eingabe == "1" || eingabe == "true" {
			return true
		}
		if strings.HasPrefix(eingabe, "n") || eingabe == "0" || eingabe == "false" {
			return false
		}
		fmt.Println("Eingabe muss ja oder nein sein.")
	}
}

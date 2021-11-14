package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	var inputString string
	var outputString string
	_, _ = fmt.Scanf("%s", &inputString)
	for i := 0; i < len(inputString); i++ {
		if unicode.IsUpper(rune(inputString[i])) {
			outputString += strings.ToLower(string(inputString[i]))
		} else {
			outputString += strings.ToUpper(string(inputString[i]))
		}
	}
	fmt.Println(outputString)
}

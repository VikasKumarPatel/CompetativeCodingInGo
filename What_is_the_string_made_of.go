package main

import (
	"fmt"
)

func main() {
	var numberString string
	_, err := fmt.Scanf("%s", &numberString)
	if err != nil {
		return
	}
	findNumberOfDashes(numberString)
}

// findNumberOfDashes scan and print number of dashes in given string within range of 1-100 chars
func findNumberOfDashes(numberString string) {
	if len(numberString) == 0 || len(numberString) > 100 {
		return
	}
	sum := 0
	for _, number := range numberString {
		switch number {
		case '0':
			sum += 6
			break
		case '1':
			sum += 2
			break
		case '2':
			sum += 5
			break
		case '3':
			sum += 5
			break
		case '4':
			sum += 4
			break
		case '5':
			sum += 5
			break
		case '6':
			sum += 6
			break
		case '7':
			sum += 3
			break
		case '8':
			sum += 7
			break
		case '9':
			sum += 6
			break
		}
	}
	if sum > 0 {
		fmt.Println(sum)
	}
}

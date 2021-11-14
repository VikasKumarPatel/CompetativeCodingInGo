package main

import "fmt"

func main() {
	var n int
	_, err := fmt.Scanf("%d", &n)
	if err != nil {
		return
	}
	divisibleBy10Number(n)
}

func divisibleBy10Number(arraySize int) {
	var tmpScan string
	for i := 0; i < arraySize; i++ {
		_, _ = fmt.Scanf("%s", &tmpScan)
		if i == arraySize-1 {
			if tmpScan[len(tmpScan)-1] == '0' {
				fmt.Println("Yes")
			} else {
				fmt.Println("No")
			}
		}
	}
}

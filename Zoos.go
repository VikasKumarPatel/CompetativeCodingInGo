package main

import "fmt"

func main() {
	var inputString string
	_, err := fmt.Scanf("%s", &inputString)
	if err != nil {
		fmt.Println(err)
		return
	}
	var x, y int64

	for i := 0; i < len(inputString); i++ {
		if inputString[i] == 'z' || inputString[i] == 'Z' {
			x++
		}
		if inputString[i] == 'o' || inputString[i] == 'O' {
			y++
		}
	}
	if 2*x == y {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

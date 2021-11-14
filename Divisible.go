package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	var N int
	_, err := fmt.Scanf("%d", &N)
	if err != nil {
		return
	}
	divisibleNumber(N)
}

func divisibleNumber(n int) {
	if n%2 != 0 {
		return
	}
	half := n / 2
	var tmpScan string
	var number int64
	var numString string
	for i := 0; i < n; i++ {
		_, _ = fmt.Scanf("%s", &tmpScan)
		if half > i {
			val, _ := strconv.ParseInt(string(tmpScan[0]), 10, 64)
			number = number + (val * int64(math.Pow10(n-1-i)))
			numString += string(tmpScan[0])
		} else {
			val, _ := strconv.ParseInt(string(tmpScan[len(tmpScan)-1]), 10, 64)
			number = number + (val * int64(math.Pow10(n-1-i)))
			numString += string(tmpScan[len(tmpScan)-1])

		}
	}
	var evenSum int64
	var oddSum int64
	for i := 0; i < len(numString); i++ {
		val, _ := strconv.ParseInt(string(numString[i]), 10, 64)
		if i%2 == 0 {
			evenSum += val
		} else {
			oddSum += val
		}
	}
	var difference int64
	if evenSum < oddSum {
		difference = oddSum - evenSum
	} else {
		difference = evenSum - oddSum
	}
	if difference == 0 || difference%11 == 0 {
		fmt.Println("OUI")
	} else {
		fmt.Println("NON")
	}
}

package main

import "fmt"

func main() {
	var testCases uint8
	_, err := fmt.Scanf("%d", &testCases)
	if err != nil || testCases > 10 {
		return
	}
	var k uint8
	var participants uint8
	var greenCost, purpleCost int
	var minColor, maxColor int
	var l uint8

	for k = 0; k < testCases; k++ {

		_, _ = fmt.Scanf("%d", &greenCost)
		_, _ = fmt.Scanf("%d", &purpleCost)
		if greenCost < purpleCost {
			minColor = greenCost
			maxColor = purpleCost
		} else {
			minColor = purpleCost
			maxColor = greenCost
		}
		_, err = fmt.Scanf("%d", &participants)
		if err != nil || participants > 10 {
			return
		}
		i := 0
		iFlag, jFlag := 0, 0
		for l = 0; l < participants; l++ {
			_, _ = fmt.Scanf("%d", &i)
			if i == 1 {
				iFlag++
			}
			_, _ = fmt.Scanf("%d", &i)
			if i == 1 {
				jFlag++
			}
		}
		if iFlag < jFlag {
			fmt.Println((iFlag * maxColor) + (jFlag * minColor))
		} else {
			fmt.Println((jFlag * maxColor) + (iFlag * minColor))
		}
		iFlag, jFlag = 0, 0
	}
}

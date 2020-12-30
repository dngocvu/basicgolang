package main

import "fmt"

func findSum(array [] int, target int ) {
	giveNumberLength := len(array)
	var resultSum[2] int
	var i, j int
	for i = 0; i < giveNumberLength - 1; i++ {
		for j = 1; j < giveNumberLength; j++ {
			if (array[i] + array[j]) == target {
				resultSum[0] = i
				resultSum[1] = j
				fmt.Println("Result", resultSum)
				//return resultSum
			}
		}
	}
	//return
}

func main(){
	var giveNumber = [] int {2, 3, 7, 6}
	var target int = 9
	findSum(giveNumber, target)
	//result := findSum(giveNumber, target)
	//fmt.Println("Result", result)
}
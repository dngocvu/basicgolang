package main

import (
	"fmt"
	"math"
)

func ReverseNumber(inputNumber int32) (revNumber int32){
	
	if(inputNumber == 0){
		fmt.Println("0")
	}else{
		var flag bool = false
		if inputNumber < 0 {
			inputNumber = 0 - inputNumber
			flag = true
		}
		for inputNumber > 0 {
				revNumber = revNumber*10 + inputNumber%10
				inputNumber = inputNumber/10
			}
		if flag {
			revNumber = 0 - revNumber
			flag = false
		}
	}
	return revNumber
}
	
func main(){

	var num int32
	var numbers int32
	fmt.Printf("Enter your number: ")
	fmt.Scanf("%d", &num)
	fmt.Printf("the number just enter is: %d, type %T ", num, num)
	numbers = num % 10
	fmt.Printf("the length of number just enter is %d", numbers)
	fmt.Println("max integer: ", math.MaxInt32)
	fmt.Println("min integer: ", math.MinInt32)
	fmt.Println("reverse: ",ReverseNumber(num))	
}
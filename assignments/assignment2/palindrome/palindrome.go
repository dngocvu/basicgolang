package main

import (
	"fmt"
	"strconv"
)
func CheckPalindrome(inputNumber int) (result bool) {
		result = false
	if inputNumber < 0 {
		result = false
	} else {
		var inputNumberString string = ""
		var afterString string = ""
		inputNumberString = strconv.Itoa(inputNumber)
		afterString = inputNumberString
		// create a rune that include string
		cp1 := []rune(inputNumberString)
		for i, j := 0, len(cp1)-1; i < j; i, j = i+1, j-1 { 
  
			// swap the letters of the string, 
			// like first with last and so on. 
			cp1[i], cp1[j] = cp1[j], cp1[i] 
		}
		// casting rune to string and compare
		if afterString == string(cp1){
			// return the reversed string. 
			result = true
	
		}
	}
	return result
}

func main(){
	var number int
	fmt.Println("Enter your number")
	fmt.Scanf("%d", &number)
	fmt.Println("This number is a palidrome: ", CheckPalindrome(number))
}
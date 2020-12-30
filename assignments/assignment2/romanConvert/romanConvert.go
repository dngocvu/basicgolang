package main

import "fmt"

func convertRomanNumber(romanString string) int{
	var number int
	var romanMap = make(map[string]int)
	romanMap["I"] = 1
	romanMap["V"] = 5
	romanMap["X"] = 10
	romanMap["L"] = 50
	romanMap["C"] = 100
	romanMap["D"] = 500
	romanMap["M"] = 1000
	// not include in map -> has no mean
	length_string := len(romanString)
	if  length_string <= 1{
		number = romanMap[romanString]
	} else {
		var i int = 0
		romanString = romanString + "0"
		for i < length_string {
			fmt.Println("i: ", i)
			if romanMap[string(romanString[i])] < romanMap[string(romanString[i+1])] {
				number = number + romanMap[string(romanString[i+1])] - romanMap[string(romanString[i])]
				i = i + 2
			} else {
				number = number + romanMap[string(romanString[i])]
				i++
				fmt.Println("number: ", i)
			}
		}
	}
	return number
}

func main(){
	var string_roman string
	var result int
	fmt.Printf("Enter your roman number: ")
	fmt.Scanf("%s", &string_roman)
	result = convertRomanNumber(string_roman)
	fmt.Println("Your roman number is: ", string_roman)
	fmt.Println("Your digital number is: ", result)

}
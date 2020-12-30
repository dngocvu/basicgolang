package main

import "fmt"

func LongestPal(inputString string)(result string){
	n := len(inputString)
	var lenPal int
	var low, high int = 0, 0
	// length = 0
	if n == 0 {
		return ""
	}
	// length = 1
	if n == 1 {
		return inputString
	}
	result = result + string(inputString[0])
	for i := 1; i < n; i++{
		low = i - 1
		high = i
		for low >= 0 && high < n && inputString[low] == inputString[high]{
			if lenPal < (high - low + 1){
				lenPal = high - low + 1
				result = inputString[low:lenPal]
			}
			low --
			high ++
		}
		low = i - 1
		high = i + 1
		for low >= 0 && high < n && inputString[low] == inputString[high]{
			if lenPal < (high - low + 1){
				lenPal = high - low + 1
				result = inputString[low:lenPal]
			}
			low --
			high ++
		}
	}
	return result
}

func main() {
	var str string = "abbaba"
	result := LongestPal(str)
	fmt.Println("result:",result)
}
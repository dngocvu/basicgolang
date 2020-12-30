package main

import "fmt"

func Max(a, b int)(max int){
	if a > b {
		max = a
	} else {
		max = b
	}
	return
}
func LongestSubstring(inputString string)(result int){
	var limit int = 256
	lenString := len(inputString)
	var start int = 0
	count := make([]int,limit)
	for i := 0; i < lenString; i++ {
		start = Max(start, count[inputString[i]] + 1)
		result = Max(result, i - start + 1)
		count[inputString[i]] = i
	}
	return result + 1
}

func main() {
	var str string = "ab"
	result := LongestSubstring(str)
	fmt.Println("result:",result)
}
package main

import "fmt"

func LongestCommonPrefix(arrInput []string)(string){
	n := len(arrInput)
	var minLen int
	var result string
	minLen = len(arrInput[0])
	for i := 1; i < n; i++ {
		if minLen > len(arrInput[i]){
			minLen = len(arrInput[i])
		}
	}
	for i := 0; i < minLen ; i++ {
		currentMember := arrInput[0][i]
		for j := 1; j < n; j ++ {
			if arrInput[j][i] != currentMember {
				return result
			}
		}
		result = result + string(currentMember)
	}
	return result
}
func main(){
	var arr = []string {"lower", "flow", "flight"}
	result := LongestCommonPrefix(arr)
	if result == ""{
		fmt.Println("None")
	} else {
		fmt.Println("result:",result)
	}
}
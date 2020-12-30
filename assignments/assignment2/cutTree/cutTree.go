package main

import "fmt"

func CheckInCrease(subarr [] int)(increase bool){
	increase = true
	for i := 1; i < len(subarr); i++ {
		if subarr[i-1] > subarr[i]{
			return false
		}
	}
	return
}
//
func RemoveArray(s []int, index int) []int {
	arrayBefore := make([]int, 0)
	// create before array
	arrayBefore = append(arrayBefore, s[:index]...)
	// add after array without index element 
    return append(arrayBefore, s[index+1:]...)
}
func CutTree(arr []int)(res int){
	var count int = 0
	var subArr [] int
	for i := 0; i < len(arr) ; i++ {
		subArr = RemoveArray(arr, i)
		fmt.Println(subArr)
		if CheckInCrease(subArr) {
			count = count + 1
		}
	}
	return count
}

func main(){
	var arr = [] int {1, 2, 3, 3, 5, 6, 7}
	fmt.Println(CutTree(arr))
}
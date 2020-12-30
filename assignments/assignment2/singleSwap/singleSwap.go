package main

import "fmt"

func Swap(a *int, b *int){
	tem := *a
	*a = *b
	*b = tem
}

func Sort(arr []int)( []int){
	n := len(arr)
	for i:= 0; i < n - 1; i++ {
		min := i
		for j := i+1; j < n; j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		Swap(&arr[min], &arr[i])
	}
	return arr
}
func singleSwap(arr []int)(result bool){
	n := len(arr)
	var count int = 0
	cpy := make([]int, n)
	for i := 0; i < n; i++{
		cpy[i] = arr[i]
	}
	cpy = Sort(cpy)
	for i := 0; i < n; i++ {
		if arr[i] != cpy[i] {	
			count ++
		}
	}
	if count == 0 || count == 2 {
		result = true
	} else {
		result = false
	}
	return result
}

func main(){

	var arr = [] int {1, 2, 1, 7, 5}
	fmt.Println(singleSwap(arr))

}
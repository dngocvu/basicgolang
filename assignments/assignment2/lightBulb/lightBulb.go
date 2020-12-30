package main

import "fmt"

func Max(a int, b int)(max int ){
	if a > b {
		max = a
	} else {
		max = b
	}
	return
}

func lightBulb(arr []int)(result int){
	var count int = 0
	var max int
	// i : moment
	for i := 0; i < len(arr) ; i++ {
		max = Max(max, arr[i])
		count ++
		if max == count {
			result ++ 
		}
	}
	return result
}


func main(){

	var arr = [] int {2, 3, 4, 1, 5}
	fmt.Println(lightBulb(arr))

}
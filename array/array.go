package main

import "fmt"

// return array from function
func returnArray(x, y, z int) [3]int{
	var arr [3] int	
	arr[0] = x
	arr[1] = y
	arr[2] = z
	return arr
}
// passing array to function 
func passArray(myArray [11] int) int{
	var sum int
	for i := 0; i <= 10; i++{
		sum = sum + myArray[i]
	}
	return sum
}

func main(){
	// initialise array
	var balance [10] float64 // option 1
	var apple = [5] float32 {1,2,3,4,5} // option 2
	var banana = [] float32 {2,3,4,5,6,7} // option 3
	
	mingo := banana

	fmt.Println("balance is: ", balance)
	fmt.Println("apple is: ", apple)
	fmt.Println("mingo is: ", mingo)
	fmt.Println("banana is: ", banana)
	fmt.Println("apple's 5 element is: ", apple[4])
	fmt.Printf("type of mingo is: %T", mingo)
	var a int
	var x,y,z int = 56,60,70 
	for i := 0; i < 5; i++ {
		fmt.Printf("apple's element %d is %.f \n", i,apple[i])
	}
	fmt.Println("a = ", a)

	var b [3] int
	b = returnArray(x, y, z)
	fmt.Println("array just type is ", b)
	var c = [11] int {1,5,5,4,3}
	var getValue int
	getValue = passArray(c)
	fmt.Println("sum is ", getValue)

	// mutildemension arrays
	// var mutilArray [][] int32 // option 1
	var mutilArrayValue = [4][4] int32 {
		{1, 2, 3, 4},
		{2, 3, 4, 5},
		{3, 4, 5, 6},
		{4, 5, 6, 7}}

	for i := 0; i < 4; i++{
		for j := 0; j < 4; j++{
			fmt.Println(mutilArrayValue[i][j])
		}
	}
}
package main

import (
	"fmt"
	"math"
)
func main(){
	var age, tall int
	
	var(
		status = "Single"
		weight = 55
		girl_friend int
	)
	age = 20
	tall = 165
	// short hand declaration
	name, member := "Dang Ngoc Vu", 8
	fmt.Println("My name is:", name, ", Member :", member)  
	fmt.Println("My age is:", age, ", My height is:", tall)
	fmt.Println("My weight is:", weight,", Girl friends: ", girl_friend, ", and I am ", status)
	var me, papa float64
	me = 23
	papa = 45
	result := math.Max(me, papa)
	fmt.Println("the older age is: ", result)
}

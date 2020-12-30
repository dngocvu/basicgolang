package main

import "fmt"

func sum(a int, b int) int{
	return a+b
}
func calcBill(price int, pcs int) int{
	var sumBill int  = price*pcs
	return sumBill
}
// return multi values
func rect(length int, width int) (int, int){
	perimater := (length + width)*2
	area := length*width 	
	return perimater, area
}
// return by name declare
func rect2(length2, width2 int) (perimater2, area2 int){
	perimater2 = (length2 + width2)*2
	area2 = length2*width2 	
	return 
}
func main(){
	var a,b int = 5, 6
	var c,d int = 2, 3
	var e,f int = 20, 40
	var g,h int = 50, 60
	var i,k int = 40, 60

	perimater, area := rect(e,f)
	perimater2, area2 := rect2(g,h)
	// discard one of values return
	perimater3,_:= rect2(i,k)
	fmt.Println("sum = ", sum(a,b))
	fmt.Println("Your bill is ", calcBill(c,d))
	fmt.Println("Property of rectangle perimater:", perimater,"are: ", area)
	fmt.Println("Property of rectangle2 perimater2:", perimater2,"are: ", area2)
	fmt.Println("Property of rectangle2 perimater3:", perimater3)

}
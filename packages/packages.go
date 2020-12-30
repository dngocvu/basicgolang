package main

// using custom packages

import	(
	"fmt"
	"packages/rectangle"
	"packages/triangle"
	"log"
)
// packages level variable
var adge1, adge2, adge3, sin float64 = 5,6,7,30
	var length, width float64 = -50, 60 // negative length
// init function
func init(){
	println("main package initialized")	
	if length < 0 {
		log.Fatal("length is less than zero")
	}
	// exit program
	if width < 0 {
		log.Fatal("width is less than zero")
	}
	// exit program
}
func main(){
	
	a := triangle.AreaTri(adge1, adge2, sin)
	b := triangle.ParimeterTri(adge1, adge2, adge3)
	c := rectangle.AreaRect(length, width)
	d := rectangle.ParimeterRect(length, width)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

}
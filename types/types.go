package main

import (
	"fmt"
	"unsafe"
)

func main(){
	a := true
	b := false
	var c,d,e bool 
	var f,g,h int = 5,6,7 // another way to declare and assign
	c = true
	fmt.Println("a", a, "b", b)
	fmt.Println("c", c)
	d = a && b
	fmt.Println("d", d)
	e = a || b
	fmt.Println("e", e)
	fmt.Println(f, g, h)
	fmt.Printf("type of a is %T, size of a is %d \n", a, unsafe.Sizeof(a))
	// complex type
	c1 := complex(5, 7)
	c2 := 8 + 27i
	sum := c1 + c2
	fmt.Printf("type of sum is %T, size of sum is %d \n", sum, unsafe.Sizeof(sum))
	fmt.Println("sum: ", sum)
	// string 
	var str1,str2, myname string
	str1 = "Dang Ngoc"
	str2 = "Vu"
	myname = str1 +" "+ str2
	fmt.Println("My name is: ", myname)
	// convertion
	var1 := 34
	var2 := 50.5
	sumVar := var1 + int(var2)
	fmt.Println("Sum var = ", sumVar)
	i := 10
	var j float64 = float64(i)
	fmt.Println("j", j)
}
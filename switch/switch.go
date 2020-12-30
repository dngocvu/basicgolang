package main

import "fmt"

func main(){
	// syntax
	finger := 7
	switch finger{
	case 1:
		fmt.Println("Thumb")
	case 2:
		fmt.Println("Index")
	case 3:
		fmt.Println("Middle")
	case 4:
		fmt.Println("Ring")
	case 5:
		fmt.Println("Little")
	default:
		fmt.Println("incorrect finger numberd")
	}
	// mutil expressions
	letter := 'i'
	switch letter{
	case 'a', 'e', 'i', 'o', 'u' :
		fmt.Println("vowel")
	default:
		fmt.Println("Not a vowel")
	}
	fmt.Printf("Type of letter is %T \n", letter)
	// miss expressions
	num := 10
	switch {
	case num < 10 && num > 0:
		fmt.Println("less")
	case num == 10:
		fmt.Println("equal")
	case num > 10:
		fmt.Println("more")
	default:
		fmt.Println("invalue")
	}
	// fallthrough: check the output of a function is ok or not
}
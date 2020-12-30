package main

import "fmt"

func main(){
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}
	// break
	for i := 0; i <= 8; i++ {
		if i == 5 {
			break
		}
		fmt.Println("number:",i)	
	}
	//continue
	for i := 0; i <= 6; i++ {
		if i == 4 {
			continue
		}
		fmt.Println("check:", i)
	}
	// omit initialisation and post
	var i int = 0
	for i < 10 {
		fmt.Println("number_2:", i)
		i += 2
	}
	// mutil initialisation
	for down, up := 10, 1; down > 0 && up < 10; down, up = down - 1, up + 1 {
		sum := down * up
		fmt.Printf("%d * %d = %d \n", down, up, sum)
	}
	// infinite loop
	// don't do it
	/*
	for {
		fmt.Println("1")
	}
	*/
}
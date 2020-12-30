package main 


import "fmt"


func main(){
	// option 1
	var num int = 11
	if num % 2 == 0 {
		fmt.Println("The number is an even number")
	}else{
		fmt.Println("The number is an odd number")
	}
	// option2: num2 is just use in if's scope
	if num2 := 20; num2 % 2 == 0 {
		fmt.Println("The number_2 is an even number")
	}else{
		fmt.Println("The number_2 is an odd number")
	}
	//fmt.Println(num2)
	var age int = 20
	if age < 30 && age > 5 {
		fmt.Println("you are young")
	}else if age > 0 && age <= 5 {
		fmt.Println("you are a baby")
	}else {
		fmt.Println("you are old")
	}

}
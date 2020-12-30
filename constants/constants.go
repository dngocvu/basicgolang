package main

import(
	"fmt"
) 

func main(){
	var num int = 10
	var str string = "Dang Ngoc Vu"
	fmt.Printf("number is: %d, string is: %s \n", num, str)
	// const can not be modified
	const a = 55
	// a = 60
	b := a
	b = 50
	fmt.Println("const: ", b)
	// can not assigned a const variable by a value of a function call
	// const c = math.Sqrt(4)
	const hello = "Hello World"
	fmt.Printf("type of hello is %T \n", hello)
	const typeHello string = "Hello World too"
	fmt.Printf("type of typeHello is %T \n", typeHello)
	// const boolean
	const trueConst = true
	// trueConst = false -> fault
	varBool := trueConst
	fmt.Println("varBool is: ", varBool)
	// key "type" is assign data type

	type myBool bool
	var changedBool myBool = false
	fmt.Println("changed Bool is:", changedBool)

	//numeric constants
	const number = 10
	var intVar int = number
	var int32Var int32 = number
	var int64Var int64 = number
	var float32Var float32 = number
	//var float64Var float64 = number
	//var complex32Var complex64 = number
	
	fmt.Println("int", intVar, "int32", int32Var, "int64", int64Var, "float32", float32Var)
}
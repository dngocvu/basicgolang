package main

import ( 
	"fmt"
	"regexp"
)
func RemoveCharacter(inputRemove string)(outputRemove string){
	reg, _ := regexp.Compile("[^a-zA-Z0-9]+")
	outputRemove = reg.ReplaceAllString(inputRemove, "")
	return
}
func FormatPhoneNumber(inputString string)(result string){
	afterRemoveString := RemoveCharacter(inputString)
	lenString := len(afterRemoveString)
	var count int
	switch lenString % 10 {
	case 1:
		if lenString > 10 {
			for i := 0; i < lenString; i++ {
				fmt.Printf("%s",string(afterRemoveString[i]) )
				count = count + 1
				if count % 3 == 0 && count != 21{
					fmt.Printf("-")
				}
			}
		}
	case 2:
	case 3:
	case 4:
	case 5:
	case 6:
	case 7:
	case 8:
	case 9:
	case 0:
	}
	
	return 
}

func main() {
	var str string = "00-44  48 555555111111111155555555 8"
	fmt.Println(FormatPhoneNumber(str))
}
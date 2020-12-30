package main

import(
	"fmt"
	"strings"
	"strconv"
)

func Swap(strswap_1, strswap_2 *string, len_Str_1 int, len_Str_2 int ) {
	if len_Str_1 > len_Str_2 {
		temp := *strswap_1
		*strswap_1 = *strswap_2
		*strswap_2 = temp
	}
}
func Reverse(str_reserve string) string {
	r := []rune(str_reserve)
    for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
        r[i], r[j] = r[j], r[i]
    }
    return string(r)
}
func SubZeroFormatBefore(str_before_sub string, diff int) string{
	var str_before string = ""
	for i := 0; i < diff; i = i + 1 {
		str_before_sub = "0" + str_before_sub 
	} 
	str_before = str_before_sub
	return str_before	

}
func SubZeroFormatAfter(str_after_sub string, diff int) string{
	var str_after string = ""
	for i := 0; i < diff; i = i + 1 {
		str_after_sub = str_after_sub + "0"
	} 
	str_after = str_after_sub
	return str_after
}
func SubZeroFormatCombination(str_conbination_sub string, diff1 int, diff2 int) string{
	var str_combination string = ""
	for i := 0; i < diff1; i = i + 1 {
		str_conbination_sub = str_conbination_sub + "0"
	} 
	str_combination = str_conbination_sub
	for i := 0; i < diff2; i = i + 1 {
		str_conbination_sub = "0" + str_conbination_sub 
	} 
	str_combination = str_conbination_sub
	return str_combination
}
func RemovePoint(inputRemoveString1 , inputRemoveString2 string)(outputRemoveString1, outputRemoveString2 string){
	outputRemoveString1 = strings.Replace(inputRemoveString1, ".", "", 1)
	outputRemoveString2 = strings.Replace(inputRemoveString2, ".", "", 1)
	return outputRemoveString1, outputRemoveString2
}
func SubPoint(subPointString string, pointPosition int)(result string) {
    if pointPosition != 0{
		result = subPointString[:pointPosition - 1] + "." + subPointString[pointPosition - 1:]
	}else{
		result = subPointString
	}
	return result
}
func CheckPoint(str_point string) (cpoint int) {
	cp := []rune(str_point)
    for i := 0; i < len(cp); i = i + 1 {
        if cp[i] == '.' {
			cpoint = i + 1
		} 
	}
    return cpoint
}
func SubString(subString1, subString2 string)(outputSubString string){
	var bigger bool = false
	for i:= (len(subString2) - 1); i >= 0; i-- {
		if int((subString1[i] - '0')) > int((subString2[i] - '0')){
			bigger = true
			break
		} else if int((subString1[i] - '0')) < int((subString2[i] - '0')){
			bigger = false
			break
		}
	}
	if bigger {
		var carry int = 0
		var subDigital int
		for i := 0; i < len(subString2); i++ {
			subDigital = int(subString1[i] - '0') - int(subString2[i] - '0') - carry
			fmt.Println("subDigital: ",subDigital)
			if subDigital < 0 {
				subDigital = subDigital + 10
				carry = 1
			} else {
				carry = 0
			}
			outputSubString = outputSubString + strconv.Itoa(subDigital)
			fmt.Printf("loop %d: %s \n",i,outputSubString)
		}
	} else {
		var carry int = 0
		var subDigital int
		for i := 0; i < len(subString2); i++ {
			subDigital = int(subString2[i] - '0') - int(subString1[i] - '0') - carry
			fmt.Println("subDigital: ",subDigital)
			if subDigital < 0 {
				subDigital = subDigital + 10
				carry = 1
			} else {
				carry = 0
			}
			outputSubString = outputSubString + strconv.Itoa(subDigital)
			fmt.Printf("loop %d: %s \n",i,outputSubString)
		}
		outputSubString = outputSubString + "-"
	}
	return outputSubString
}
func SubTwoLargeNumBer(str_1 string, str_2 string)(result string){
	//get length of two string\
	var sub, subString string
	var subString1, subString2 string
	var maxPointPosition int
	var str_final_1, str_final_2 string
	len_Str_1 := len(str_1)
	len_Str_2 := len(str_2)
	//get digital position
	cpBefore_1 := CheckPoint(str_1)
	cpBefore_2 := CheckPoint(str_2)
	if cpBefore_1 < cpBefore_2 {
		maxPointPosition = cpBefore_2
	} else {
		maxPointPosition = cpBefore_1
	}
	cpAfter_1 := len_Str_1 - cpBefore_1
	cpAfter_2 := len_Str_2 - cpBefore_2
	//check where digital point and format string
	//to xxxx.xxxx and yyyy.yyyy
	if cpBefore_1 > 0 && cpBefore_2 > 0{
		if (cpAfter_1 < cpAfter_2 && cpBefore_1 < cpBefore_2) {
			diff_After_1 := cpAfter_2 - cpAfter_1
			diff_Before_1 := cpBefore_2 - cpBefore_1
			str_final_1 = SubZeroFormatCombination(str_1,diff_After_1,diff_Before_1)
		}
		if (cpAfter_1 > cpAfter_2 && cpBefore_1 > cpBefore_2) {
			diff_After_2 := cpAfter_1 - cpAfter_2
			diff_Before_2 := cpBefore_1 - cpBefore_2
			str_final_2 = SubZeroFormatCombination(str_2,diff_After_2,diff_Before_2)
		}
		if cpAfter_1 < cpAfter_2 {
			diff_After_1 := cpAfter_2 - cpAfter_1
			str_final_1 = SubZeroFormatAfter(str_1, diff_After_1)
		}
		if cpAfter_1 > cpAfter_2 {
			diff_After_2 := cpAfter_1 - cpAfter_2
			str_final_2 = SubZeroFormatAfter(str_2, diff_After_2)
		}
		if cpBefore_1 < cpBefore_2 {
			diff_Before_1 := cpBefore_2 - cpBefore_1
			str_final_1 = SubZeroFormatBefore(str_1, diff_Before_1)
		}
		if cpBefore_1 > cpBefore_2 {
			diff_Before_2 := cpBefore_1 - cpBefore_2
			str_final_2 = SubZeroFormatBefore(str_2, diff_Before_2)
		}
	} else if cpBefore_1 == 0 && cpBefore_2 == 0 {
		if len(str_1) < len (str_2){
			inDiff := len(str_2) - len (str_1)
			str_final_1 = SubZeroFormatBefore(str_1, inDiff)
			str_final_2 = str_2
		} else if len(str_1) > len (str_2){
			inDiff := len(str_1) - len (str_2)
			str_final_2 = SubZeroFormatBefore(str_2, inDiff)
			str_final_1 = str_1
		} else {
			str_final_1 = str_1
			str_final_2 = str_2
		}
	} else if cpBefore_1 == 0 && cpBefore_2 != 0{
		if cpBefore_2 < len(str_1) {
			inDiff1 := len(str_2) - cpBefore_2
			inDiff2 := len(str_1) + 1 - cpBefore_2
			str_final_1 = SubZeroFormatAfter(str_1, inDiff1)
			str_final_2 = SubZeroFormatBefore(str_2, inDiff2)
			maxPointPosition = len(str_1) + 1
		} else if cpBefore_2 > len(str_1) {
			inDiff1 := cpBefore_2 - len(str_1)
			inDiff2 := cpAfter_2
			str_final_1 = SubZeroFormatCombination(str_1, inDiff1, inDiff2)
			str_final_2 = str_2
			maxPointPosition = cpBefore_2 + 1
		} else {
			str_final_1 = str_1
			str_final_2 = str_2
		}
	} else if cpBefore_1 != 0 && cpBefore_2 == 0 {
		if cpBefore_1 < len(str_2) {
			inDiff1 := len(str_1) - cpBefore_1
			inDiff2 := len(str_2) + 1 - cpBefore_1
			str_final_2 = SubZeroFormatAfter(str_2, inDiff1)
			str_final_1 = SubZeroFormatBefore(str_1, inDiff2)
			maxPointPosition = len(str_2) + 2
		} else if cpBefore_1 > len(str_2) {
			inDiff1 := cpBefore_1 - len(str_2)
			inDiff2 := cpAfter_1
			str_final_2 = SubZeroFormatCombination(str_2, inDiff1, inDiff2)
			str_final_1 = str_1
			maxPointPosition = cpBefore_1 + 2
		} else {
			str_final_1 = str_1
			str_final_2 = str_2
		}

	}
	// final format
	subString1, subString2 = RemovePoint(str_final_1, str_final_2)
	// reverse subString1, subString2
	subString1 = Reverse(subString1)
	subString2 = Reverse(subString2)
	// subString = subString1 + subString2
	subString = SubString(subString1, subString2)
	// reverse subString
	subString = Reverse(subString)
	// add point
	sub = SubPoint(subString, maxPointPosition)
	// final
	return sub
}
func main(){
	//
	var number_1 string = "1223.12123"
	var number_2 string = "14523411231"
	//
	//SubTwoLargeNumBer(number_1, number_2)
	fmt.Println("Sub: ", SubTwoLargeNumBer(number_1, number_2))
	fmt.Println("END")
}

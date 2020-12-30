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
func AddZeroFormatBefore(str_before_add string, diff int) string{
	var str_before string = ""
	for i := 0; i < diff; i = i + 1 {
		str_before_add = "0" + str_before_add 
	} 
	str_before = str_before_add
	return str_before	

}
func AddZeroFormatAfter(str_after_add string, diff int) string{
	var str_after string = ""
	for i := 0; i < diff; i = i + 1 {
		str_after_add = str_after_add + "0"
	} 
	str_after = str_after_add
	return str_after
}
func AddZeroFormatCombination(str_conbination_add string, diff1 int, diff2 int) string{
	var str_combination string = ""
	for i := 0; i < diff1; i = i + 1 {
		str_conbination_add = str_conbination_add + "0"
	} 
	str_combination = str_conbination_add
	for i := 0; i < diff2; i = i + 1 {
		str_conbination_add = "0" + str_conbination_add 
	} 
	str_combination = str_conbination_add
	return str_combination
}
func RemovePoint(inputRemoveString1 , inputRemoveString2 string)(outputRemoveString1, outputRemoveString2 string){
	outputRemoveString1 = strings.Replace(inputRemoveString1, ".", "", 1)
	outputRemoveString2 = strings.Replace(inputRemoveString2, ".", "", 1)
	return outputRemoveString1, outputRemoveString2
}
func AddPoint(addPointString string, pointPosition int)(result string) {
    if pointPosition != 0{
		result = addPointString[:pointPosition - 1] + "." + addPointString[pointPosition - 1:]
	}else{
		result = addPointString
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
func SumString(sumString1, sumString2 string)(sumString string){
	var carry int = 0
	var sumDigital int
	//
	for i := 0; i < len(sumString1); i++ {
		sumDigital = int((sumString1[i] - '0') + (sumString2[i] - '0')) + carry
		carry = sumDigital / 10;
		sumDigital = sumDigital % 10
		sumString = sumString + strconv.Itoa(sumDigital)
	}
	return sumString
}
func SumTwoLargeNumBer(str_1 string, str_2 string)(result string){
	//get length of two string\
	var sum, sumString string
	var sumString1, sumString2 string
	var maxPointPosition int
	var str_final_1, str_final_2 string
	var flag bool = false
	var substringAfter_1, substringAfter_2 string
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
			str_final_1 = AddZeroFormatCombination(str_1,diff_After_1,diff_Before_1)
		}
		if (cpAfter_1 > cpAfter_2 && cpBefore_1 > cpBefore_2) {
			diff_After_2 := cpAfter_1 - cpAfter_2
			diff_Before_2 := cpBefore_1 - cpBefore_2
			str_final_2 = AddZeroFormatCombination(str_2,diff_After_2,diff_Before_2)
		}
		if cpAfter_1 < cpAfter_2 {
			diff_After_1 := cpAfter_2 - cpAfter_1
			str_final_1 = AddZeroFormatAfter(str_1, diff_After_1)
		}
		if cpAfter_1 > cpAfter_2 {
			diff_After_2 := cpAfter_1 - cpAfter_2
			str_final_2 = AddZeroFormatAfter(str_2, diff_After_2)
		}
		if cpBefore_1 < cpBefore_2 {
			diff_Before_1 := cpBefore_2 - cpBefore_1
			str_final_1 = AddZeroFormatBefore(str_1, diff_Before_1)
		}
		if cpBefore_1 > cpBefore_2 {
			diff_Before_2 := cpBefore_1 - cpBefore_2
			str_final_2 = AddZeroFormatBefore(str_2, diff_Before_2)
		}
	} else if cpBefore_1 == 0 && cpBefore_2 == 0 {
		if len(str_1) < len (str_2){
			inDiff := len(str_2) - len (str_1)
			str_final_1 = AddZeroFormatBefore(str_1, inDiff)
			str_final_2 = str_2
		} else if len(str_1) > len (str_2){
			inDiff := len(str_1) - len (str_2)
			str_final_2 = AddZeroFormatBefore(str_2, inDiff)
			str_final_1 = str_1
		} else {
			str_final_1 = str_1
			str_final_2 = str_2
		}
	} else if cpBefore_1 == 0 && cpBefore_2 != 0{
		flag = true
		substringAfter_2 = str_2[cpBefore_2 - 1:len_Str_2]
		str_2 = str_2[0:cpBefore_2 - 1]
		if len(str_1) < len (str_2){
			inDiff := len(str_2) - len (str_1)
			str_final_1 = AddZeroFormatBefore(str_1, inDiff)
			str_final_2 = str_2
		} else if len(str_1) > len (str_2){
			inDiff := len(str_1) - len (str_2)
			str_final_2 = AddZeroFormatBefore(str_2, inDiff)
			str_final_1 = str_1
		} else {
			str_final_1 = str_1
			str_final_2 = str_2
		}
		maxPointPosition = 0
	} else if cpBefore_1 != 0 && cpBefore_2 == 0 {
		flag = true
		substringAfter_1 = str_1[cpBefore_1 - 1:len_Str_1]
		str_1 = str_1[0:cpBefore_1 - 1]
		if len(str_1) < len (str_2){
			inDiff := len(str_2) - len (str_1)
			str_final_1 = AddZeroFormatBefore(str_1, inDiff)
			str_final_2 = str_2
		} else if len(str_1) > len (str_2){
			inDiff := len(str_1) - len (str_2)
			str_final_2 = AddZeroFormatBefore(str_2, inDiff)
			str_final_1 = str_1
		} else {
			str_final_1 = str_1
			str_final_2 = str_2
		}
		maxPointPosition = 0
	}
	// final format
	sumString1, sumString2 = RemovePoint(str_final_1, str_final_2)
	// reverse sumString1, sumString2
	sumString1 = Reverse(sumString1)
	sumString2 = Reverse(sumString2)
	// sumString = sumString1 + sunString2
	sumString = SumString(sumString1, sumString2)
	// reverse sumString
	sumString = Reverse(sumString)
	// add point
	sum = AddPoint(sumString, maxPointPosition)
	// add substring
	if flag {
		sum = sum + substringAfter_1 + substringAfter_2
		flag = false
	}
	// final
	return sum
}
func main(){
	//
	var number_1 string = "14565566666.2341234124"
	var number_2 string = "123456.78900009999999999"
	//
	SumTwoLargeNumBer(number_1, number_2)
	fmt.Println("Sum: ", SumTwoLargeNumBer(number_1, number_2))
	fmt.Println("END")
}
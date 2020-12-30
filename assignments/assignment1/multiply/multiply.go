package main

import(
	"fmt"
	"strings"
	//"strconv"
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
func MtpPoint(mtpPointString string, pointPosition int)(result string) {
    if pointPosition != 0{
		result = mtpPointString[:pointPosition] + "." + mtpPointString[pointPosition:]
	}else{
		result = mtpPointString
	}
	return result
}
func RemovePoint(inputRemoveString1 , inputRemoveString2 string)(outputRemoveString1, outputRemoveString2 string){
	outputRemoveString1 = strings.Replace(inputRemoveString1, ".", "", 1)
	outputRemoveString2 = strings.Replace(inputRemoveString2, ".", "", 1)
	return outputRemoveString1, outputRemoveString2
}
func CheckPoint(str_point string) (cpoint int) {
	cp := []rune(str_point)
    for i := len(cp) - 1; i >= 0; i = i - 1 {
        if cp[i] == '.' {
			cpoint = len(cp) - 1 - i
		} 
	}
    return cpoint
}
func MtpString(mtpString1, mtpString2 string)(outputMtpString string){
	data := make([]byte, len(mtpString1) + len(mtpString2))
	for i := len(mtpString1) - 1; i >= 0; i-- {
		for j := len(mtpString2) - 1; j >= 0; j-- {
			outputMtpString := (mtpString1[i] - '0')*(mtpString2[j] - '0') + data[i+j+1]
			data[i+j], data[i+j+1] = data[i+j] + outputMtpString / 10, outputMtpString % 10
		}
	}
	for i := 0; i < len(data); i++ {
		data[i] += '0'
	}
	for i, outputMtpString := range data {
		if outputMtpString != '0' {
			return string(data[i:])
		}
	}
	return
}
func MtpTwoLargeNumBer(str_1 string, str_2 string)(result string){
	//get length of two string\
	var mtp, mtpString string
	var mtpString1, mtpString2 string
	var maxPointPosition int
	//get digital position
	cpAfter_1 := CheckPoint(str_1)
	cpAfter_2 := CheckPoint(str_2)
	maxPointPosition = cpAfter_1 + cpAfter_2
	//check where digital point and format string
	mtpString1, mtpString2 = RemovePoint(str_1, str_2)
	// mtpString = mtpString1 + mtpString2
	mtpString = MtpString(mtpString1, mtpString2)
	// reverse mtpString
	mtp = Reverse(mtpString)
	// add point
	mtp = MtpPoint(mtp, maxPointPosition)
	mtp = Reverse(mtp)
	// final
	return mtp
}
func main(){
	//
	var number_1 string = "14523.1233"
	var number_2 string = "123.1212"
	//
	//MtpTwoLargeNumBer(number_1, number_2)
	fmt.Println("Multiply: ", MtpTwoLargeNumBer(number_1, number_2))
	fmt.Println("END")
}

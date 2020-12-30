package main


import ( 
	"fmt"
	//"regexp"
	"strings"
	"strconv"
)
func FindSleepingTime(inputString string)(result int){
	var subStringSlice [] string // hh:mm hh:mm
	//var result int
	subStringSlice = CreateSubString(inputString)
	result = CalTime(subStringSlice)
	return
}
// Split , remove day and add into a string follow format: -hh:mm hh:mm
// first element and last element group together
func CreateSubString(inputSubstring string)(subStringArray [] string){
	var strList [] string
	for _, line := range strings.Split(strings.TrimSuffix(inputSubstring, "\n"), "\n") {
		strList = append(strList, line)
	}
	var parseStringOneLine string = ""
	var parseStringOneLineSlice [] string
	for i := 0; i < len(strList); i++ {
		parseStringOneLine = parseStringOneLine + " " + strList[i]
	}
	parseStringOneLineSlice = strings.Fields(parseStringOneLine)
	//
	parseStringSliceNew := make([]string, len(strList))
	var k int = 0
	for i := 1; i < len(parseStringOneLineSlice); i = i + 2 {
		parseStringSliceNew[k] = parseStringOneLineSlice[i]
		k += 1
	}
	return parseStringSliceNew
} 
// convert to integer for calculate
func CalTime(sliceTime [] string)(result int){
	var parseSliceTimeSpace [] string
	var parseSliceTimeSpace_Final [] string
	var parseSliceTimeLine string
	for i := 0; i < len(sliceTime); i++ {
		parseSliceTimeLine = parseSliceTimeLine + "-" + sliceTime[i]
	}
	parseSliceTimeSpace = strings.Split(parseSliceTimeLine, "-")
	parseSliceTimeSpace = parseSliceTimeSpace[1:len(parseSliceTimeSpace)]
	var parseSliceTimeGroup [] string
	parseSliceTimeGroup = append(parseSliceTimeGroup, parseSliceTimeSpace[len(parseSliceTimeSpace)-1] + "-" + parseSliceTimeSpace[0])
	for i := 1; i < len(parseSliceTimeSpace)-1 ; i = i + 2 {
		parseSliceTimeGroup = append(parseSliceTimeGroup, parseSliceTimeSpace[i] + "-" + parseSliceTimeSpace[i+1])
	}
	//
	var parseSliceTimeLine_Final string
	for i := 0; i < len(sliceTime); i++ {
		parseSliceTimeLine_Final = parseSliceTimeLine_Final + "-" + parseSliceTimeGroup[i]
	}
	parseSliceTimeLine_Final = strings.ReplaceAll(parseSliceTimeLine_Final,":","")
	parseSliceTimeSpace_Final = strings.Split(parseSliceTimeLine_Final, "-")
	parseSliceTimeSpace_Final = parseSliceTimeSpace_Final[1:len(parseSliceTimeSpace_Final)]
	//
	var resultArr [] int
	var start, end int
	for i := 0; i < len(parseSliceTimeSpace_Final); i = i + 2{
		start, _ = (strconv.Atoi(parseSliceTimeSpace_Final[i]))
		end, _ = (strconv.Atoi(parseSliceTimeSpace_Final[i+1]))
		resultArr = append(resultArr, 144000 - start*60 + end*60)
	}
	var max int = resultArr[0] 
	for i := 1; i < len(resultArr); i++ {
		if max < resultArr[i] {
			max = resultArr[i]
		}
	}
	return max/100
}

func main () {
	var str string =
`Mon 02:00-23:00
Tue 01:00-17:00
Wed 01:00-23:00
Thu 01:00-23:00
Fri 01:00-23:00
Sat 01:00-23:00
Sun 01:00-21:00`
	
	fmt.Println(FindSleepingTime(str))

}
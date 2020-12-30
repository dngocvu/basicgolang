package main


import ( 
	"fmt"
	//"regexp"
	//"strings"
)

// 1. find length of test case 
// 2. group each group test case together
// 3. count how many ok in group case and single case 
// 4. compute
// 5. printf


func computeScore(inputT []string, inputR []string)(result int) {
	var minT int = 0
	var countGroup int = 0
	var subgroupsLen int = 0
	// map
	var stringMapMain = make(map[string]string)
	var stringMap = make(map[string]string)
	// 1.
	minT = len(inputT[0])
	for i := 1; i < len(inputT); i++ {
		if minT > len(inputT[i]) {
			minT = len(inputT[i])
		}
	}
	// 2.
	for i := 0; i < len(inputT); i++ {
		if minT == len(inputT[i]) {
			countGroup += 1
		}
	}
	subgroupMain := make([]string, countGroup)
	for i := 0; i < len(inputT); i++ {
		if len(inputT[i]) == minT {
			subgroupMain = append(subgroupMain, inputT[i])
			stringMapMain[inputT[i]] = inputR[i]
		}
	}
	subgroupsLen = len(inputT) - countGroup
	subgroup := make([]string, subgroupsLen)
	for i := 0; i < len(inputT); i++ {
		if len(inputT[i]) == (minT + 1){
			subgroup = append(subgroup, inputT[i]) 
			stringMap[inputT[i]] = inputR[i]
		}
	}
	subgroupMain = subgroupMain[(len(subgroupMain)/2):len(subgroupMain)]
	subgroup = subgroup[(len(subgroup)/2):len(subgroup)]
	fmt.Println(subgroup)
	fmt.Println(subgroupMain)
	fmt.Println(stringMapMain)
	fmt.Println(stringMap)
	//
	var countSubgroupMainOK int = 0
	for i := 0; i < len(subgroupMain); i++ {
		if stringMapMain[subgroupMain[i]] == "OK" {
			countSubgroupMainOK += 1
		}
	}
	var diffSign string
	var countSubgroupCase int = 0
	if len(subgroup) > 0 {
		diffSign = string(subgroup[0][minT-1])
		countSubgroupCase = 1
	}
	//var countSubgroupOK int = 0
	arrSubgroup := make([]string, countSubgroupCase)
	arrSubgroup = append(arrSubgroup, diffSign)
	for i := 1; i < len(subgroup); i++ {
		if string(subgroup[i][minT-1]) != diffSign{
			arrSubgroup = append(arrSubgroup, string(subgroup[i][minT-1]))
			countSubgroupCase += 1
		}
	}
	arrSubgroup = arrSubgroup[(len(arrSubgroup)/2):len(arrSubgroup)]
	var countSubgroupNO int = 0
	fmt.Println(arrSubgroup)
	for j := 0; j < len(arrSubgroup); j++ {
		for i := 0; i < len(subgroup); i++ {
			if string(subgroup[i][minT-1]) == arrSubgroup[j] && stringMapMain[subgroup[i]] == "OK" {
				continue
			} else if string(subgroup[i][minT-1]) == arrSubgroup[j] && stringMapMain[subgroup[i]] != "OK"{
				break
			}
			countSubgroupNO -= 1
		}
		countSubgroupNO += 1
	}

	//
	sumCase := countGroup + countSubgroupCase
	fmt.Println(sumCase)
	fmt.Println(countSubgroupMainOK)
	fmt.Println(countSubgroupCase)
	fmt.Println(countSubgroupNO)
	sumCaseOk := countSubgroupMainOK + countSubgroupCase - countSubgroupNO
	fmt.Println(sumCaseOk)
	sum := (sumCaseOk*100)/sumCase
	return sum
}

func main () {
	var strT = [] string {"codility1","codility3","codility2","codility4a","codility4b","codility4c"} 
	var strR = [] string {"OK","OK","OK","OK","Runtime","OK"} 
	fmt.Println(computeScore(strT, strR))

}
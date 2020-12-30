package main


import ( 
	"fmt"
	//"regexp"
	//"strings"
)

// 2. check
// 3. printf

func findChange(inputS string, inputT string)(result string){

	nS := len(inputS)
	nT := len(inputT)
	//var diff int = 0
	var count int
	var same int
	if nS > nT {
		result = "IMPOSSIBLE"
		return result
	}
	if nT > nS + 1{
		result = "IMPOSSIBLE"
		return result
	}
	if nT == nS + 1 {
		for i := 0; i < nS; i++ {
			if string(inputS[i]) ==  string(inputT[i]) {
				result = "ADD " + string(inputT[len(inputT) - 1])
			} else {
				result = "IMPOSSIBLE"
				return result
			}
		}
	}
	if nT == nS {
		for i := 0; i < nS; i++ {
			if string(inputS[i]) !=  string(inputT[i]) {
				count = count + 1
			}
			for j := 0; j < nS; j++ {
				if string(inputS[i]) ==  string(inputT[j]) {
					same = same + 1
				}
			}
		}
		if count == 0 {
			result = "NOTHING"
			return result
		}
		if count == 1 {
			for i := 0; i < nS; i++ {
				if string(inputS[i]) !=  string(inputT[i]) {
					result = "CHANGE " + string(inputS[i]) + " " + string(inputT[i])
					return result
				}
			}
		}
		if count >= 2 && same == nS {
			for i := 0; i < nS; i++ {
				if string(inputS[i]) !=  string(inputT[i]) {
					result = "MOVE " + string(inputS[i])
					return result
				}
			}
		} else {
			result = "IMPOSSIBLE"
			return result
		}
	}
	return
}
	


func main() {

	var strS string = "beans"
	var strT string = "eansb"
	fmt.Println(findChange(strS, strT))

}
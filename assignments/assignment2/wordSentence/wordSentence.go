package main


import ( 
	"fmt"
	//"regexp"
	"strings"
)

// 1. split sentence and add into a []tring
// 2. format sentence to correct format : recursion
// 3. count space
// 4. print
func SplitAny(splitInput string, seps string) []string {
    splitter := func(r rune) bool {
        return strings.ContainsRune(seps, r)
    }
    return strings.FieldsFunc(splitInput, splitter)
}
func wordSentence(inputS string)(result int) {
	//n := len(inputS)
	var countWords int
	var maxWord int
	var containerS [] string
	var containerL [] string
	var containerF [] string
	//var containerN string
	containerS = SplitAny(inputS, ".!?")
	for i := 0; i < len(containerS); i++{
		for j := 1; j < len(containerS[i]); j++{
			// replace = loop
			for k := 0; k < len(containerS[i]); k++{
				if string(containerS[i][j-1]) == " " && string(containerS[i][j]) == " "{
					containerS[i] = strings.ReplaceAll(containerS[i], "  "," ",)
				}
			}
		}
	}
	containerL = containerS
	for i := 0; i < len(containerL); i++{
		if string(containerL[i][len(containerL[i])- 1]) == " "{
			containerL[i] = strings.TrimSuffix(containerL[i], " ") 
		}
	}
	containerF = containerL
	for i := 0; i < len(containerF); i++{
		if string(containerF[i][0]) == " "{
			containerF[i] = strings.TrimPrefix(containerF[i], " ") 
		}
	}
	fmt.Println(containerF)
	arrContainer := make([]int, len(containerF))
	for i := 0; i < len(containerF); i++{
		countWords = 0
		for j := 0; j < len(containerF[i]); j++{
			if string(containerF[i][j]) == " " {
				countWords += 1	
			}
		}
		arrContainer[i] = countWords
		fmt.Println(arrContainer)
	}
	maxWord = arrContainer[0]
	for i := 1; i < len(arrContainer); i++{
		if arrContainer[i] >= maxWord {
			maxWord = arrContainer[i]
		}
	}

	return maxWord + 1
}
	

func main () {
	var strS string = "  We test coders...   Give us a try"
	var strS2 string = "Forget           CVs..Save time .   x x"
	fmt.Println(wordSentence(strS))
	fmt.Println(wordSentence(strS2))

}
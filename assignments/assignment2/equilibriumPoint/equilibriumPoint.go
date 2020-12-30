package main

import "fmt"


func equilibriumPoint(arr [][]int)(result int){
	N := len(arr)
	M := len(arr[1])
	var sumL, sumR, sumU, sumD int
	fmt.Println(N, M)
	var count int = 0
	for i := 1; i < N - 1; i++ {
		for j := 1; j < M - 1; j++ {
			sumL = 0
			sumR = 0
			sumU = 0
			sumD = 0
			//sum left
			for k := 0; k < j; k++{
				for	h := 0; h < N; h++ {
					sumL += arr[h][k]		
				}
			}
			fmt.Println(sumL)
			//sum right
			for k := j+1; k < M; k++{
				for	h := 0; h < N; h++ {	
					sumR += arr[h][k]
				}
			}
			fmt.Println(sumR)
			//sum up
			for k := 0; k < i; k++{
				for	h := 0; h < M; h++ {
					sumU += arr[k][h]		
				}
			}
			fmt.Println(sumU)
			//sum down
			for k := i+1; k < N; k++ {
				for	h := 0; h < M; h++ {	
					sumD += arr[k][h]
				}
			}
			fmt.Println(sumD)
			if sumU == sumD && sumL == sumR {
				count += 1
			} 
		}	
	}
	return count
}

func main(){

	var arr = [][] int {{2, 7, 5},
						{3, 1, 1},
						{2, 1, -7},
						{0, 2, 1},
						{1, 6, 8},
						{-1, 2, -1}}
	fmt.Println(equilibriumPoint(arr))

}
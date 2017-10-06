package main

import (
	"fmt"
)

func max(x, y int) (z int)  {
	z = x
	if y > x {
		z = y
	}
	return z
}

func cuttingRod(N *int) int {

	pieces := make([]int, *N)
	for i:=0; i<*N; i++ {
		fmt.Scan(&pieces[i])
	}

	fmt.Println("Pieces ", pieces)

	dp := make([]int, *N+1)
	dp[0] = 0
	for i:=1; i<=*N; i++ {
		for j:=0; j<i; j++ {
			dp[i] = max(dp[i], dp[j] + pieces[i-j-1])
		}
	}

	return dp[*N]
}

func main()  {

	var T int
	fmt.Print("enter test cases \n")
	fmt.Scan(&T)

	for ; T > 0; T-- {
		N := 0
		fmt.Print("enter rod length \n")
		fmt.Scan(&N)

		fmt.Println("cutting rod :", cuttingRod(&N))
	}
}
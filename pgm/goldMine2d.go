package main

import (
	"fmt"
)

func max(x, y uint8) uint8 {
	if y > x {
		return y
	}
	return x
}

func maxGoldMine(a [][]uint8, m, n int) uint8 {

	var sum uint8

	dp := make([][] uint8, m)
	for i := range a {
		dp[i] = make([]uint8, n)
	}

	for i:=0;i<n;i++ {
		for j:=0;j<m;j++ {
			dp[j][i] = a[j][i]

			if i-1 >= 0 {
				if dp[j][i] < dp[j][i-1] + a[j][i] {
					dp[j][i] = dp[j][i-1] + a[j][i]
				}
				if j-1 >=0 && dp[j][i]  < dp[j-1][i-1] + a[j][i] {
					dp[j][i] = dp[j-1][i-1] + a[j][i]
				}
				if j+1 < m && dp[j][i]  < dp[j+1][i-1] + a[j][i] {
					dp[j][i] = dp[j+1][i-1] + a[j][i]
				}
			}

			if sum < dp[j][i] {
				sum = dp[j][i]
			}

		}
	}

	return sum
}


func main()  {

	var T int
	fmt.Print("enter test cases \n")
	fmt.Scan(&T)

	for ; T > 0;  {
		T--
		var m,n int
		fmt.Print("enter array dimension length \n")
		fmt.Scan(&m, &n)

		a := make([][]uint8, m)
		for i := range a {
			a[i] = make([]uint8, n)
		}

		fmt.Print("enter matrix element value \n")
		for i:=0; i<m; i++ {
			for j:=0; j<n; j++ {
				fmt.Scan(&a[i][j])
			}
		}

		fmt.Println("Max gold collected by person :", maxGoldMine(a, m, n))
	}
}
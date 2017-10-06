package main

import (
	"fmt"
	"math"
)

func printSubsets(s string)  {

	length := len(s)
	totalSubset := math.Pow(2, float64(length)) - 1

	for i := 1; i <= int(totalSubset); i++ {

		for j := uint(0); j<uint(length); j++ {
			if ok := i & (1<<j); ok > 0 {
				fmt.Print(string(s[j]))
			}
		}
		fmt.Println()
	}
}

func main()  {

	var T int
	fmt.Print("Enter test cases \n")
	fmt.Scan(&T)

	for ; T > 0; T-- {
		var s string

		fmt.Scan(&s)

		//fmt.Println(string(s[0]), s[1])
		printSubsets(s)
	}
}
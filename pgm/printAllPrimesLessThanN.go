package main

import "fmt"

func main()  {

	var T int
	fmt.Scan(&T)

	for ;T > 0; {

		var N int
		fmt.Scan(&N)

		fmt.Printf("1")

		var arr = make([]bool, N+1)
		for i:=2;i<=N;i++ {
			if arr[i] == false {
				fmt.Printf(", %v", i)
				for p := i; p<=N; p+=i {
					arr[p] = true
				}
			}
		}
		T--
		fmt.Println()
	}

}

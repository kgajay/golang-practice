package main

import "fmt"

func minCoinChange (arr []int, N, target int) int {

	dp := make([]int, target+1)
	dp[0] = 1

	for i:=0; i<N; i++ {
		for j:=arr[i]; j<=target; j++ {
			dp[j] = dp[j] + dp[j-arr[i]]
		}
	}
	//fmt.Println(dp)
	return dp[target]
}

func main()  {

	var T int
	_, err := fmt.Scanf("%d", &T)

	fmt.Println(T, err)

	for i :=0 ; i<T; i++ {
		var N int
		_, err := fmt.Scanf("%d", &N)
		fmt.Println(N, err)
		//var arr []int
		arr := make([]int, N)
		for i2:=0 ; i2<N; i2++ {
			//_, err := fmt.Scanf("%d", &arr[i2])
			//fmt.Println(err)
			fmt.Scan(&arr[i2])
		}

		fmt.Println(arr)

		var target int
		fmt.Scan(&target)
		//fmt.Println(target)

		fmt.Println(minCoinChange(arr, N, target))
	}
}
package main

import "fmt"

func addElemInMap(x, low, high int, freq map[int]int)  {

	_, ok := freq[x]
	if ok {
		freq[x] += high - low + 1
	} else {
		freq[x] = high - low + 1
	}

}

func binarySearchCnt(arr []int, freq map[int]int, low, high int)  {

	if low <= high {

		if arr[low] == arr[high] {
			addElemInMap(arr[low], low, high, freq)
		} else {
			mid := (low + high) / 2;
			binarySearchCnt(arr, freq, low, mid)
			binarySearchCnt(arr, freq, mid+1, high)
		}
	}
}

func findFrequency(arr []int, freq map[int]int, n int)  {

	binarySearchCnt(arr, freq, 0, n)
}

func main()  {

	var T int
	fmt.Print("Enter test cases \n")
	fmt.Scan(&T)

	for ; T > 0; T-- {
		var N int
		fmt.Scan(&N)
		arr := make([]int, N)
		for i2:=0 ; i2<N; i2++ {
			//_, err := fmt.Scanf("%d", &arr[i2])
			//fmt.Println(err)
			fmt.Scan(&arr[i2])
		}

		freq := make(map[int]int)

		findFrequency(arr, freq, N-1)

		for key, value := range freq {
			fmt.Println("Key:", key, "Value:", value)
		}
	}
}
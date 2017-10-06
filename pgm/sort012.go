package main

import "fmt"

func swap(a, b int) (int, int) {
	return b, a
}

func sort012(a []int, n int)  {

	var low, mid, high int
	high = n-1

	for ;mid <= high; {
		switch a[mid] {
		case 0:
			a[low], a[mid] = swap(a[low], a[mid])
			low++
			mid++
			break
		case 1:
			mid++
			break
		case 2:
			a[mid], a[high] = swap(a[mid], a[high])
			high--
			break
		}
	}
}

func main()  {

	var T int
	fmt.Print("Enter test cases \n")
	fmt.Scan(&T)

	for ; T > 0; T-- {
		var n int
		fmt.Scan(&n)

		a := make([]int, n, n)
		for i:=0; i<n; i++ {
			fmt.Scan(&a[i])
		}
		sort012(a, n)

		for i:=0; i<n; i++ {
			fmt.Printf("%d ", a[i])
		}
		fmt.Println()
	}
}
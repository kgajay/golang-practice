package main

import "fmt"

func fibTR(n, i, a , b uint16) uint16 {

	if n <= i {
		return a+b
	} else {
		temp := a+b
		a = b
		b = temp
		fmt.Print(b, " ")
		return fibTR(n, i+1, a, b)
	}
}

func fib(n uint16) uint16 {

	return fibTR(n, 1, 0, 1)
}

func main()  {

	var T uint16
	fmt.Print("Enter test cases \n")
	fmt.Scan(&T)

	for ; T > 0; T-- {
		var n uint16
		fmt.Print("Enter fibonacci number \n")
		fmt.Scan(&n)

		fmt.Println(fib(n))
	}

}
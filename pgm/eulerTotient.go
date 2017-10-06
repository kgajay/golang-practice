package main

import "fmt"

func phi(n int) int {
	var result float64
	result = float64(n)
	for p:=2; p*p<=n; p++ {
		if n%p == 0 {
			for ; n%p == 0; {
				n = n/p;
			}
			result *= (1.0 - (1.0 / (float64(p))));
		}
	}

	if n > 1 {
		result *= (1.0 - (1.0 / (float64(n))));
	}

	return int(result)
}

func phi2(n int) int {
	var result = n;   // Initialize result as n

	// Consider all prime factors of n and subtract their
	// multiples from result
	for p:=2; p*p<=n; p++ {
		// Check if p is a prime factor.
		if n % p == 0 {
			// If yes, then update n and result
			for ;n % p == 0; {
				n /= p;
			}
			result -= result / p;
		}
	}

	// If n has a prime factor greater than sqrt(n)
	// (There can be at-most one such prime factor)
	if n > 1 {
		result -= result / n;
	}
	return result;
}

func main()  {

	var n int
	fmt.Print("Enter n : ")
	fmt.Scan(&n)

	for i := 1; i<=n; i++ {
		fmt.Printf("phi(%d) = %d  phi2(%d) = %d\n", i, phi(i), i, phi2(i))
	}
}
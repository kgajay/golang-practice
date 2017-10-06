package main

import ("fmt"
        "time"
        "math/rand"
        "math/cmplx"
)
import (
	"math"
	"runtime"
)

func add(x int, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9;
	y = sum - x;
	return
}

var (
	ToBe bool = false
	MaxInt uint64 = 1<<64 -1
	z complex128 = cmplx.Sqrt(-5 + 12i)
)

var c, python, java bool
var d int
var i, j int = 1, 2
const Pi = 3.14

const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func mySqrt(x float64) string {
	if (x < 0) {
		return mySqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v;
	}
	return lim
}

type Vertex struct {
	X int
	Y int
}

func main() {
    fmt.Printf("hello, world\n")
    fmt.Println("This time is ", time.Now(), time.Microsecond)
	fmt.Println("My favorite number is", rand.Intn(20))
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))

	fmt.Println(add(10, 12))

	a, b := swap("hello", "world")

	fmt.Println(a, b)
	fmt.Println(split(17))

	var i int
	fmt.Println(c, python, java, d, i, j)

	var c, d int = 1, 2
	e := 3
	f, g, h := false, true, "No!"
	fmt.Println(c, d, e, f, g, h)


	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	var a1, b1 uint = 3, 4
	var f1 float64 = math.Sqrt(float64(a1*a1 + b1*b1))
	var z1 int = int(f1)
	fmt.Printf("a1: %T, b1: %T, f1: %T, z1: %T\n", a1, b1, f1, z1)

	const Truth = true
	fmt.Println("Go Rules?", Truth, "; Pi:", Pi)

	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))

	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	sum = 1
	for ; sum < 100 ; {
		sum += 10
	}
	fmt.Println(sum)

	sum = 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)

	fmt.Println(mySqrt(2), mySqrt(-4))

	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.", os)
	}

	today := time.Now().Weekday()
	fmt.Printf("today %v\n", today)
	fmt.Println("When's Saturday?")
	switch time.Saturday {
	case today:
		fmt.Println("Today")
	case today + 1:
		fmt.Println("Tomorrow")
	case today + 2:
		fmt.Println("In two days")
	default:
		fmt.Println("Too far away")
	}

	// defer fmt.Println("world")

	fmt.Println("hello")

	i, j = 42, 2012
	p := &i
	*p = 21
	fmt.Println(i)

	p = &j
	*p = *p/37
	fmt.Println(j)

	fmt.Println(Vertex{1, 2})
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v, v.X)

	var aStr [2]string
	aStr[0] = "lo"
	aStr[1] = "bo"
	fmt.Println(aStr[1], aStr[0], aStr)

	primes := [5]int{2, 3, 5, 7, 11}
	fmt.Println(primes)

	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s)

	// Extend its length.
	s = s[:4]
	printSlice(s)

	// Drop its first two values.
	s = s[2:]
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
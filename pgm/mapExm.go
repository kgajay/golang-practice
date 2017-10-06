package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

var m2 = map[string]Vertex{
	"Bell Labs": Vertex{
		40.68433, -74.39967,
	},
	"Google": Vertex{
		37.42202, -122.08408,
	},
}


func (v Vertex) Abs() float64 {
	return math.Sqrt(v.Lat*v.Lat + v.Long*v.Long)
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Abser interface {
	Abs() float64
}

func (v *Vertex) Scale(f float64)  {
	v.Lat = v.Lat * f
	v.Long = v.Long * f

}
//func (v *Vertex) Abs() float64 {
//	return math.Sqrt(v.Lat*v.Lat + v.Long*v.Long)
//}

type I interface {
	M()
}

func describe(i I)  {
	fmt.Printf("(%v, %T)\n", i, i)
}

type T struct {
	S string
}

func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

type F float64

func (f F) M() {
	fmt.Println(f)
}

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}


func main()  {
	m = make(map[string]Vertex)
	fmt.Println(m)
	m["Bell labs"] = Vertex{10, 34.6888}
	fmt.Println(m)

	fmt.Println(m2)

	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)

	m["abcd"] = 12
	v, ok = m["abcd"]
	fmt.Println("The value:", v, "Present?", ok)

	x := Vertex{3, 4}
	fmt.Println(x.Abs())

	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())

	x.Scale(3)
	fmt.Println(x)

	var a Abser
	f = MyFloat(-math.Sqrt2)
	x = Vertex{3, 4}

	a = f  // a MyFloat implements Abser
	fmt.Println(a.Abs())
	a = x
	fmt.Println(a.Abs())

	var i I
	i = &T{"Hello"}
	describe(i)
	i.M()

	i = F(math.Pi)
	describe(i)
	i.M()

	var t *T
	i = t
	describe(i)
	i.M()

	do(21)
	do("hello")
	do(true)

}
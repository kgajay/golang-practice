package main

import (
	"fmt"
	"math/rand"
)

type Tree struct {
	Value int
	Left *Tree
	Right *Tree
}

// New returns a new, random binary tree
// holding the values 1k, 2k, ..., nk.
func New(n, k int) *Tree {
	var t *Tree
	for _, v := range rand.Perm(n) {
		t = insert(t, (1+v)*k)
	}
	return t
}



func insert(root *Tree, v int)  *Tree {

	if root == nil {
		return &Tree{v, nil, nil}
	}
	if v < root.Value {
		root.Left = insert(root.Left, v)
		return root
	}
	root.Right = insert(root.Right, v)
	return root
}

func Walk(t *Tree, ch chan int)  {

	if t == nil {
		return
	}
	Walk(t.Left, ch)
	ch <- t.Value
	Walk(t.Right, ch)
}

func Walker(t *Tree) <-chan int  {

	ch := make(chan int)
	go func() {
		Walk(t, ch)
		close(ch)
	}()
	return ch
}

func Compare(t1, t2 *Tree) bool {

	c1, c2 := Walker(t1), Walker(t2)
	for {
		v1, ok1 := <-c1
		v2, ok2 := <-c2
		if !ok1 || !ok2 {
			return ok1 == ok2
		}
		if v1 != v2 {
			break
		}
	}
	return false
}


func main()  {

	t1 := New(100, 1)
	ch := Walker(t1)
	for  i := range ch {
		fmt.Printf("%v ", i)
	}
	fmt.Println()


	fmt.Println(Compare(t1, New(100, 1)), "Same Contents")
	fmt.Println(Compare(t1, New(99, 1)), "Differing Sizes")
	fmt.Println(Compare(t1, New(100, 2)), "Differing Values")
	fmt.Println(Compare(t1, New(101, 2)), "Dissimilar")
}
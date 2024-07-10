package main

import "fmt"

func main() {
	a := 5
	fmt.Printf("%v %v\n", a, &a)
	pointerA := &a
	fmt.Printf("%v %v %T", pointerA, *pointerA, pointerA)
}

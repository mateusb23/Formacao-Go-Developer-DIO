package main

import "fmt"

var x int
var y float64

func main() {
	x := 10
	y := 10.6
	fmt.Printf("x: %d,%T e y: %g,%T", x, x, y, y)
}

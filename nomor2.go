package main

import "fmt"

func main() {
	var a, b, c, x, y, z float64
	a = 50
	b = 30
	c = 20

	x = a + b - c
	y = a * b
	z = a / b

	fmt.Println("nilai x = ", x)
	fmt.Println("nilai y = ", y)
	fmt.Println("nilai z = ", z)
}
package main

import "fmt"

var i int = 5

func main() {
	greet()
	fmt.Println(wish("santosh"))
	a, b := div(11, 5)
	fmt.Println(a, b)
	fmt.Println(addandsub(11, 5))
	increment()
	decrement()
}

// Declaring func
func greet() {
	fmt.Println("hello")
}

// func parameters with single return value
func wish(name string) string {
	return "happy birthday" + name
}

// multiple func parameters with multiple return value
func addandsub(a, b int) (int, int) {
	return a + b, a - b
}

// Named return values + naked return
func div(a, b int) (quotient, remainder int) {
	quotient = a / b
	remainder = a % b
	return
}

// shadowing
func increment() {
	fmt.Println("before", i)
	{
		i := 20
		fmt.Println(i)
	}
	fmt.Println("before", i)
}

// mutated global
func decrement() {
	fmt.Println("before", i)
	{
		i = 20
		fmt.Println(i)
	}
	fmt.Println("before", i)
}

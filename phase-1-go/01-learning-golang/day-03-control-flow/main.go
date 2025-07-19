package main

import (
	"fmt"
)

func main() {
	for {

		var n int
		fmt.Printf("Choose one:\n1. Run FizzBuzz\n2. Run Grade\n3. Login\n4. Labeled loop test\n5. While-style loop\n6. dispatcher\n")
		fmt.Scanf("%d", &n)
		switch n {
		case 1:
			fizzbuzz()
		case 2:
			grade()
		case 3:
			login()
		case 4:
			labeledLoop()
		case 5:
			whileStyle()
		case 6:
			dispatcher()
		default:
			fmt.Println("invalid number")
		}
	}

}

func fizzbuzz() {
	var n int
	fmt.Printf("Enter a number for fizzbuzz: ")
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {

		switch {
		case i%15 == 0:
			fmt.Printf("FizzBuzz\n")
			// continue
		case i%3 == 0:
			fmt.Println("Fizz")
			// continue
		case i%5 == 0:
			fmt.Println("Buzz")
			// continue
		default:
			fmt.Println(i)
		}
	}

}

func whileStyle() {
	var n int
	fmt.Printf("Enter number to display from 1: ")
	fmt.Scanf("%d", &n)
	for i := 0; i < n; i++ {
		fmt.Println(i)
	}
}

func dispatcher() {
	var n int
	fmt.Println("Enter an integer for dispatcher, either 1 or 2.")
	fmt.Scanf("%d", &n)

	switch n {
	case 1:
		grade()
	case 2:
		login()
	default:
		fmt.Println("exiting.....")
		return
	}
}

func grade() {
	var marks int
	fmt.Printf("Enter your marks: ")
	fmt.Scanf("%d", &marks)
	switch {
	case marks >= 90:
		fmt.Println("A")
	case marks >= 80:
		fmt.Println("B")
	case marks >= 70:
		fmt.Println("C")
	case marks < 70:
		fmt.Println("F")
	default:
		fmt.Println("invalid number")
	}
}

func login() {
	var password, username string

	fmt.Printf("Enter username and password: ")
	fmt.Scanf("%s%s", &username, &password)
	if password != "admin123" {
		fmt.Println("Access for " + username + " is denied")
		return
	}
	fmt.Println("Access for " + username + " is granted")
}

func labeledLoop() {
outer:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("Checking i=%d, j=%d (i*j=%d)\n", i, j, i*j)
			if i*j == 4 {
				fmt.Println("Breaking outer loop because i*j == 4")
				break outer
			}
		}
	}
	fmt.Println("Exited labeled loop")
}

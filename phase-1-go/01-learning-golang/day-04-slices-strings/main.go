package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	appendGrowth()
	copySlice()
	arraydemo()
	stringMutate()
}

func arraydemo() {
	num := [5]int{7, 8, 9, 4, 5}
	slice := num[1:4]
	fmt.Println(num)
	fmt.Println(slice)
	slice[0] = 999
	fmt.Println(num)
	fmt.Println(slice)
}

func appendGrowth() {
	nums := make([]int, 0)
	for i := 0; i < 5; i++ {
		nums = append(nums, i)
		fmt.Printf("Length: %d, Capacity: %d\n", len(nums), cap(nums))
	}
}

func copySlice() {
	a := []int{1, 2, 3, 4, 5}
	b := make([]int, len(a))
	copy(b, a)
	b[0] = 99
	fmt.Println("Original slice:", a)
	fmt.Println("Copied slice:", b)
}

func stringMutate() {
	name := "john"
	for i, v := range name {
		fmt.Printf("Index %d: %c\n", i, v)
	}

	nameBytes := []byte(name)
	nameBytes[0] = 'd'
	updatedName := string(nameBytes)
	fmt.Println(updatedName)

	greet := "नमस्ते"
	fmt.Println("Length in bytes:", len(greet))
	fmt.Println("Length in runes:", utf8.RuneCountInString(greet))
	for i, v := range greet {
		fmt.Printf("Index %d: %c\n", i, v)
	}
}

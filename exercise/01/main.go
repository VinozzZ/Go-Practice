package main

import (
	"fmt"
)

func greatestNum(nums ...int) int {
	var max int
	for _, num := range nums {
		if max < num {
			max = num
		}
	}
	return max
}

func foo(num ...int) {
	for _, n := range num {
		fmt.Println("hey", n)
	}
}

func main() {
	half := func(num int) (int, bool) {
		halfNum := num / 2
		isEven := num%2 == 0
		// result := make([]string, 2)
		// result[0] = string(halfNum)
		// result[1] = isEven
		return halfNum, isEven
	}
	num, even := half(4)
	fmt.Println(num, even)

	randomNums := []int{1, 2, 3, 4, 5, 7}
	max := greatestNum(randomNums...)
	fmt.Println(max)

	foo(1, 2)
	foo(1, 2, 3)
	aSlice := []int{1, 2, 3, 4}
	foo(aSlice...)
	foo()
}

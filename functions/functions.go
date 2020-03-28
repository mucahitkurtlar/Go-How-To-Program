package main

import "fmt"

func myFunction()(a, b int) {
	a = 13
	b = 42
	return
}

func variadic(words ...string) {
	for i, w := range words {
		fmt.Println(i, w)
	}
}

func factorial(num uint)(uint) {
	if num == 0 {
		return 1
	}
	return num * factorial(num - 1)
}

func main() {
//	var x int
//	var y int
	closure := func(num int) int {
		return -num
	}
	x, _ := myFunction()
	fmt.Printf("%d\n\n", x)
	variadic("a", "b", "c")
	fmt.Println(closure(5))
	fmt.Println(factorial(5))
}

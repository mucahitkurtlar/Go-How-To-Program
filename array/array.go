package main

import "fmt"

func main() {
	var arr0 [3]int
	for i := 0; i < 3; i++ {
		arr0[i] = i
	}
	fmt.Println(arr0)

	arr1 := []string{"a", "b", "c", "d", "e"}
	fmt.Println(arr1)
	var arr2 []string = arr1[1:3] //slice işlemi matematikteki [x, y) gibi gerçekleşiyor
	fmt.Println(arr2)
	fmt.Println(cap(arr2))
	var arr3 []int
	if arr3 == nil {
		fmt.Println("Nil") //Array'lerin boş değeri nil'dir
	}
}

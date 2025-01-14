package main

import "fmt"

func foo() (int, int) {
	return 3, 7
}

func main() {
	a, b := foo()
	fmt.Println(a)
	fmt.Println(b)

	/**
	특정 반환 값만 원할 경우 _ 키워드 사용
	*/
	_, c := foo()
	fmt.Println(c)

	d, _ := foo()
	fmt.Println(d)
}

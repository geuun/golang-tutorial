package _pointer

import "fmt"

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0
}

func Main() {
	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	fmt.Println("pointer:", &i)

	var a int = 10
	var b int = 20

	var intPointer1 *int = &a
	var intPointer2 *int = &b

	fmt.Println("a:", intPointer1)
	fmt.Println("b:", intPointer2)
}

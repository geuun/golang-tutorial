package _closures

import "fmt"

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func Main() {
	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	// 3

	newInt2 := intSeq()
	fmt.Println(newInt2())
	fmt.Println(newInt2())
	fmt.Println(newInt2())
	fmt.Println(newInt2())
	//4

	fmt.Println(nextInt())
	// 4
	fmt.Println(newInt2())
	// 5
}

package main

import (
	"fmt"
	"slices"
)

func main() {
	/**
	Slice는 동적 타입의 배열
	컴파일 시점에 배열의 크기가 정해진다.
	*/

	var s []string
	fmt.Println("uninit:", s, s == nil, len(s) == 0)
	// uninit: [] true true

	s = make([]string, 3)
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))
	// emp: [   ] len: 3 cap: 3

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])
	fmt.Println("len:", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("append:", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("copy:", c)

	l := s[2:5]
	fmt.Println("s[2:5]:", s[2:5])
	fmt.Println("s[2:5]:", l)

	l = s[:5]
	fmt.Println("s[:5]:", s[:5])
	fmt.Println("s[:5]:", l)

	l = s[2:]
	fmt.Println("s[2:]:", s[2:])
	fmt.Println("s[2:]", l)

	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	t2 := []string{"g", "h", "i"}
	if slices.Equal(t, t2) {
		fmt.Println(" t === t2")
	}

	twoD := make([][]int, 3)
	fmt.Println("dcl:", twoD)
	// [[] [] []]
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			fmt.Println("j", j)
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d:", twoD)
}

package _slice

import (
	"fmt"
	"slices"
)

func Main() {
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

	/**
	make(슬라이스 타입, 슬라이스 길이, 슬라이스 용량)
	*/
	slice := make([]int, 5, 10)
	fmt.Println("slice:", slice)
	fmt.Println("len:", len(slice))
	fmt.Println("cap:", cap(slice))
	for i := 0; i < 5; i++ {
		slice[i] = i
	}
	fmt.Println("slice:", slice)

	slice = make([]int, 0, 5)
	for i := 0; i < 11; i++ {
		slice = append(slice, i)
		fmt.Println(cap(slice), len(slice), slice)
	}

	/**
	[1/5] 	5 1 [0]
	[2/5]	5 2 [0 1]
	[3/5]	5 3 [0 1 2]
	[4/5]	5 4 [0 1 2 3]
	[5/5]	5 5 [0 1 2 3 4]
	[6/10]	10 6 [0 1 2 3 4 5]  >> capacity 두배로 증가
	[7/10]	10 7 [0 1 2 3 4 5 6]
	[8/10]	10 8 [0 1 2 3 4 5 6 7]
	[9/10]	10 9 [0 1 2 3 4 5 6 7 8]
	[10/10]	10 10 [0 1 2 3 4 5 6 7 8 9]
	[11/20]	20 11 [0 1 2 3 4 5 6 7 8 9 10] >> capacity 두배로 증가
	*/
}

package main

import (
	"fmt"
	"math"
)

const s string = "constant"

func main() {
	fmt.Println("s = ", s)

	const n = 500000000

	const d = 3e20 / n
	fmt.Println("int64(d) =", int64(d))

	fmt.Println(math.Sin(n))
}

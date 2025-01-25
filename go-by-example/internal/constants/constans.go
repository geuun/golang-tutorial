package _constants

import (
	"fmt"
	"math"
)

const s string = "constants"

func Main() {
	fmt.Println("s = ", s)

	const n = 500000000

	const d = 3e20 / n
	fmt.Println("int64(d) =", int64(d))

	fmt.Println(math.Sin(n))
}

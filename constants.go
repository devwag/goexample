package main

import (
	"fmt"
	"math"
)

const s string = "const s"

func main() {
	fmt.Println(s)

	const n = 5000

	const d = 3e20 / n

	fmt.Println(d)

	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))
}

package main

import "fmt"

func main() {
	var a [10]int

	fmt.Println("emp", a)

	a[4] = 100
	fmt.Println("val4:", a[4])
	fmt.Println("set", a)

	b := [5]int{0, 1, 2, 3}
	fmt.Println("b:", b)

	var twoD [2][3]int

	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d", twoD)
}

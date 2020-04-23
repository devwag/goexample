package main

import "fmt"

func main() {

	s := make([]string, 3)

	s[0] = "d"
	s[1] = "e"
	s[2] = "s"
	s = append(s, "k")

	s = append(s, "top")

	fmt.Println("original slice", s)

	c := s[2:4]

	fmt.Println("slice slice1", c)

	c = s[0:4]

	fmt.Println("slice slice2", c)

	c = s[1:4]

	fmt.Println("slice slice3", c)

	t := []string{"g", "h", "i"}

	t = append(t, "k")

	fmt.Println("init slice", t)

	twoD := make([][]int, 3)

	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)

		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}

	}

	fmt.Println("2d: ", twoD)

}

package main

import "fmt"

func main() {

	mps := make(map[string]int)

	mps["one"] = 123
	mps["two"] = 456
	mps["three"] = 789

	fmt.Println("maps: ", mps)
	fmt.Println("len: ", len(mps))
	fmt.Println("val ", mps["two"])

	delete(mps, "two")
	delete(mps, "2")

	i, prs := mps["two"]
	fmt.Println("prs: ", prs)
	fmt.Println("prs: ", i)

	k, prs3 := mps["three"]
	if prs3 {
		fmt.Println("prs3: val", k)

	}

}

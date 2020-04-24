package main

import "fmt"

func main() {

	arr := []int{0, 1, 2, 3, 4}

	for _, n := range arr {
		fmt.Println(n)
	}

	for i, n := range arr {

		if n == 3 {
			fmt.Println(i, n)
		}
	}

	kvs := map[string]string{"a": "apple", "b": "banana", "c": "cucumber"}

	for k, v := range kvs {
		fmt.Println(k, v)
	}

	for i, c := range "bello" {
		fmt.Println(i, c)
	}
}

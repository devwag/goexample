package main

import (
	"fmt"
	"time"
)

func main() {

	i := 4
	fmt.Print("Write ", i, " as ")

	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Println("default")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("its weekend")
	default:
		fmt.Println("its weekday")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I am a bool")

		case int:
			fmt.Println("I am an int")

		case string:
			fmt.Println("I am an string")

		default:
			fmt.Printf("Dont know type %T\n", t)
		}
	}

	whatAmI(true)
	whatAmI(1)
	whatAmI("Hey")
	whatAmI(123e5)
	whatAmI(3.5)

}

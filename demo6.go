package main

import (
	"fmt"
	"time"
)

func main() {
	a := 2
	switch a {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	case 3:
		fmt.Println("3")
	case 4, 5:
		fmt.Println("4 or 5")
	default:
		fmt.Println("other")
	}
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It is before noon")
	default:
		fmt.Println("It is after noon")
	}
}

package main

import "fmt"

func add3(n int) {
	n += 2

}
func add2ptr(n *int) {
	*n += 2
}
func main() {
	n := 5
	add3(n)
	fmt.Println(n)
	add2ptr(&n)
	fmt.Println(n)
}

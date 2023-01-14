package main

import "fmt"

func main() {
	nums := []int{2, 2, 3, 4}
	sum := 0
	for i, num := range nums {
		sum += num
		if num == 2 {
			fmt.Println("index:", i, "num:", num)
		}
	}
	fmt.Println("sum:", sum)
	m := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range m {
		fmt.Println(k, v)
	}
	for k := range m {
		fmt.Println("key", k)
	}
}

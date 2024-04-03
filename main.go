package main

import "fmt"

func main() {
	var sum int
	for i := 0; i < 5; i++ {
		sum += i
	}
	fmt.Printf("sum is %d\n", sum)
}

package main

import "fmt"

func main() {
	var arr []int = []int{1, 2, 3, 6, 7}
	for i, val := range arr {
		fmt.Println(i, val)
	}
}

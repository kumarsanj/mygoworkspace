package main

import "fmt"

func sum(i, j int) int {
	return i + j
}

func main() {
	i := 0
	j := 2
	i = i + 1
	fmt.Println("hello world 3")
	fmt.Println(i)
	k := sum(i, j)
	fmt.Println(k)
}

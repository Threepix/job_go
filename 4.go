package main

import "fmt"

func posled(n int) int {
	var b, c, res, i int
	b = 3
	c = 3
	res = 0
	for i = 0; i < 100000000; i++ {
		res += b*2 + c - (n - 1)
	}
	return posled(n-1) + res
}

func main() {
	fmt.Println(posled(1))
}

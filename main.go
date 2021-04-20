package main

import "fmt"

func pl(a int16, b int16) int16 {
	var c int16 = a + b
	return c
}
func main() {
	var x int16
	x = pl(23, 34)
	fmt.Println(x)
	x = 777
	fmt.Println(x)
}

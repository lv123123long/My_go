package main

import (
	"fmt"
)

func test(x, y int, s string) (int, string) {
	n := x + y
	return n, fmt.Sprint(s, n)
}



func main(){
	var whatever [5]struct{}
	for i := range whatever{
		defer fmt.Println(i)
	}
}

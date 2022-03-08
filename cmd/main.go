package main

import (
	"fmt"
)

func main() {
	a := []int{0, 1, 2}
	var out []*int
	for i := 0; i < 3; i++ {
		fmt.Println(i)
		out = append(out, &a[i])
	}
	for _, o := range out {
		fmt.Println(*o)
	}
}

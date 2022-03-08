package main

import (
	"fmt"
	"log"
)

func main() {
	a := []int{0, 1, 2}
	var out []*int
	for i := 0; i < 3; i++ {
		fmt.Println(i)
		out = append(out, &a[i])
		log.Println(out)
	}
	for _, o := range out {
		fmt.Println(*o)
		log.Println(*o)
	}
}

package main

import "fmt"

func main() {
	var out []*int
	for i := 0; i < 3; i++ {
		fmt.Println(i)
		out = append(out, &i)
	}
	for _, o := range out {
		fmt.Println(*o)
	}
}

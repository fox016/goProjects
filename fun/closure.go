package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	var current, prev, call = 1, 0, 0
	return func() int {
		defer func() { call++ }()
		if call == 0 {
			return prev
		} else if call == 1 {
			return current
		} else {
			current, prev = current+prev, current
			return current
		}
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

package main

// https://tpaschalis.github.io/goroutines-size/
func main() {
	foo(0)
}

func foo(i int) int {
	if i < 1e8 {
		return foo(i + 1)
	}
	return -1
}

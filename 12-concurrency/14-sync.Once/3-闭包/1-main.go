package main

func main() {
	i := incr()
	println(i())
	println(i())
}

func incr() func() int {
	var x int

	return func() int {
		x++
		return x
	}
}

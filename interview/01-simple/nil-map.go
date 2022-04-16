package main

func main() {
	var s []string
	s = append(s, "a")

	var m map[string]int // panic: assignment to entry in nil map
	m["a"] = 1

}

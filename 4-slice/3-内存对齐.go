package main

func main() {
	s1 := []int{1, 2}
	println(len(s1), cap(s1)) // 2 2

	s1 = append(s1, 3, 4, 5)
	println(len(s1), cap(s1)) // 5 6

}

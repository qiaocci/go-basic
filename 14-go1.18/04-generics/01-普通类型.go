package main

func main() {
	ints := map[string]int{
		"a": 1,
		"b": 2,
	}
	floats := map[string]float64{
		"a": 1.1,
		"b": 2.2,
	}
	println(SumInts(ints))
	println(SumFloat(floats))
}

func SumInts(m map[string]int) int {
	var s int
	for _, v := range m {
		s += v
	}
	return s
}

func SumFloat(m map[string]float64) float64 {
	var f float64
	for _, v := range m {
		f += v
	}
	return f
}

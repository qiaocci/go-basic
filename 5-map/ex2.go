package main

import "fmt"

func main() {
	m := map[string]interface{}{
		"a": "1",
		"b": "bbb",
	}

	fmt.Println(m["a"].(string))
	value, ok := m["c"].(string)
	if ok {
		fmt.Println(value)
	} else {
		fmt.Println("error")
	}

}

package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", indexHandler)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		panic(err)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	num := rand.Intn(2)
	if num == 0 {
		time.Sleep(time.Second * 3)
		fmt.Fprintf(w, "slow response")
		return
	}
	fmt.Fprintf(w, "quick response")
}

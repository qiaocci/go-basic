package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)

var mux sync.RWMutex

func main() {
	db := priceDB{}
	db.db = make(map[string]dollars)
	db.db["鞋子"] = 1000
	db.db["袜子"] = 30
	//mux := http.NewServeMux()
	http.Handle("/list", http.HandlerFunc(db.list))
	http.HandleFunc("/price", db.price)
	http.Handle("/updatePrice", http.HandlerFunc(db.updatePrice))
	log.Fatal(http.ListenAndServe(":8000", nil))
}

type dollars float32

func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}

type priceDB struct {
	sync.Mutex
	db map[string]dollars
}

func (p *priceDB) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range p.db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (p *priceDB) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := p.db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	fmt.Fprintf(w, "%s\n", price)

}

func (p *priceDB) updatePrice(w http.ResponseWriter, req *http.Request) {
	mux.Lock()
	defer mux.Unlock()

	query := req.URL.Query()
	item := query.Get("item")
	price := query.Get("price")
	_, ok := p.db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}

	value, err := strconv.ParseFloat(price, 32)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "parse price error: %v\n", price)
		return
	}

	// 修改价格
	p.Lock()
	p.db[item] = dollars(value)
	p.Unlock()
	fmt.Fprintf(w, "update ok, new price is %v\n", p.db[item])
}

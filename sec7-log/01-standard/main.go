package main

import (
	"log"
	"net/http"
	"os"
)

func main()  {
	SetupLogger()
	SimpleHTTPGet("www.google.com")
	SimpleHTTPGet("http://www.google.com")
}

func SetupLogger() {
	logFileLocation, _ := os.OpenFile("test.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0744)
	log.SetOutput(logFileLocation)
}

func SimpleHTTPGet(url string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("failed to get url:%v err:%v", url, err)
	} else {
		log.Printf("Status code for %s: %s", url, resp.Status)
		resp.Body.Close()
	}
}

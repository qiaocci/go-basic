package main

import (
	"go.uber.org/zap"
	"net/http"
)

var logger *zap.Logger

func main() {
	InitLogger()
	defer logger.Sync()
	SimpleHTTPGet("www.google.com")
	SimpleHTTPGet("http://www.google.com")
}

func InitLogger() {
	logger, _ = zap.NewDevelopment()
}

func SimpleHTTPGet(url string) {
	resp, err := http.Get(url)
	if err != nil {
		logger.Error("failed to get url ...", zap.String("url", url), zap.Error(err))
	} else {
		logger.Info("success...", zap.String("statusCode", resp.Status), zap.String("url", url))
		resp.Body.Close()
	}

}

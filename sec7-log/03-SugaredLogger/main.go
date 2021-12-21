package main

import (
	"go.uber.org/zap"
	"net/http"
)

var logger *zap.Logger
var sugarLogger *zap.SugaredLogger

func main() {
	InitLogger()
	defer sugarLogger.Sync()
	SimpleHTTPGet("www.google.com")
	SimpleHTTPGet("http://www.google.com")
}

func InitLogger() {
	logger, _ = zap.NewProduction()
	sugarLogger = logger.Sugar()
}

func SimpleHTTPGet(url string) {
	sugarLogger.Debugf("Trying to hit GET request for %s", url)
	resp, err := http.Get(url)
	if err != nil {
		sugarLogger.Errorf("failed to get url ... url:%s, err:%v\n", url, err)
	} else {
		sugarLogger.Infof("success... url:%v statusCode:%v\n", url, resp.Status)
		resp.Body.Close()
	}

}

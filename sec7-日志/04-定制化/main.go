package main

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
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
	//logger, _ = zap.NewProduction()
	//sugarLogger = logger.Sugar()
	writerSyncer := getLogWriter()
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writerSyncer, zapcore.DebugLevel)

	logger := zap.New(core)
	sugarLogger = logger.Sugar()
}

func getEncoder() zapcore.Encoder {
	// json encoder
	//return zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
	// 普通encoder
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogWriter() zapcore.WriteSyncer {
	//file, _ := os.Create("./test.log")
	//return zapcore.AddSync(file)
	lumberJackLogger := &lumberjack.Logger{
		Filename:   "./test.log",
		MaxSize:    10,
		MaxBackups: 5,
		MaxAge:     30,
		Compress:   false,
	}
	return zapcore.AddSync(lumberJackLogger)
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

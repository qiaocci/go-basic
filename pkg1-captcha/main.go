package main

import (
	"github.com/mojocn/base64Captcha"
	"log"
)

func main() {
	driver := base64Captcha.NewDriverDigit(80, 240, 6, 0.7, 80)
	store := base64Captcha.DefaultMemStore
	c := base64Captcha.NewCaptcha(driver, store)
	id, b64s, err := c.Generate()
	id, b64s, err = c.Generate()
	if err != nil {
		log.Printf("generate failed, err:%v\n", err)
		return
	}
	log.Printf("id=%v\nb64s=%v\n", id, b64s)
}

package main

import (
	"fmt"
	"gocmictest/pkg/setting"
	"gocmictest/routers"
	"net/http"
)

func main() {
	r := routers.InitRouter()
	s := &http.Server{
		Addr:           fmt.Sprintf("192.168.10.104:%d", setting.HTTPPort),
		Handler:        r,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}

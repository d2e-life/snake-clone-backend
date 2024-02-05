package server

import (
	"flag"
	"log"
	"net/http"
	"time"
)

func Run() {
	// 정적 파일 호스팅
	http.Handle("/", http.FileServer(http.Dir("static")))
	server := &http.Server{
		Addr:              *flag.String("addr", ":8080", "http service address"),
		ReadHeaderTimeout: 3 * time.Second,
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

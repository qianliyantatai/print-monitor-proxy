package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"time"
)

func main() {
	target, err := url.Parse("https://qenbvkyonzvlmvrsqbfe.supabase.co")
	if err != nil {
		log.Fatal(err)
	}
	proxy := httputil.NewSingleHostReverseProxy(target)
	proxy.FlushInterval = 100 * time.Millisecond
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	server := &http.Server{
		Addr:         ":" + port,
		Handler:      proxy,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 122 * time.Second,
		IdleTimeout:  90 * time.Second,
	}
	log.Println("proxy listening on :" + port)
	log.Fatal(server.ListenAndServe())
}

package main

import (
	"github.com/CedricJAnslinger/HorseManagement/router"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	log.Println("Server status: Starting server")

	// Create a new router
	log.Println("Server status: Creating router")
	r := router.NewRouter(router.PathNotFoundHandler, router.MethodNotFoundHandler)

	fs := http.FileServer(http.Dir("website"))

	r.HandleFunc("GET", "/", router.Redirect("/calendar_month.html"))
	r.AddDirectoryWeb("website", fs)	// Add directory website with the actual content of the website

	// Configure server
	log.Println("Server status: Creating server")
	server := &http.Server{
		Addr:         "0.0.0.0:8080",
		// Set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler: r,
	}

	// Run server in a goroutine so that it doesn't block.
	// TODO: Run on TLS
	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	log.Println("Server status: Running")

	c := make(chan os.Signal, 1)
	// Allow graceful shutdowns so that the os can safely shutting down the process
	signal.Notify(c, os.Interrupt)

	// Block until we receive the signal.
	<-c

	log.Println("Server status: offline")
	os.Exit(0)
}

package main

import (
	"fmt"
	"github.com/CedricJAnslinger/HorseManagement/router"
	"log"
	"net/http"
)

func main() {
	log.Println("Server status: Starting server")

	// First, create a new router
	log.Println("Server status: Creating router")
	r := router.NewRouter(PathNotFoundHandler, MethodNotFoundHandler)

	// Not the final api, just for testing
	r.Handle("GET", "/", HomeHandler)
	r.Handle("GET", "/month", MonthHanlder)
	r.Handle("GET", "/horses", HorsesHandler)
	r.Handle("GET", "/horses/:name", HorseHandler)

	log.Println("Server status: Start listening")
	http.ListenAndServe(":8080", r)
}

// PathNotFoundHandler handles a request for which no path exists.
// @parameter - w: http.ResponseWriter >> Interface used by an HTTP handler to construct an HTTP response.
// @parameter - r: http.Request(Pointer) >> Received HTTP request.
func PathNotFoundHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Couldn't find path!")
}

// PathNotFoundHandler handles a request for which the wrong HTTPMethod was used
// @parameter - w: http.ResponseWriter >> Interface used by an HTTP handler to construct an HTTP response.
// @parameter - r: http.Request(Pointer) >> Received HTTP request.
func MethodNotFoundHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Couldn't accept method on this path")
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Project Horse-Management")
}

func MonthHanlder(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Display one month")
}

func HorsesHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Display all horses")
}

func HorseHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Display horse:  %s", r.Form.Get("name"))
}


package router

import (
	"fmt"
	"net/http"
)

// PathNotFoundHandler handles a request for which no path exists.
// @parameter - w: http.ResponseWriter >> Interface used by an HTTP handler to construct an HTTP response.
// @parameter - r: http.Request(Pointer) >> Received HTTP request.
func PathNotFoundHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "HTTP Response Status:", http.StatusNotFound)
}

// PathNotFoundHandler handles a request for which the wrong HTTPMethod was used
// @parameter - w: http.ResponseWriter >> Interface used by an HTTP handler to construct an HTTP response.
// @parameter - r: http.Request(Pointer) >> Received HTTP request.
func MethodNotFoundHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "HTTP Response Status:", http.StatusMethodNotAllowed)
}


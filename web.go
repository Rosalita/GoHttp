// Really basic web server
// running this program and accessing the url http://localhost:8080/chocolate
// loads a page which displays text "Hello I am a webserver and I love Chocolate"

package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// handler function is of the type http.HandlerFunc
	// It takes two arguments.
	// http.Response.Writer assembles the HTTP server's response by writing to it
	// http.Request is a datastructure that represents the client HTTP request.

	fmt.Fprintf(w, "Hello I am a webserver and I love %s!", r.URL.Path[1:])
	// r.URL.Path is the path component of the request URL.
	// The trailing [1:] means "create a subslice of path from the 1st character to the end."
	// This drops the leading "/" from the path name
}

func main() {
	// this tells the net/http package to handle all requests made to to the web root "/" with the handler function

	http.HandleFunc("/", handler)
	// then we tell the webserver to listen and serve on port 8080
	http.ListenAndServe(":8080", nil)
}

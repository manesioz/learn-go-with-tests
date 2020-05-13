package app

import (
	"fmt"
	"net/http"
)

//PlayerServer function acts as a server, taking a request and writing a response
func PlayerServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "20")
}

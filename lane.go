package lane

import (
	"fmt"
	"net/http"
)

// Lane object
type Lane struct {
	Routes []string
}

var mux = http.NewServeMux()

// Start start webserver
func Start(addr string) error {
	fmt.Println("Lane started")

	return http.ListenAndServe(addr, mux)
}

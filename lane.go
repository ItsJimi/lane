package lane

import (
	"fmt"
	"net/http"
)

// Context context
type Context struct {
	Response http.ResponseWriter
	Request  *http.Request
}

// Send send simple response
func (c *Context) Send(str string) {
	c.Response.Write([]byte(str))
}

// Get handle get method
func Get(pattern string, handler func(Context)) {
	http.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			return
		}

		handler(Context{
			Response: w,
			Request:  r,
		})
	})
}

// Start start webserver
func Start(addr string) error {
	fmt.Println("Lane started")
	return http.ListenAndServe(addr, nil)
}

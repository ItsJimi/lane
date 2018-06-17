package lane

import "net/http"

// Context context
type Context struct {
	Response http.ResponseWriter
	Request  *http.Request
}

// Send send simple response
func (c *Context) Send(str string) {
	c.Response.Write([]byte(str))
}

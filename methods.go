package lane

import (
	"net/http"
	"time"
)

func timeHandler(w http.ResponseWriter, r *http.Request) {
	tm := time.Now().Format(time.RFC1123)
	w.Write([]byte("The time is: " + tm))
}

// Get handle GET method
func Get(pattern string, handler func(Context)) {
	th := http.HandlerFunc(timeHandler)

	mux.Handle(pattern, th)
}

// Get handle GET method
// func Get(pattern string, handler func(Context)) {
// 	mux.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Println(r.Method)
// 		if r.Method != "GET" {
// 			http.NotFound(w, r)
// 			return
// 		}

// 		handler(Context{
// 			Response: w,
// 			Request:  r,
// 		})
// 	})
// }

// Post handle POST method
func Post(pattern string, handler func(Context)) {
	mux.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.NotFound(w, r)
			return
		}

		handler(Context{
			Response: w,
			Request:  r,
		})
	})
}

// Put handle PUT method
func Put(pattern string, handler func(Context)) {
	mux.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "PUT" {
			http.NotFound(w, r)
			return
		}

		handler(Context{
			Response: w,
			Request:  r,
		})
	})
}

// Patch handle PATCH method
func Patch(pattern string, handler func(Context)) {
	mux.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "PATCH" {
			http.NotFound(w, r)
			return
		}

		handler(Context{
			Response: w,
			Request:  r,
		})
	})
}

// Delete handle DELETE method
func Delete(pattern string, handler func(Context)) {
	mux.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "DELETE" {
			http.NotFound(w, r)
			return
		}

		handler(Context{
			Response: w,
			Request:  r,
		})
	})
}

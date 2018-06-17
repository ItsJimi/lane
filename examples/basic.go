package main

import (
	"github.com/ItsJimi/lane"
)

func main() {
	lane.Get("/", func(c lane.Context) {
		c.Send("Hello")
	})

	lane.Start(":8080")
}

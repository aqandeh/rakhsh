package main

import (
	"github.com/aqandeh/rakhsh"
)

func main() {
	app := rakhsh.New()

	app.GET("/", func(c *rakhsh.Context) {
		c.JSON(200, map[string]string{
			"framework": "Rakhsh",
			"status": "running",
		})
	})

	app.GET("/test", func(c *rakhsh.Context) {
		c.String(200, "Hello from Rakhsh 🚀")
	})

	app.Listen(":8080")
}

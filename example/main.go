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

	app.Listen(":8080")
}

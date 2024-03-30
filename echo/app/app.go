package main

// run
// /Users/okawara/code/go/src/echo/app
// go run app.go

import (
	"app/presentation"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func main() {

	e := echo.New()
	e.Logger.SetLevel(log.DEBUG)
	// e.GET("/", func(c echo.Context) error {
	// 	return c.String(http.StatusOK, "Hello, World!")
	// })

	e.GET("/users/:id", getUser)
	presentation.Register(e)

	e.Start(":1323")

}

func getUser(c echo.Context) error {

	presentation.Hello()
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

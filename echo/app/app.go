package main

// run
// /Users/okawara/code/go/src/echo/app
// go run app.go

import (
	"app/presentation"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func main() {

	e := echo.New()
	e.Logger.SetLevel(log.DEBUG)
	// e.GET("/users/:id", getUser)
	presentation.Register(e)

	e.Start(":1323")

}

// func getUser(c echo.Context) error {

// 	presentation.Hello()
// 	id := c.Param("id")
// 	return c.String(http.StatusOK, id)
// }

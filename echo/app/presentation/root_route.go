package presentation

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func printStr() {

	fmt.Println("hello")
}

// ポインタ型の変数を引数
func Register(api *echo.Echo) {
	api.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	api.GET("/users/:id", getUser)
}

func getUser(c echo.Context) error {

	printStr()
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

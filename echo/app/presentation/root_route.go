package presentation

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Hello() {

	fmt.Println("hello")
}

// ポインタ型の変数を引数
func Register(api *echo.Echo) {
	api.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

}

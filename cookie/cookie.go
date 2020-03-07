package cookie

import (
	"net/http"

	"github.com/labstack/echo"
)

func Write(c echo.Context, name string, value string) {
	cookie := new(http.Cookie)
	cookie.Name = name
	cookie.Value = value
	c.SetCookie(cookie)
}

func Read(c echo.Context, name string) string {
	cookie, err := c.Cookie(name)
	if err != nil {
		return ""
	}
	return cookie.Value
}

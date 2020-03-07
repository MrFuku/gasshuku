package auth

import (
	"strconv"

	"github.com/MrFuku/gasshuku/cookie"
	"github.com/labstack/echo"
)

func UserLogin(c echo.Context, id int) {
	cookie.Write(c, "logginUser", strconv.Itoa(id))
}

func UserLoggedin(c echo.Context) bool {
	id := cookie.Read(c, "logginUser")
	return id != ""
}

func LoginUser(c echo.Context) int {
	id := cookie.Read(c, "logginUser")
	if id != "" {
		return 0
	}
	res, _ := strconv.Atoi(id)
	return res
}

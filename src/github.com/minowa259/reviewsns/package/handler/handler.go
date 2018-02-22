package handler

import (
	"github.com/labstack/echo"
	"net/http"

	"../config"
	"fmt"
)

var Conf config.Data = config.LoadConfig()


// HomePage でrootの表示管理
func HomePage() echo.HandlerFunc {
	data := struct {
		Pagetitle string
		Config config.Data
	} {
		Pagetitle: "ログイン・新規登録",
		Config: Conf,
	}
	fmt.Printf("%#v\n", data)
	return func(c echo.Context) error {
		return c.Render(http.StatusOK, "index", data)
	}
}


func LoginPage() echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Get("csrf").(string)
		data := struct {
			Pagetitle string
			Csrftoken string
			Config config.Data
		} {
			Pagetitle: "レビュー専用のポータル",
			Csrftoken: token,
			Config: Conf,
		}
		return c.Render(http.StatusOK, "login", data)
	}
}

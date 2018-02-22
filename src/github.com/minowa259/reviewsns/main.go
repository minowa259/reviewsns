package main

import (
	"./package/handler"
	"./package/db"

	"github.com/labstack/echo"
	"html/template"
	"io"
	"github.com/labstack/echo/middleware"
)

// Template の宣言を行う
type Template struct {
	templates *template.Template
}

// Render でレンダリング
func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {

	/*
		DBに接続
	*/
	db := db.SnsDB{}
	db.Connect()
	defer db.Close() //切断


	/*
		echoのインスタンスを作成
	*/
	e := echo.New()
	e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		TokenLookup: "header:X-XSRF-TOKEN",
	}))

	/*
		templateのレンダラーを設定
	*/
	t := &Template{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}

	e.Renderer = t


	/*
		静的ファイルのパスを設定
	*/
	e.Static("/assets", "static")

	/*
		ルーティングの設定
	*/
	e.GET("/", handler.HomePage())
	e.GET("/login", handler.LoginPage())


	// サーバーを開始
	e.Logger.Fatal(e.Start(":1111"))
}

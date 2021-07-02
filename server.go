package main

import (
	"io"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func getUser(c echo.Context) error {
	// User ID from path `users/:id`
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

func show(c echo.Context) error {
	// Get team and member from the query string
	team := c.QueryParam("team")
	member := c.QueryParam("member")
	return c.String(http.StatusOK, "team:"+team+", member:"+member)
}

func save(c echo.Context) error {
	// Get name and email
	name := c.FormValue("name")
	email := c.FormValue("email")
	return c.String(http.StatusOK, "name:"+name+", email:"+email)
}

func save2(c echo.Context) error {
	// Get name
	name := c.FormValue("name")
	// Get avatar
	avatar, err := c.FormFile("avatar")
	if err != nil {
		return err
	}

	// Source
	src, err := avatar.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// Destination
	dst, err := os.Create(avatar.Filename)
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	return c.HTML(http.StatusOK, "<b>Thank you! "+name+"</b>")
}

func sendJson(c echo.Context) error {

	type User struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	}

	users := []User{
		{Id: 1, Name: "test1"},
		{Id: 2, Name: "test2"},
	}

	if err := c.Bind(users); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, users)
	// or
	// return c.XML(http.StatusCreated, u)
}

func main() {
	// echo オブジェクトの指定
	e := echo.New()
	// ルート
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	// パス
	e.GET("/users/:id", getUser)

	// クエリ
	e.GET("/show", show)

	// ポスト
	e.POST("/save", save)

	// リクエストにJsonを返す
	e.GET("/users", sendJson)

	e.Static("/static", "static")

	e.Logger.Fatal(e.Start(":1323"))
}

// やりたいこと
// ユーザから送られてきたものを受信する
// 階層を持つJsonを送る
// 送られてきたJsonを適切にパースする
// form,buttonの情報を取得する
// DBからJsonを生成
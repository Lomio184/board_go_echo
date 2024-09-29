package host

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"board-echo/core/render"
	"board-echo/core/client"
)

var posts = []client.Post{}

func createEchoInst() *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Renderer = render.CreateRenderer() 
	return e 
}

func EchoManager() {
	e := createEchoInst()
	addPost, newPost, showPosts := client.PostManager(&posts)

	e.GET("/", echo.HandlerFunc(showPosts))
	e.GET("/new", echo.HandlerFunc(newPost))
	e.POST("/add", echo.HandlerFunc(addPost))

	e.Logger.Fatal(e.Start(":8080"))
}
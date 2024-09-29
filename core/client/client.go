package client

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type Post struct {
	ID      string
	Title   string
	Content string
}

type HandlerFunc func(echo.Context) error

func PostManager(posts *[]Post) (HandlerFunc, HandlerFunc, HandlerFunc) {
	addPost := func(c echo.Context) error {
		title := c.FormValue("title")
		content := c.FormValue("content")
		id := uuid.New().String()

		*posts = append(*posts, Post{ID: id, Title: title, Content: content})
		c.Redirect(http.StatusSeeOther, "/")

		return nil
	}

	// removePost := func(index int) {
	// 	if index < 0 || index >= len(*posts) {
	// 		fmt.Println("Invalid Index")
	// 		return
	// 	}

	// 	*posts = append((*posts)[:index], (*posts)[index+1:]...)
	// 	fmt.Println("Post remove at index : ", index)
	// }

	newPost := func(c echo.Context) error {
		c.Render(http.StatusOK, "new_post.html", nil)
		return nil
	}

	showPosts := func(c echo.Context) error {
		c.Render(http.StatusOK, "index.html", map[string]interface{}{
			"posts": *posts})

		return nil
	}

	return addPost, newPost, showPosts
}

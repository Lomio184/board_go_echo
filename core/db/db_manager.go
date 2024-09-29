package db

import (
    "database/sql"
	"board-echo/core/client"
    _ "github.com/mattn/go-sqlite3"
)


func InitDB(filepath string) (*sql.DB, error) {
    db, err := sql.Open("sqlite3", filepath)
    if err != nil {
        return nil, err
    }
    return db, nil
}

func GetPosts(db *sql.DB) ([]client.Post, error) {
    rows, err := db.Query("SELECT id, title, content FROM posts")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var posts [] client.Post
    for rows.Next() {
        var post client.Post
        err := rows.Scan(&post.ID, &post.Title, &post.Content)
        if err != nil {
            return nil, err
        }
        posts = append(posts, post)
    }
    return posts, nil
}
package models

import (
	"errors"
	"html"
	"time"
)

type Post struct {
	ID         int    `json:"id,omitempty"`
	Title      string `json:"title,omitempty"`
	Body       string `json:"body,omitempty"`
	AuthorId   uint64 `json:"author_id,omitempty"`
	AuthorNick string `json:"author_nick,omitempty"`
	Likes      int    `json:"likes"`
	CreatedAt  string `json:"created_at,omitempty"`
}

func (post *Post) Prepare() error {
	if err := post.validate(); err != nil {
		return err
	}
	post.format()
	return nil
}

func (post *Post) validate() error {
	if post.Title == "" {
		return errors.New("title is required")
	}
	if post.Body == "" {
		return errors.New("body is required")
	}
	return nil
}

func (post *Post) format() {
	post.Title = html.EscapeString(post.Title)
	post.Body = html.EscapeString(post.Body)
	post.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
}

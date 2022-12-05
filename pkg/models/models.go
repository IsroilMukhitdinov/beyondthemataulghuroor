package models

import "time"

type Snippet struct {
	Id        int
	Author    string
	CreatedAt time.Time
	Content   string
	Comments  []Comment
}

type Comment struct {
	Id          int
	CommentedBy string
	RepliedTo   string
	CreatedAt   time.Time
	Content     string
}

type User struct {
	Id        int
	Name      string
	JoinedAt  time.Time
	Snippets  []Snippet
	Commments []Comment
	Password  string
	Email     string
	Followers []User
	Following []User
	Settings  map[string]string
}

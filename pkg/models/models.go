package models

import (
	"database/sql"
	"time"
)

type Conn struct {
	DB *sql.DB
}

type Snippet struct {
	Id        int
	Author    string
	CreatedAt time.Time
	Content   string
}

type User struct {
	Id        int
	Name      string
	JoinedAt  time.Time
	Snippets  []Snippet
	Password  string
	Email     string
	Followers []User
	Following []User
	Settings  map[string]string
}

func (c *Conn) Latest() ([]*Snippet, error) {
	stmt := `select author, created_at, content from snippets;`

	rows, err := c.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	snippets := []*Snippet{}

	for rows.Next() {
		snippet := &Snippet{}

		err = rows.Scan(&snippet.Author, &snippet.CreatedAt, &snippet.Content)
		if err != nil {
			return nil, err
		}

		snippets = append(snippets, snippet)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return snippets, nil
}

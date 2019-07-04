package db

import (
	"database/sql"
	"time"

	"github.com/sirupsen/logrus"
)

// "database/sql"

// Blog ...
type Blog struct {
	ID            string    `json:"id"`
	Title         string    `json:"title"`
	Description   string    `json:"description"`
	Author        string    `json:"author"`
	FormattedDate string    `json:"formatted_date"`
	DocType       string    `json:"doc_type"`
	MDSrc         string    `json:"md_src"`
	HTMLSrc       string    `json:"html_src"`
	CreatedAt     time.Time `json:"created_at"`
	IsFeatured    bool      `json:"is_featured"`
	IsPublic      bool      `json:"is_public"`
	IsDeleted     bool      `json:"is_deleted"`
}

// CreateTable ..
func (b *Blog) CreateTable(pg *Postgres) {
	str := `
	CREATE TABLE blogs (
		id TEXT UNIQUE NOT NULL,
		title TEXT NOT NULL,
		description TEXT,
		author TEXT,
		formatted_date TEXT,
		doc_type TEXT,
		md_src TEXT,
		html_src TEXT,
		created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		is_featured BOOLEAN,
		is_public BOOLEAN,
		is_deleted BOOLEAN 
	);`

	_, err := pg.DB.Exec(str)
	if err != nil {
		logrus.Warnf("%v\n", err)
	}
}

// str := `
// 	INSERT INTO admins (email, google_id)
// 	VALUES ($1, $2)`

// 	return PG.DB.Exec(str, a.Email, a.GoogleID)

// Insert .
func (b *Blog) Insert() (sql.Result, error) {
	str := `
	INSERT INTO blogs (id, title, description, author, formatted_date,
		doc_type, md_src, html_src, created_at, is_featured, is_public, is_deleted)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)`

	return PG.DB.Exec(str, b.ID, b.Title, b.Description, b.Author, b.FormattedDate,
		b.DocType, b.MDSrc, b.HTMLSrc, b.CreatedAt, b.IsFeatured, b.IsPublic, b.IsDeleted)
}

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
	Likes         int       `json:"likes"`
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
		is_deleted BOOLEAN,
		likes INT NOT NULL DEFAULT 0
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

// Query ...
func (b *Blog) Query() error {
	str := `SELECT * FROM admins WHERE id=$1;`

	row := PG.DB.QueryRow(str, b.ID)

	return row.Scan(&b)
}

// Update .
func (b *Blog) Update(new *Blog) (sql.Result, error) {
	// TODO: Find a better strategy (there must be some other tool)

	str := `
	UPDATE blogs
	SET id = $2, title = $3, description = $4, author = $5, formatted_date = $6, doc_type = $7,
	md_src = $8, html_src = $9, created_at = $10, is_featured = $11, is_public = $12, is_deleted = $13
	WHERE id = $1;`

	return PG.DB.Exec(str, b.ID, new.ID, new.Title, new.Description, new.Author, new.FormattedDate,
		new.DocType, new.MDSrc, new.HTMLSrc, new.CreatedAt, new.IsFeatured, new.IsPublic, new.IsDeleted)
}

// Delete .
func (b *Blog) Delete() (sql.Result, error) {
	str := `
	UPDATE blogs
	SET is_deleted = $2
	WHERE id = $1;`

	return PG.DB.Exec(str, b.ID, true)
}

// DeletePerm .
func (b *Blog) DeletePerm() (sql.Result, error) {
	str := `
	DELETE FROM blogs
	WHERE id = $1;`

	return PG.DB.Exec(str, b.ID)
}

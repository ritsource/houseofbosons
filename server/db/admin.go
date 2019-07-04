package db

import (
	"database/sql"

	"github.com/sirupsen/logrus"
)

// Admin ...
type Admin struct {
	Email    string `json:"email"`
	GoogleID string `json:"id"`
}

// CreateTable ..
func (a *Admin) CreateTable(pg *Postgres) {
	str := `
	CREATE TABLE admins (
		google_id TEXT,
		email TEXT UNIQUE NOT NULL
	);`

	_, err := pg.DB.Exec(str)
	if err != nil {
		logrus.Warnf("%v\n", err)
	}
}

// Insert ...
func (a *Admin) Insert() (sql.Result, error) {
	str := `
	INSERT INTO admins (email, google_id)
	VALUES ($1, $2)`

	return PG.DB.Exec(str, a.Email, a.GoogleID)
}

// Query ...
func (a *Admin) Query() error {
	str := `SELECT email, google_id FROM admins WHERE email=$1;`

	row := PG.DB.QueryRow(str, a.Email)

	return row.Scan(&a.Email, &a.GoogleID)
}

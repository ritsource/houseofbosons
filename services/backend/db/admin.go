package db

import (
	"database/sql"
	"fmt"
)

// Admin ...
type Admin struct {
	Email    string
	GoogleID string
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
	str := `SELECT * FROM admin WHERE email=$1;`
	// Query

	row := PG.DB.QueryRow(str, a.Email)
	fmt.Printf("row - %+v\n", row)

	return nil
	// err := row.Scan(a.Email, a.GoogleID)
	// switch  {
	// case condition:

	// }
}

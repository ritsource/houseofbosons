package db

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

var PG Postgres

func init() {
	config := Config{}

	config.host = os.Getenv("POSTGRES_HOST")
	config.user = os.Getenv("POSTGRES_USER")
	config.password = os.Getenv("POSTGRES_PASSWORD")
	config.dbname = os.Getenv("POSTGRES_DB_NAME")

	var err error
	config.port, err = strconv.Atoi(os.Getenv("POSTGRES_PORT"))
	if err != nil {
		panic("invalid POSTGRES_PORT")
	}

	PG.cfg = config

	err = PG.Connect()
	if err != nil {
		logrus.Panic(err)
	}
	logrus.Infoln("successfully connected to database")

	PG.PrepTables()
}

// Postgres ...
type Postgres struct {
	DB  *sql.DB
	cfg Config
}

type Config struct {
	host     string
	port     int
	user     string
	password string
	dbname   string
}

// Connect ...
func (pg *Postgres) Connect() error {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		pg.cfg.host, pg.cfg.port, pg.cfg.user, pg.cfg.password, pg.cfg.dbname)

	var err error
	pg.DB, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		return err
	}

	return pg.DB.Ping()
}

// Close ...
func (pg *Postgres) Close() {
	pg.Close()
}

// PrepTables .
func (pg *Postgres) PrepTables() {
	pg.createAdminsTable()
}

func (pg *Postgres) createAdminsTable() {
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

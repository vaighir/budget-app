package drivers

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

const maxOpenConn = 10
const maxIdleConn = 5
const maxConnLifetime = 2 * time.Minute

type DB struct {
	SQL *sql.DB
}

var conn = &DB{}

func ConnectSQL(dsn string) (*DB, error) {
	d, err := NewDatabase(dsn)

	if err != nil {
		log.Fatal(err)
	}

	d.SetMaxOpenConns(maxOpenConn)
	d.SetMaxIdleConns(maxIdleConn)
	d.SetConnMaxLifetime(maxConnLifetime)

	conn.SQL = d

	return conn, nil
}

func NewDatabase(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)

	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

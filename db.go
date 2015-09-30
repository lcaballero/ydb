package ydb

import (
	_ "code.google.com/p/go-sqlite/go1/sqlite3"
	"database/sql"
	"death"
	"syscall"
	"log"
)

type DB struct {
	Db *sql.DB
	Filename string
	Driver string
	Err error
}

type Transaction func(tx *sql.Tx) error
type Prepared func(stmt *sql.Stmt) error
const DefaultSqlDriver = "sqlite3"

func NewDb(filename string) *DB {
	return &DB{
		Filename: filename,
		Driver: DefaultSqlDriver,
	}
}

func (d *DB) Start() *DB {
	d.Db, d.Err = sql.Open(d.Driver, d.Filename)
	if d.Err != nil {
		log.Println(d.Err)
	}
	go func() {
		death.NewDeath(syscall.SIGINT, syscall.SIGTERM).WaitForDeath(d.Db)
	}()
	return d
}

func (d *DB) Begin(txfn Transaction) {
	tx, err := d.Db.Begin()
	if err != nil {
		log.Printf("While beginning transaction %v\n", err)
	}
	defer tx.Commit()

	err = txfn(tx)
	if err != nil {
		log.Printf("While executing transaction %v", err)
	}
}

func (d *DB) Prepare(code string, fn Prepared) {
	d.Begin(func(tx *sql.Tx) error {
		stmt, err := tx.Prepare(code)
		if err != nil {
			return err
		}
		defer stmt.Close()
		return fn(stmt)
	})
}


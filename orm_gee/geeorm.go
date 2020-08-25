package geeorm

import (
	"database/sql"
	"go-experiment/orm_gee/dialect"
	"go-experiment/orm_gee/log"
	"go-experiment/orm_gee/session"
)

type Engine struct {
	db      *sql.DB
	dialect dialect.Dialect
}

func NewEngine(driver, source string) (e *Engine, err error) {
	db, err := sql.Open(driver, source)
	if err != nil {
		log.Error(err)
		return
	}
	// Send a ping to make sure the database connection is alive.
	if err = db.Ping(); err != nil {
		log.Error(err)
		return
	}
	// make sure the specific dialect exists
	dial, ok := dialect.GetDialect(driver)
	if !ok {
		log.Error(err)
		return
	}
	e = &Engine{db: db, dialect: dial}
	log.Info("Connect database success")
	return
}

func (engin *Engine) Close() {
	if err := engin.db.Close(); err != nil {
		log.Error("Failed to close database")
	}
	log.Info("Close database successfully")
}

func (engin *Engine) NewSession() *session.Session {
	return session.New(engin.db, engin.dialect)
}

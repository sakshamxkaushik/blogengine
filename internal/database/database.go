package database

import (
	"github.com/go-pg/pg/v10"
)

// TODO create localdatabase
func NewDBOptions() *pg.Options {
	return &pg.Options{
		Addr:     "babar.db.elephantsql.com",
		Database: "ebhxorlp",
		User:     "ebhxorlp",
		Password: "yJsZVCcDF8dZqkAh6KnQWlVh22yn1_tY",
	}
}

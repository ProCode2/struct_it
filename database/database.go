package database

import "github.com/procode2/structio/database/postgres"

var Db Storer = postgres.NewPostgresStore()

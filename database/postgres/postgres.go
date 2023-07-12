package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/procode2/structio/models"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
)

type PostgresStore struct {
	db *bun.DB
}

func (p *PostgresStore) Init() {
	ctx := context.Background()

	_, err := p.db.NewCreateTable().
		IfNotExists().
		Model((*models.User)(nil)).
		Exec(ctx)
	if err != nil {
		log.Fatal(err)
	}

	_, err = p.db.NewCreateTable().
		IfNotExists().
		Model((*models.Path)(nil)).
		Exec(ctx)
	if err != nil {
		log.Fatal(err)
	}

	_, err = p.db.NewCreateTable().
		IfNotExists().
		Model((*models.Level)(nil)).
		Exec(ctx)
	if err != nil {
		log.Fatal(err)
	}

	_, err = p.db.NewCreateTable().
		IfNotExists().
		Model((*models.Bit)(nil)).
		Exec(ctx)
	if err != nil {
		log.Fatal(err)
	}
}

func NewPostgresStore() *PostgresStore {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("can not load .env at the moment")
	}
	fmt.Println("DB_USER", os.Getenv("DB_USER"))

	dsn := fmt.Sprintf("postgres://%s:%s@localhost:%s/%s?sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"))
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	db := bun.NewDB(sqldb, pgdialect.New())

	if db == nil {
		log.Fatal("Could not connect to DB")
	}

	db.AddQueryHook(bundebug.NewQueryHook(
		bundebug.WithVerbose(true),
		bundebug.FromEnv("BUNDEBUG"),
	))

	return &PostgresStore{
		db: db,
	}
}

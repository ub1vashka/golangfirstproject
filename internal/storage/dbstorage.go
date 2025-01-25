package storage

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/golang-migrate/migrate/v4"
	"github.com/jackc/pgx/v5"
	"github.com/ub1vashka/golangfirstproject/internal/domain/models"
	"github.com/ub1vashka/golangfirstproject/internal/logger"
)

type DBStorage struct {
	conn *pgx.Conn
}

func NewDB(ctx context.Context, addr string) (*DBStorage, error) {
	conn, err := pgx.Connect(ctx, addr)
	if err != nil {
		return nil, err
	}
	return &DBStorage{conn: conn}, nil
}

func (dbs *DBStorage) GetBooks() ([]models.Book, error) {
	log := logger.Get()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rows, err := dbs.conn.Query(ctx, "SELECT * FROM books")
	if err != nil {
		log.Error().Err(err).Msg("failed get data from table books")
		return nil, err
	}
	var books []models.Book
	for rows.Next() {
		var book models.Book
		if err = rows.Scan(&book.BID, &book.Lable, &book.Author, &book.Description, &book.WritedAt); err != nil {
			log.Error().Err(err).Msg("failed scan raws data")
			return nil, err
		}
		books = append(books, book)
	}
	return books, nil
}

func Migrations(dbDsn string, migratePath string) error {
	log := logger.Get()
	migrPath := fmt.Sprint("file://", migratePath)
	m, err := migrate.New(migrPath, dbDsn)
	if err != nil {
		log.Error().Err(err).Msg("failed to db connect")
		return err
	}
	if err := m.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			log.Debug().Msg("no migrations apply") // если никаких изменений не было в миграциях
			return nil
		}
		log.Error().Err(err).Msg("run migrations failed")
		return err
	}
	log.Debug().Msg("all migrations apply")
	return nil
}

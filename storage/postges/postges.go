package postgres

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"

	"booking/storage"
)

type Store struct {
	db             *pgxpool.Pool
	user           storage.UserRepoI
	role           storage.RoleRepoI
}

func NewPostgres(psqlConnString string) storage.StorageRepoI {
	config, err := pgxpool.ParseConfig(psqlConnString)
	if err != nil {
		log.Panicf("Unable to parse connection string.: %v", err)
	}

	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		log.Panicf("Unable to connect to the database: %v", err)
	}

	return &Store{
		db: pool,
	}
}

func (s *Store) CloseDB() {
	s.db.Close()
}

func (s *Store) User() storage.UserRepoI {
	if s.user == nil {
		s.user = &userRepo{
			db: s.db,
		}
	}
	return s.user
}

func (s *Store) Role() storage.RoleRepoI {
	if s.role == nil {
		s.role = &roleRepo{
			db: s.db,
		}
	}
	return s.role
}
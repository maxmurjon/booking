package postgres

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"

	"booking/storage"
)

type Store struct {
	db     *pgxpool.Pool
	role   storage.RoleRepoI
	user   storage.UserRepoI
	doctor storage.DoctorRepoI
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

func (s *Store) Doctor() storage.DoctorRepoI {
	if s.doctor == nil {
		s.doctor = &doctorRepo{
			db: s.db,
		}
	}
	return s.doctor
}

package storage

import (
	"github.com/Sharofiddin701/medium-clone-simple/storage/postgres"
	"github.com/Sharofiddin701/medium-clone-simple/storage/repo"
	"github.com/jmoiron/sqlx"
)

type StorageI interface {
	User() repo.UserStorageI
}

type StoragePg struct {
	UserRepo repo.UserStorageI
}

func NewStorage(psqlConn *sqlx.DB) StorageI {
	return &StoragePg{
		UserRepo: postgres.NewUserStorage(psqlConn),
	}
}

func (s *StoragePg) User() repo.UserStorageI {
	return s.UserRepo
}

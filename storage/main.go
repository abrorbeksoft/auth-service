package storage

import "github.com/abrorbeksoft/auth-service/storage/repos"

type StorageI interface {
	UserStorage() repos.UserStorageI
}
package repos

import "github.com/abrorbeksoft/auth-service/genproto/auth"

type UserStorageI interface {
	Create(user *auth.User) (string,error)
}
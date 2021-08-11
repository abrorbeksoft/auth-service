package postgres

import (
	"database/sql"
	"github.com/abrorbeksoft/auth-service/genproto/auth"
	"github.com/abrorbeksoft/auth-service/storage/repos"
)

type userRepo struct {
	session *sql.DB
}

func NewUserRepo(ses *sql.DB) repos.UserStorageI {
	return &userRepo{session: ses}
}

func (rep *userRepo) Create(user *auth.User) (string,error) {
	query, _ :=rep.session.Query(`SELECT * FROM users`);
	query.Next();
	return "Hello",nil
}
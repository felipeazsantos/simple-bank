package api

import (
	"time"

	db "github.com/felipeazsantos/simple_bank/db/sqlc"
	"github.com/felipeazsantos/simple_bank/util"
)

func randomUser() (db.User, error) {
	hashedPassword, _ := util.HashPassword(util.RandomString(6))
	return db.User{
		Username:          util.RandomOwner(),
		HashedPassword:    hashedPassword,
		FullName:          util.RandomString(6),
		Email:             util.RandomEmail(),
		PasswordChangedAt: time.Now(),
		CreatedAt:         time.Now(),
	}, nil
}

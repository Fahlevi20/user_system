package User

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	dbPool *pgxpool.Pool
}

// Create constructor function which capable created and return new instance from userRepository
func NewUserRepository(dbPool *pgxpool.Pool) UserRepository {
	return UserRepository{dbPool: dbPool}
}

// create logic
// RegusterUser for registering user to database
func (r *UserRepository) RegisterUser(user *RegisterUser) error {
	//Create local variable for keep the error
	var err error

	query := "INSERT INTO users (first_name, last_name, password, email) VALUES ($1,$2,$3,$4)"
	_, err = r.dbPool.Exec(context.Background(), query, user.First_name, user.Last_name, user.Password, user.Email)

	if err != nil {
		log.Println(err)
	}
	return nil
}

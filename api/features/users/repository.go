package User

import (
	"context"
	"log"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type userRepository struct{
	dbPool *pgxpool.Pool
}

//Create constructor function which capable created and return new instance from userRepository 
func NewUserRepository(dbPool *pgxpool.Pool) UserRepository{ //create function parameter can received one argument from dbPool with type pgxpool.Pool
	return &userRepository(dbPool:dbPool) // Create new instance using Pointer from userRepository
}

// create logic

func (r *userRepository) RegisterUser(user *RegisterUser) error{
	query:= "INSERT INTO users (first_name, last_name, password, email) VALUES ($1,$2,$3,$4)"
	_,err= r.dbPool.Exec(context.Background(),query, user.First_name, user.Last_name, user.Password, user.Email)
	
	if err!= nil {
		log.Println(error())
	}
	return nil
}
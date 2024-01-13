package User

import "github.com/jackc/pgx/v5/pgxpool"

type userRepository struct{
	dbPool *pgxpool.Pool
}

func NewUserRepository(dbPool *pgxpool.Pool) UserRepository{
	return &userRepository(dbPool:dbPool)
}
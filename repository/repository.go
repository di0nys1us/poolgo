package repository

import "github.com/jmoiron/sqlx"

type Repository interface {
}

type DatabaseRepository struct {
	database *sqlx.DB
}

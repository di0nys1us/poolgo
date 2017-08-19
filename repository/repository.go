package repository

import "github.com/di0nys1us/poolgo/models"

type Repository interface {
	FindUser(id int) *models.User
}

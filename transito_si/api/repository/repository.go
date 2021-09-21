package repository

import (
	"unal.edu.co/rest-example/api/dto"
	"unal.edu.co/rest-example/api/model"
)

type UserRepository interface {
	SearchUser(dto.SearchUser) ([]model.User, error)
	CreateUser(model.User) (model.User, error)
	CreateTable() ()
}

package auth

import "ws/src/user"

type UserController struct {
	UserRepo *user.Repository
}

func NewController(repo *user.Repository) *UserController {
	return &UserController{UserRepo: repo}
}

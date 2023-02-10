package models

type User struct {
	Username string `binding:"required"`
	Password string `binding:"required"`
	Name     string `binding:"required"`
}

type Login struct {
	Username string `binding:"required"`
	Password string `binding:"required"`
}

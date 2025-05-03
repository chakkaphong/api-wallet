package entities

import "time"

type User struct {
	UserId    int
	Username  string
	Email     string
	FirstName string
	LastName  string
	CreatedAt time.Time
}

func (User) TableName() string {
	return "users"
}

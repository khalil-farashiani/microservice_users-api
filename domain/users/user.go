package users

import "time"

type User struct {
	Id          int64
	FirstName   string
	LastName    string
	Email       string
	DateCreated time.Time
}

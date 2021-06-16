package model

import "time"

type Account struct {
	AccountID int64
	FirstName string
	Lastname string
	UserName string
	EMail string
	Password string
	Birthdate time.Time
	Role Role
}

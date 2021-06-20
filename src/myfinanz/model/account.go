package model

import "time"

type Account struct {
	AccountID int64 `bson:"_id,omitempty"`
	FirstName string `bson:"first_name,omitempty"`
	Lastname string `bson:"lastname,omitempty"`
	UserName string `bson:"user_name,omitempty"`
	EMail string `bson:"e_mail,omitempty"`
	Password string `bson:"password,omitempty"`
	Birthdate time.Time `bson:"birthdate,omitempty"`
	Role Role `bson:"role,omitempty"`
}

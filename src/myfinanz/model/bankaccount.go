package model

type BankAccount struct {
	BankAccountID  int64 //
	OwnerName  string // `gorm:"notNull;size:60"`
	NameOfBank string // `gorm:"notNull;size:40"`
	Iban       string // `gorm:"notNull;size:20"`
}

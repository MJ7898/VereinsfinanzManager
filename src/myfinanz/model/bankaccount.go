package model

type BankAccount struct {
	SchemaVersion string `bson:"schema_version,omitempty"`
	BankAccountID int64  `bson:"bank_account_id,omitempty"`
	OwnerName     string `bson:"owner_name,omitempty"`
	NameOfBank    string `bson:"name_of_bank,omitempty"`
	Iban          string `bson:"iban,omitempty"`
}

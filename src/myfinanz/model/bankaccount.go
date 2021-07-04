package model

type BankAccount struct {
	SchemaVersion string `bson:"schema_version,omitempty" json:"schema_version"`
	BankAccountID int64  `bson:"bank_account_id,omitempty" json:"bank_account_id"`
	OwnerName     string `bson:"owner_name,omitempty" json:"owner_name"`
	NameOfBank    string `bson:"name_of_bank,omitempty" json:"name_of_bank"`
	Iban          string `bson:"iban,omitempty" json:"iban"`
}

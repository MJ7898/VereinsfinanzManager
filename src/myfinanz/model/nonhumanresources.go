package model

import "time"

type NonHumanResources struct {
	SchemaVersion string `bson:"schema_version,omitempty"`
	// ContractID             int64     `bson:"contract_id,omitempty"`
	Name         string       `bson:"name,omitempty"`
	Cost         string       `bson:"cost,omitempty"`
	Validity     string       `bson:"validity,omitempty"`
	TimeStamp    time.Time    `bson:"time_stamp,omitempty"`
	ResourceType ResourceType `bson:"resource_type,omitempty"`
}

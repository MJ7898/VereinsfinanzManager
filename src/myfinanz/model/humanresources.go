package model

import "time"

type HumanResources struct {
	SchemaVersion string `bson:"schema_version,omitempty"`
	// ContractID             int64     `bson:"contract_id,omitempty"`
	Name                   string    `bson:"player_name,omitempty"`
	Value                  string    `bson:"value,omitempty"`
	Salary                 string    `bson:"salary,omitempty"`
	ContractRuntime        time.Time `bson:"contract_runtime,omitempty"`
	ContractRuntimeInYears float32   `bson:"contract_runtime_in_years,omitempty"`
}

package model

import "time"

type Contract struct {
	ContractID             int64
	PlayerName             string
	Value                  string
	Salary                 string
	ContractRuntime        time.Time
	ContractRuntimeInYears float32
}

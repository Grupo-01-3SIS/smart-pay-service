package entity

import "time"

type Budget struct {
	AreaName       string
	ExecutiveName  string
	Email          string
	TypeCostCenter string
	Value          float64
	StartDate      time.Time
	EndDate        time.Time
}

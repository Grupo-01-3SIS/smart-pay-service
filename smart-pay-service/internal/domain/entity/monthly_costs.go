package entity

import "time"

type MonthlyCosts struct {
	AreaName        string
	Date            time.Time
	Responsible     string
	Description     string
	Value           float64
	ExpenseCategory string
	PaymentMethod   string
	Observation     string
}

type EmployeeMonthlyCosts struct {
	AreaName     string
	EmployeeName string
	JobTitle     string
	Position     string
	Salary       string
}

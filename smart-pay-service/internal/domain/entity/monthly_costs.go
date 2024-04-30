package entity

import "time"

const FilePatternMC = `^custos-mensais.*\.csv$`
const FilePatternEMC = `custos-funcionarios.*\.csv$`

// DateFormat é um novo tipo que representa o formato de data "yyyy/MM/dd".
type DateFormat string

// UnmarshalCSV implementa a interface csv.Unmarshaler.
func (d *DateFormat) UnmarshalCSV(s string) error {
	t, err := time.Parse("2006/01/02", s)
	if err != nil {
		return err
	}
	*d = DateFormat(t.Format("2006/01/02"))
	return nil
}

// MarshalCSV implementa a interface csv.Marshaler.
func (d DateFormat) MarshalCSV() (string, error) {
	return string(d), nil
}

type MonthlyCosts struct {
	AreaName        string     `csv:"id_Cc"`
	Date            DateFormat `csv:"data"`
	Responsible     string     `csv:"responsavel_custo"`
	Description     string     `csv:"desc_transacao"`
	Value           float64    `csv:"valor"`
	ExpenseCategory string     `csv:"categoria_despesa"`
	PaymentMethod   string     `csv:"metodo_pagto"`
	Observation     string     `csv:"obs"`
}

type EmployeeMonthlyCosts struct {
	AreaName     string `csv:"id_Cc"`
	EmployeeName string `csv:"nome"`
	JobTitle     string `csv:"cargo"`
	Position     string `csv:"senioridade"`
	Salary       string `csv:"salario"`
}

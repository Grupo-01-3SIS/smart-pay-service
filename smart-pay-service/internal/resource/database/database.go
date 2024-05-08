package database

import (
	"database/sql"
	"fmt"
	"os"
	"smart-pay-service/internal/domain/entity"
	"smart-pay-service/internal/domain/gateway"
	"strconv"

	log_zap "smart-pay-service/config/log"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	"go.uber.org/zap"
)

var _ gateway.CostCenter = (*Database)(nil)

type Database struct {
	log *zap.Logger
	db  *sql.DB
}

func NewDatabase() *Database {
	return &Database{
		log: log_zap.NewLogger().Named("layer-database"),
		db:  connectionDb(),
	}
}

func (d *Database) InsertCoastCenter(obj []*entity.CostCenterInfo) (int, error) {

	var id int
	for _, info := range obj {
		err := d.db.QueryRow(`INSERT INTO smartpay.orcamento_trimestral (data_inicio, data_fim, orcamento_trimestral)
			VALUES ($1, $2, $3)
			ON CONFLICT DO NOTHING
			RETURNING idorcamento_trimestral`, info.StartDate, info.EndDate, info.Value).Scan(&id)
		if err != nil && err != sql.ErrNoRows {
			d.log.Error(err.Error())
			return 0, err
		}

		if id == 0 {
			err := d.db.QueryRow(`SELECT idorcamento_trimestral FROM smartpay.orcamento_trimestral
				WHERE data_inicio = $1 AND data_fim = $2`, info.StartDate, info.EndDate).Scan(&id)
			if err != nil {
				d.log.Error(err.Error())
				return 0, err
			}
		}

		_, err = d.db.Exec(`INSERT INTO smartpay.centro_de_custos (nome_centro, tipo, orcamento_trimestral, fk_orcamento_trimestral)
			VALUES ($1, $2, $3, $4)`, info.AreaName, info.TypeCostCenter, info.Value, id)
		if err != nil {
			d.log.Error(err.Error())
			return 0, err
		}
	}

	return id, nil
}

func (d *Database) InsertCoastEmployee(obj []*entity.EmployeeMonthlyCosts, id int) error {
	for _, info := range obj {
		_, err := d.db.Exec(`INSERT INTO smartpay.funcionario (nome_funcionarios, cargo, senioridade, salario, fk_centro_de_custos)
		VALUES($1, $2, $3, $4, $5)`, info.EmployeeName, info.JobTitle, info.Position, info.Salary, id)
		if err != nil {
			d.log.Error(err.Error())
			return err
		}
	}
	return nil
}

func (d *Database) InsertCoastVariable(obj []*entity.MonthlyCosts, id int) error {
	for _, info := range obj {
		_, err := d.db.Exec(`INSERT INTO smartpay.gastos_variaveis (valor, categoria_despesa, desc_transacao, metodo_pagto, obs, data, responsavel, fk_centro_de_custos)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`, info.Value, info.ExpenseCategory, info.Description, info.PaymentMethod, info.Observation, info.Date, info.Responsible, id)
		if err != nil {
			d.log.Error(err.Error())
			return err
		}
	}
	return nil
}

func connectionDb() *sql.DB {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Erro ao carregar arquivo .env:", err)
		panic("falha ao carregar o .env")
	}

	portInt, err := strconv.Atoi(os.Getenv("DATABASE_PORT"))
	if err != nil {
		panic(err)
	}

	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", os.Getenv("DATABASE_HOST"), portInt, os.Getenv("DATABASE_USER"), os.Getenv("DATABASE_PASSWORD"), os.Getenv("DATABASE_NAME"))

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil
	}

	// defer db.Close()
	return db
}

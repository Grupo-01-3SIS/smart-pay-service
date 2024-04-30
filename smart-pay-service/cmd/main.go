package main

import (
	"context"
	"fmt"
	"os"
	"smart-pay-service/internal/domain/entity"
	"smart-pay-service/internal/domain/service"
	s3_client "smart-pay-service/internal/s3"
	"strings"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Erro ao carregar arquivo .env:", err)
		panic("falha ao carregar o .env")
	}
	s3Client, err := s3_client.NewBucketHandler(
		context.Background(),
		os.Getenv("AWS_REGION"),
		os.Getenv("AWS_ENDPOINT"),
	)
	if err != nil {
		panic("falha ao criar o handler do bucket")
	}

	objs := s3Client.ListObjects()

	// for para fazer download dos objetos que tem no bucket
	for _, obj := range objs {
		fileName := strings.Split(obj, "/")
		err := s3Client.DownloadCsv("smartpay", obj, fileName[len(fileName)-1])
		if err != nil {
			panic("falha no download")
		}
	}

	service := service.NewCostCenterService(nil)

	filePatterns := []string{`^custos-mensais.*\.csv$`, `^custos-funcionarios.*\.csv$`, `^centro-de-custos.*\.csv$`}

	// for para dar o marshal nos 3 csv
	for _, pattern := range filePatterns {
		_, err := service.UnmarshalCSVData(context.Background(), pattern)
		if err != nil {
			panic("falha no marshal")
		}
	}
}

// Metodo para printar o resultado do Marshal do csv custos-mensais...csv
func PrintMonthlyCosts(seq []*entity.MonthlyCosts) {
	// Iterar sobre cada MonthlyCosts
	for _, monthlyCost := range seq {
		// Imprimir o valor de cada campo da struct MonthlyCosts
		fmt.Println("AreaName:", monthlyCost.AreaName)
		fmt.Println("Date:", monthlyCost.Date)
		fmt.Println("Responsible:", monthlyCost.Responsible)
		fmt.Println("Description:", monthlyCost.Description)
		fmt.Println("Value:", monthlyCost.Value)
		fmt.Println("ExpenseCategory:", monthlyCost.ExpenseCategory)
		fmt.Println("PaymentMethod:", monthlyCost.PaymentMethod)
		fmt.Println("Observation:", monthlyCost.Observation)
		fmt.Println()
	}
}

// Metodo para printar o resultado do Marshal do csv custos-funcionarios...csv
func PrintEmployeeMonthlyCosts(seq []*entity.EmployeeMonthlyCosts) {
	for _, employeeMonthlyCost := range seq {
		fmt.Println("AreaName:", employeeMonthlyCost.AreaName)
		fmt.Println("EmployeeName:", employeeMonthlyCost.EmployeeName)
		fmt.Println("JobTitle:", employeeMonthlyCost.JobTitle)
		fmt.Println("Position:", employeeMonthlyCost.Position)
		fmt.Println("Salary:", employeeMonthlyCost.Salary)
		fmt.Println()
	}
}

// Metodo para printar o resultado do Marshal do csv centro-de-custos...csv

func PrintCostCenter(seq []*entity.CostCenterInfo) {
	// Iterar sobre cada MonthlyCosts
	for _, costCenter := range seq {
		fmt.Println("AreaName:", costCenter.AreaName)
		fmt.Println("ExecutiveName:", costCenter.ExecutiveName)
		fmt.Println("Email:", costCenter.Email)
		fmt.Println("TypeCostCenter:", costCenter.TypeCostCenter)
		fmt.Println("Value:", costCenter.Value)
		fmt.Println("StartDate:", costCenter.StartDate)
		fmt.Println("EndDate:", costCenter.EndDate)
		fmt.Println()
	}
}

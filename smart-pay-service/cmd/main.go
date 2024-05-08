package main

import (
	"context"
	"fmt"
	"os"
	"smart-pay-service/internal/domain/service"
	"smart-pay-service/internal/resource/database"
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
			println(err)
			panic("falha no download")
		}
	}

	gtw := database.NewDatabase()
	service := service.NewCostCenterService(gtw)
	service.RunService(context.Background())
}

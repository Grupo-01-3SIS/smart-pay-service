package s3_client

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	log_zap "smart-pay-service/config/log"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/joho/godotenv"
	"github.com/spf13/afero"
	"go.uber.org/zap"
)

var mesesMap = map[string]string{
	"january":   "janeiro",
	"february":  "fevereiro",
	"march":     "março",
	"april":     "abril",
	"may":       "maio",
	"june":      "junho",
	"july":      "julho",
	"august":    "agosto",
	"september": "setembro",
	"october":   "outubro",
	"november":  "novembro",
	"december":  "dezembro",
}

type S3Handler struct {
	region string
	client *s3.Client
	log    *zap.Logger
}

type keys struct {
	keyId     string
	secretKey string
}

func NewBucketHandler(ctx context.Context, region string, endpoint string) (*S3Handler, error) {
	keys, err := getKeys()
	if err != nil {
		return nil, err
	}
	cfg, err := config.LoadDefaultConfig(
		context.TODO(),
		config.WithRegion(region),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(
			keys.keyId,
			keys.secretKey, "",
		)),
		config.WithEndpointResolverWithOptions(aws.EndpointResolverWithOptionsFunc(
			func(service, region string, options ...interface{}) (aws.Endpoint, error) {
				return aws.Endpoint{
					PartitionID:   "aws",
					URL:           endpoint,
					SigningRegion: region,
				}, nil
			})),
	)
	if err != nil {
		log.Fatal(err)
	}

	return &S3Handler{
		region: region,
		client: s3.NewFromConfig(cfg),
		log:    log_zap.NewLogger().Named("layer-s3client"),
	}, nil
}

func (h *S3Handler) ListObjects() []string {
	h.log.Info("List objects...")
	output, err := h.client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
		Bucket: aws.String("smartpay"),
	})

	if err != nil {
		h.log.Error(err.Error())
		log.Fatal(err)
	}

	response := make([]string, 0, len(output.Contents))
	for _, object := range output.Contents {
		response = append(response, aws.ToString(object.Key))
	}

	response = currentMonthCsv(response)
	h.log.Info(fmt.Sprintf("response List Object %v", response))
	return response
}

func (h *S3Handler) DownloadCsv(bucketName string, objectKey string, fileName string) error {
	var aferoFs = afero.NewOsFs()
	result, err := h.client.GetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectKey),
	})
	if err != nil {
		log.Printf("Couldn't get object %v:%v. Here's why: %v\n", bucketName, objectKey, err)
		return err
	}
	defer result.Body.Close()

	nameDir := strings.Split(objectKey, "/")[0]
	existsDir, _ := afero.Exists(aferoFs, fmt.Sprintf("tools/%s", nameDir))

	if !existsDir {
		err = aferoFs.Mkdir(fmt.Sprintf("tools/%s", nameDir), 0755)
		if err != nil {
			log.Printf("Couldn't create dir %s error: %s", nameDir, err)
			return err
		}
	}

	file, err := aferoFs.Create(fmt.Sprintf("tools/%s/%s", nameDir, fileName))
	if err != nil {
		log.Printf("Couldn't create file %v. Here's why: %v\n", fileName, err)
		return err
	}
	defer file.Close()
	body, err := io.ReadAll(result.Body)
	if err != nil {
		log.Printf("Couldn't read object body from %v. Here's why: %v\n", objectKey, err)
	}
	_, err = file.Write(body)
	return err
}

func getKeys() (*keys, error) {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Erro ao carregar arquivo .env:", err)
		return nil, err
	}

	return &keys{
		keyId:     os.Getenv("AWS_ACCESS_KEY_ID"),
		secretKey: os.Getenv("AWS_SECRET_ACCESS_KEY"),
	}, nil
}

func currentMonthCsv(objetos []string) []string {
	currentMonth := strings.ToLower(time.Now().Month().String())

	currentMonthPortuguese, ok := mesesMap[currentMonth]
	if !ok {
		fmt.Println("Nome do mês não encontrado:", currentMonth)
		return nil
	}

	var csvMonth []string
	for _, objeto := range objetos {
		dir := strings.Split(objeto, "/")[0]

		if strings.HasSuffix(strings.ToLower(dir), currentMonthPortuguese) {
			if strings.HasSuffix(strings.ToLower(objeto), ".csv") {
				csvMonth = append(csvMonth, objeto)
			}
		}
	}

	return csvMonth
}

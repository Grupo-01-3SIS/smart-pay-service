package s3_client

import (
	"context"
	"fmt"
	"log"
	"os"
	log_zap "smart-pay-service/config/log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

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
		h.log.Info(fmt.Sprintf("key=%s size=%d", aws.ToString(object.Key), object.Size))
		response = append(response, aws.ToString(object.Key))
	}

	return response
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

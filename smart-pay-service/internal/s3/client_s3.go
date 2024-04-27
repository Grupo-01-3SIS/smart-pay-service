package s3_client

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/joho/godotenv"
)

type S3Handler struct {
	region string
	client *s3.Client
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
					SigningRegion: "us-east-1",
				}, nil
			})),
	)
	if err != nil {
		log.Fatal(err)
	}

	return &S3Handler{
		region: region,
		client: s3.NewFromConfig(cfg),
	}, nil

}

func (h *S3Handler) ListObjects() {
	output, err := h.client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
		Bucket: aws.String("smartpay"),
	})

	if err != nil {
		log.Fatal(err)
	}

	log.Println("first page results:")
	for _, object := range output.Contents {
		log.Printf("key=%s size=%d", aws.ToString(object.Key), object.Size)
	}
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

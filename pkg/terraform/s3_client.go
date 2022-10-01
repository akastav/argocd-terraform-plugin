package ke_terraform

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func main() {
	var bucketName = os.Getenv("S3_BUCKET")
	var accessKeyId = os.Getenv("S3_ACCESS_KEY_ID")
	var accessKeySecret = os.Getenv("S3_SECRET_ACCESS_KEY")

	// Get findKey from command -f
	findKey := flag.String("f", "", "The key of bucket object to search")
	flag.Parse()

	if *findKey == "" {
		fmt.Println("You must supply the key of a bucket (-f key of object)")
		return
	}

	// Create custom endpoint resolver for returning correct URL for S3 storage in ru-central1 region
	customResolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		return aws.Endpoint{
			PartitionID:   "yc",
			URL:           "https://storage.yandexcloud.net",
			SigningRegion: "ru-central1",
		}, nil

	})

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithEndpointResolverWithOptions(customResolver),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(accessKeyId, accessKeySecret, "")))
	if err != nil {
		log.Fatal(err)
	}

	// Create S3 client
	s3client := s3.NewFromConfig(cfg)

	getBucketObject(*s3client, bucketName, *findKey)
}

func getBucketObject(client s3.Client, bucketName string, findKey string) {

	// Get Bucket list
	result, err := client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
		Bucket: aws.String(bucketName),
	})
	if err != nil {
		log.Fatal(err)
	}

	// // Get Bucket object
	// for _, object := range result.Contents {
	// 	log.Printf("object=%s size=%d Bytes last modified=%s", aws.ToString(object.Key), object.Size, object.LastModified.Format("2006-01-02 15:04:05 Monday"))
	// }

	// Find requested object
	for _, object := range result.Contents {
		if aws.ToString(object.Key) == findKey {
			log.Printf("object=%s size=%d Bytes last modified=%s", aws.ToString(object.Key), object.Size, object.LastModified.Format("2006-01-02 15:04:05 Monday"))
		}
	}
}

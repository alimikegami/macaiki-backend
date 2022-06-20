package cloudstorage

import (
	"fmt"
	"mime/multipart"
	"path/filepath"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type S3 struct {
	AwsAccessKey string
	AwsSecretKey string
	Region       string
	BucketName   string
}

func CreateNewS3Instance(accessKey string, secretKey string, region string, bucketName string) *S3 {
	return &S3{AwsAccessKey: accessKey, AwsSecretKey: secretKey, Region: region, BucketName: bucketName}
}

func (s *S3) CreateAWSSession() (*session.Session, error) {
	sess, err := session.NewSession(
		&aws.Config{
			Region: aws.String(s.Region),
			Credentials: credentials.NewStaticCredentials(
				s.AwsAccessKey,
				s.AwsSecretKey,
				"",
			),
		})

	if err != nil {
		fmt.Println("AWS session was not created successfully...")
	}

	return sess, err
}

func (s *S3) UploadImage(fileName string, dirName string, img *multipart.FileHeader) (*s3manager.UploadOutput, error) {
	src, err := img.Open()
	if err != nil {
		return nil, err
	}
	defer src.Close()

	sess := session.Must(s.CreateAWSSession())

	// Create an uploader with the session and default options
	uploader := s3manager.NewUploader(sess)

	// Upload the file to S3.
	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(s.BucketName),
		Key:    aws.String(dirName + "/" + fileName + filepath.Ext(img.Filename)),
		Body:   src,
		ACL:    aws.String("public-read"),
	})

	return result, err
}

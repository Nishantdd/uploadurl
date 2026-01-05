package service

import (
	"io"

	"github.com/Nishantdd/uploadurl/backend/config"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type S3Service struct {
	bucket  string
	session *session.Session
	active  bool
}

var S3Client *S3Service = &S3Service{active: false}
var cfg *config.Config = config.Load()
var (
	bucket    = cfg.AWS.S3BucketName
	accessKey = cfg.AWS.AccessKey
	secretKey = cfg.AWS.SecretKey
	region    = cfg.AWS.Region
)

func NewS3Client() (*S3Service, error) {
	if S3Client.active {
		return S3Client, nil
	}

	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(region),
		Credentials: credentials.NewStaticCredentials(accessKey, secretKey, ""),
	})
	if err != nil {
		return nil, err
	}

	S3Client = &S3Service{bucket: bucket, session: sess, active: true}
	return S3Client, nil
}

func (s *S3Service) UploadFile(file io.Reader, fileName string) (string, error) {
	uploader := s3manager.NewUploader(s.session)

	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(fileName),
		Body:   file,
	})
	if err != nil {
		return "", err
	}

	return result.Location, nil
}

func (s *S3Service) DeleteFile(fileName string) error {
	svc := s3.New(s.session)
	_, err := svc.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(fileName),
	})
	if err != nil {
		return err
	}

	return svc.WaitUntilObjectNotExists(&s3.HeadObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(fileName),
	})
}

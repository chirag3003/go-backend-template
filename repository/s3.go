package repository

import (
	"context"
	"io"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	awsS3 "github.com/chirag3003/go-backend-template/helpers/aws"
)

type s3Repository struct {
	s3Uploader   *manager.Uploader
	s3Downloader *manager.Downloader
}

type S3Repository interface {
	Upload(context context.Context, key string, file io.Reader) (*manager.UploadOutput, error)
}

func NewS3Repository() S3Repository {
	return &s3Repository{
		s3Uploader:   awsS3.GetS3Uploader(),
		s3Downloader: awsS3.GetS3Downloader(),
	}
}

func (r *s3Repository) Upload(context context.Context, key string, file io.Reader) (*manager.UploadOutput, error) {
	res, err := r.s3Uploader.Upload(context, &s3.PutObjectInput{
		Bucket: aws.String(os.Getenv("S3_BUCKET")),
		Key:    aws.String(key),
		Body:   file,
		ACL:    "public-read",
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

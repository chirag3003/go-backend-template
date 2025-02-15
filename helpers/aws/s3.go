package aws

// var s3Client *
import "github.com/aws/aws-sdk-go-v2/service/s3"
import "github.com/aws/aws-sdk-go-v2/feature/s3/manager"

var s3Client *s3.Client
var s3Uploader *manager.Uploader
var s3ownloader *manager.Downloader

func setupS3() {
	client := s3.NewFromConfig(awsConfig)
	s3Client = client

	uploader := manager.NewUploader(client)
	downloader := manager.NewDownloader(client)

	s3Uploader = uploader
	s3ownloader = downloader
}

func GetS3Client() *s3.Client {
	return s3Client
}

func GetS3Uploader() *manager.Uploader {
	return s3Uploader
}

func GetS3Downloader() *manager.Downloader {
	return s3ownloader
}

package utils

import (
	"bytes"
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

var s *session.Session

func S3Init() (err error) {
	s3Config := &aws.Config{
		Region: aws.String(os.Getenv("AWS_BUCKET_REGION")),
		Credentials: credentials.NewStaticCredentials(
			os.Getenv("AWS_KEY_ID"),
			os.Getenv("AWS_KEY_SECRET"),
			""),
	}

	s, err = session.NewSession(s3Config)

	return
}

func S3Client(file *multipart.FileHeader, name string) (err error, location string) {
	err, fileBytes := getFileBytes(file)

	objectInput := &s3manager.UploadInput{
		Bucket:      aws.String(os.Getenv("AWS_BUCKET_NAME")),
		Key:         aws.String(name),
		Body:        bytes.NewReader(fileBytes),
		ContentType: aws.String(http.DetectContentType(fileBytes)),
	}

	ctx := context.Background()

	objectOutput, err := s3manager.NewUploader(s).UploadWithContext(ctx, objectInput)

	location = objectOutput.Location
	return
}

func getFileBytes(file *multipart.FileHeader) (err error, fileBytes []byte) {
	fileReader, err := file.Open()
	if err != nil {
		return
	}

	buf := bytes.NewBuffer(nil)
	_, err = io.Copy(buf, fileReader)
	if err != nil {
		return
	}

	fileBytes = buf.Bytes()
	return
}

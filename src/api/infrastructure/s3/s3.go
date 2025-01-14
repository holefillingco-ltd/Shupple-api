package s3

import (
	"bytes"
	"encoding/base64"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/kelseyhightower/envconfig"
	"github.com/uma-co82/Shupple-api/src/api/domain"
)

type (
	S3Service struct{}
	// NOTE: 環境変数にS3AKとS3SKを設定
	Env struct {
		S3AK string
		S3SK string
	}
)

func (s S3Service) UploadToS3(image string, uid string) error {
	// 環境変数からS3Credential周りの設定を取得
	var env Env
	_ = envconfig.Process("", &env)

	sess := session.Must(session.NewSession(&aws.Config{
		Credentials: credentials.NewStaticCredentials(env.S3AK, env.S3SK, ""),
		Region:      aws.String("ap-northeast-1"),
	}))

	uploader := s3manager.NewUploader(sess)

	data, _ := base64.StdEncoding.DecodeString(image)
	wb := new(bytes.Buffer)
	wb.Write(data)

	_, err := uploader.Upload(&s3manager.UploadInput{
		Bucket:      aws.String("shupple-user-images"),
		Key:         aws.String("images/" + uid + ".png"),
		Body:        wb,
		ContentType: aws.String("image/png"),
	})

	if err != nil {
		if err, ok := err.(awserr.Error); ok && err.Code() == request.CanceledErrorCode {
			return domain.RaiseError(400, "Upload TimuOut", nil)
		} else {
			return domain.RaiseError(400, "Upload Failed", nil)
		}
	}

	return nil
}

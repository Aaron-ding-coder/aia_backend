package handler

import (
	"context"
	"fmt"
	"github.com/volcengine/ve-tos-golang-sdk/v2/tos"
	"io"
)

var (
	ak = ""
	sk = ""
	// endpoint 若没有指定 HTTP 协议（HTTP/HTTPS），默认使用 HTTPS
	endpoint   = "tos-cn-beijing.volces.com"
	region     = "cn-beijing"
	client     *tos.ClientV2
	bucketName = "dingzwtest"
)

func InitObjectStorage() {
	credential := tos.NewStaticCredentials(ak, sk)
	var err error
	client, err = tos.NewClientV2(endpoint, tos.WithCredentials(credential), tos.WithRegion(region))
	if err != nil {
		fmt.Println("Error:", err)
		panic(err)
	}
}

func upload(ctx context.Context, ObjectKey string, f io.ReadCloser) error {
	_, err := client.PutObjectV2(ctx, &tos.PutObjectV2Input{
		PutObjectBasicInput: tos.PutObjectBasicInput{
			Bucket: bucketName,
			Key:    ObjectKey,
		},
		Content: f,
	})

	return err
}

func download(ctx context.Context, ObjectKey string) (io.ReadCloser, error) {
	output, err := client.GetObjectV2(ctx, &tos.GetObjectV2Input{
		Bucket: bucketName,
		Key:    ObjectKey,
	})

	if err != nil {
		return nil, err
	}

	return output.Content, nil
}

package main

import (
	"log"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"

	"github.com/kaperys/awslocal"
)

func main() {
	endpoint, _ := awslocal.Endpoint(awslocal.ServiceS3)
	cfg := awslocal.NewConfig(
		awslocal.WithRegion("eu-west-2"),
		awslocal.WithEndpoint(endpoint),
		awslocal.WithDummyCredentials(),
	)

	aws, err := session.NewSession(cfg)
	if err != nil {
		log.Fatal(err)
	}

	svc := s3.New(aws, nil)
	log.Println(svc.ListBuckets(&s3.ListBucketsInput{}))
}

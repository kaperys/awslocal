package main

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"

	"github.com/kaperys/awslocal"
)

func main() {
	session, err := session.NewSession(aws.NewConfig())
	if err != nil {
		log.Fatal(err)
	}

	cfg := aws.NewConfig()
	awslocal.Wrap(cfg, awslocal.ServiceS3)

	svc := s3.New(session, cfg)
	log.Println(svc.ListBuckets(&s3.ListBucketsInput{}))
}

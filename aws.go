// Copyright (c) 2018, Mike Kaperys <mike@kaperys.io>
// See LICENSE for licensing information.

// Package awslocal provides a LocalStack wrapper for the AWS Go SDK.
package awslocal

import "github.com/aws/aws-sdk-go/aws"

// A Service is a reference to a supported LocalStack service.
type Service byte

const (
	// ServiceAPIGateway is a constant value used to identify the LocalStack API Gateway service.
	ServiceAPIGateway Service = iota

	// ServiceKinesis is a constant value used to identify the LocalStack Kinesis service.
	ServiceKinesis

	// ServiceDynamoDB is a constant value used to identify the LocalStack DynamoDB service.
	ServiceDynamoDB

	// ServiceDynamoDBStreams is a constant value used to identify the LocalStack DynamoDB Streams service.
	ServiceDynamoDBStreams

	// ServiceElasticsearch is a constant value used to identify the LocalStack Elasticsearch service.
	ServiceElasticsearch

	// ServiceS3 is a constant value used to identify the LocalStack S3 service.
	ServiceS3

	// ServiceFirehose is a constant value used to identify the LocalStack Firehose service.
	ServiceFirehose

	// ServiceLambda is a constant value used to identify the LocalStack Lambda service.
	ServiceLambda

	// ServiceSNS is a constant value used to identify the LocalStack Simple Notification Service.
	ServiceSNS

	// ServiceSQS is a constant value used to identify the LocalStack Simple Queue Service.
	ServiceSQS

	// ServiceRedshift is a constant value used to identify the LocalStack Redshift service.
	ServiceRedshift

	// ServiceElasticsearchService is a constant value used to identify the LocalStack Elasticsearch service.
	ServiceElasticsearchService

	// ServiceSES is a constant value used to identify the LocalStack Simple Email Service.
	ServiceSES

	// ServiceRoute53 is a constant value used to identify the LocalStack Route 53 service.
	ServiceRoute53

	// ServiceCloudFormation is a constant value used to identify the LocalStack Cloud Formation service.
	ServiceCloudFormation

	// ServiceCloudWatch is a constant value used to identify the LocalStack Cloud Watch service.
	ServiceCloudWatch

	// ServiceSSM is a constant value used to identify the LocalStack Simple Service Manager service.
	ServiceSSM

	// ServiceSecretsManager is a constant value used to identify the LocalStack Secrets Manager service.
	ServiceSecretsManager
)

// Wrap wraps the given aws.Config with the LocalStack Endpoint of the service.
func Wrap(c *aws.Config, svc Service) {
	se := map[Service]string{
		ServiceAPIGateway:           "http://localhost:4567",
		ServiceKinesis:              "http://localhost:4568",
		ServiceDynamoDB:             "http://localhost:4569",
		ServiceDynamoDBStreams:      "http://localhost:4570",
		ServiceElasticsearch:        "http://localhost:4571",
		ServiceS3:                   "http://localhost:4572",
		ServiceFirehose:             "http://localhost:4573",
		ServiceLambda:               "http://localhost:4574",
		ServiceSNS:                  "http://localhost:4575",
		ServiceSQS:                  "http://localhost:4576",
		ServiceRedshift:             "http://localhost:4577",
		ServiceElasticsearchService: "http://localhost:4578",
		ServiceSES:                  "http://localhost:4579",
		ServiceRoute53:              "http://localhost:4580",
		ServiceCloudFormation:       "http://localhost:4581",
		ServiceCloudWatch:           "http://localhost:4582",
		ServiceSSM:                  "http://localhost:4583",
		ServiceSecretsManager:       "http://localhost:4584",
	}

	if endpoint, ok := se[svc]; ok {
		c.Endpoint = aws.String(endpoint)
	}
}

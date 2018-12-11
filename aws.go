// Copyright (c) 2018, Mike Kaperys <mike@kaperys.io>
// See LICENSE for licensing information.

// Package awslocal provides configuration for the AWS Go SDK for use with LocalStack.
package awslocal

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/pkg/errors"
)

var (
	// CredentialsID is the static ID used by the AWS SDK when calling LocalStack services.
	// LocalStack currently ignores all authentication mechanisms (ie, it will accept any), however
	// the AWS SDK enforces that credentials, in some form (profile, IAM role, static, etc), are
	// provided. See: https://github.com/localstack/localstack/issues/62#issuecomment-294749459
	CredentialsID = "dummy-localstack-credentials-id"

	// CredentialsSecret is the static secret used by the AWS SDK when calling LocalStack services.
	// LocalStack currently ignores all authentication mechanisms (ie, it will accept any), however
	// the AWS SDK enforces that credentials, in some form (profile, IAM role, static, etc), are
	// provided. See: https://github.com/localstack/localstack/issues/62#issuecomment-294749459
	CredentialsSecret = "dummy-localstack-credentials-secret"

	// CredentialsToken is the static token used by the AWS SDK when calling LocalStack services.
	// LocalStack currently ignores all authentication mechanisms (ie, it will accept any), however
	// the AWS SDK enforces that credentials, in some form (profile, IAM role, static, etc), are
	// provided. See: https://github.com/localstack/localstack/issues/62#issuecomment-294749459
	CredentialsToken = "dummy-localstack-credentials-token"
)

const (
	// ServiceAPIGateway is a constant value used to identify the LocalStack API Gateway service.
	ServiceAPIGateway = iota

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

// NewConfig returns a pointer to a new aws.Config. NewConfig accepts many Option funcs, which configure
// the aws.Config type. If one or more Options are provided, they are iterated and executed.
func NewConfig(opts ...Option) *aws.Config {
	cfg := new(aws.Config)

	if opts != nil {
		for _, o := range opts {
			o(cfg)
		}
	}

	return cfg
}

// An Option func configures the aws.Config type with a particular value. Option types should only be
// concerned with one particular configuration value, such as region or endpoint.
type Option func(c *aws.Config)

// WithRegion specifies the region to be used when calling services.
func WithRegion(region string) Option {
	return func(c *aws.Config) {
		c.Region = aws.String(region)
	}
}

// WithEndpoint overrides the service endpoint used when calling services.
func WithEndpoint(endpoint string) Option {
	return func(c *aws.Config) {
		c.Endpoint = aws.String(endpoint)
	}
}

// WithDummyCredentials configures the Config with dummy static credentials. Credentials are ignored by LocalStack,
// however the Go AWS SDK endforces that credentials, in some form, are provided.
func WithDummyCredentials() Option {
	return func(c *aws.Config) {
		c.Credentials = credentials.NewStaticCredentials(CredentialsID, CredentialsSecret, CredentialsToken)
	}
}

// Endpoint retrieves the LocalStack service endpoint.
func Endpoint(svc int) (string, error) {
	se := map[int]string{
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

	endpoint, ok := se[svc]
	if !ok {
		return "", errors.New("service endpoint not found")
	}

	return endpoint, nil
}

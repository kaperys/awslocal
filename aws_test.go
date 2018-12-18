package awslocal_test

import (
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/kaperys/awslocal"
)

func TestEndpointSuccess(t *testing.T) {
	tt := []struct {
		Service  awslocal.Service
		Endpoint string
	}{
		{
			Service:  awslocal.ServiceAPIGateway,
			Endpoint: "http://localhost:4567",
		},
		{
			Service:  awslocal.ServiceKinesis,
			Endpoint: "http://localhost:4568",
		},
		{
			Service:  awslocal.ServiceDynamoDB,
			Endpoint: "http://localhost:4569",
		},
		{
			Service:  awslocal.ServiceDynamoDBStreams,
			Endpoint: "http://localhost:4570",
		},
		{
			Service:  awslocal.ServiceElasticsearch,
			Endpoint: "http://localhost:4571",
		},
		{
			Service:  awslocal.ServiceS3,
			Endpoint: "http://localhost:4572",
		},
		{
			Service:  awslocal.ServiceFirehose,
			Endpoint: "http://localhost:4573",
		},
		{
			Service:  awslocal.ServiceLambda,
			Endpoint: "http://localhost:4574",
		},
		{
			Service:  awslocal.ServiceSNS,
			Endpoint: "http://localhost:4575",
		},
		{
			Service:  awslocal.ServiceSQS,
			Endpoint: "http://localhost:4576",
		},
		{
			Service:  awslocal.ServiceRedshift,
			Endpoint: "http://localhost:4577",
		},
		{
			Service:  awslocal.ServiceElasticsearchService,
			Endpoint: "http://localhost:4578",
		},
		{
			Service:  awslocal.ServiceSES,
			Endpoint: "http://localhost:4579",
		},
		{
			Service:  awslocal.ServiceRoute53,
			Endpoint: "http://localhost:4580",
		},
		{
			Service:  awslocal.ServiceCloudFormation,
			Endpoint: "http://localhost:4581",
		},
		{
			Service:  awslocal.ServiceCloudWatch,
			Endpoint: "http://localhost:4582",
		},
		{
			Service:  awslocal.ServiceSSM,
			Endpoint: "http://localhost:4583",
		},
		{
			Service:  awslocal.ServiceSecretsManager,
			Endpoint: "http://localhost:4584",
		},
	}

	for _, tc := range tt {
		cfg := aws.NewConfig()
		awslocal.Wrap(cfg, tc.Service)

		if *cfg.Endpoint != tc.Endpoint {
			t.Fatalf("TestEndpointSuccess failed: have %q, want %q", *cfg.Endpoint, tc.Endpoint)
		}
	}
}

func TestEndpointFail(t *testing.T) {
	cfg := aws.NewConfig()
	awslocal.Wrap(cfg, 222)

	if cfg.Endpoint != nil {
		t.Fatalf("TestEndpointFail failed: expected endpoint to not be set, found %q", *cfg.Endpoint)
	}
}

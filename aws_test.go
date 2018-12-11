package awslocal_test

import (
	"reflect"
	"testing"

	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/kaperys/awslocal"
)

func TestNewConfigWithRegion(t *testing.T) {
	region := "eu-west-2"
	cfg := awslocal.NewConfig(awslocal.WithRegion(region))
	if r := *cfg.Region; r != region {
		t.Fatalf("TestNewConfigWithRegion failed: have %q, want %q", r, region)
	}
}

func TestNewConfigWithEndpoint(t *testing.T) {
	tt := []struct {
		Service  int
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
		endpoint, _ := awslocal.Endpoint(tc.Service)
		cfg := awslocal.NewConfig(awslocal.WithEndpoint(endpoint))
		if e := *cfg.Endpoint; e != tc.Endpoint {
			t.Fatalf("TestNewConfigWithEndpoint failed: have %q, want %q", e, tc.Endpoint)
		}
	}
}

func TestNewConfigWithDummyCredentials(t *testing.T) {
	creds := credentials.NewStaticCredentials(awslocal.CredentialsID, awslocal.CredentialsSecret, awslocal.CredentialsToken)

	cfg := awslocal.NewConfig(awslocal.WithDummyCredentials())

	actual, _ := cfg.Credentials.Get()
	expected, _ := creds.Get()

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("TestNewConfigWithDummyCredentials failed: have %+v, want %+v", actual, expected)
	}
}

func TestEndpointSuccess(t *testing.T) {
	tt := []struct {
		Service  int
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
		endpoint, err := awslocal.Endpoint(tc.Service)
		if err != nil {
			t.Fatalf("TestEndpointSuccess failed: %v", err)
		}

		if endpoint != tc.Endpoint {
			t.Fatalf("TestEndpointSuccess failed: have %q, want %q", endpoint, tc.Endpoint)
		}
	}
}

func TestEndpointFail(t *testing.T) {
	_, err := awslocal.Endpoint(9999)
	if err == nil {
		t.Fatalf("TestEndpointFail failed: expected an error")
	}

	x := "service endpoint not found"
	if e := err.Error(); e != x {
		t.Fatalf("TestEndpointFail failed: have %q, want %q", e, x)
	}
}

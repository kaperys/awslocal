# awslocal [![Go Report Card](https://goreportcard.com/badge/github.com/kaperys/awslocal)](https://goreportcard.com/report/github.com/kaperys/awslocal)

`awslocal` is a wrapper for the Golang AWS SDK which makes interacting with [LocalStack](https://github.com/localstack/localstack) services easier. Simply wrap your service configuration with `awslocal.Wrap()`.

## Example

```go
session, err := session.NewSession(aws.NewConfig())
if err != nil {
    log.Fatal(err)
}

cfg := aws.NewConfig()

// Wrap the service cfg for use with LocalStack
awslocal.Wrap(cfg, awslocal.ServiceS3)

svc := s3.New(session, cfg)
log.Println(svc.ListBuckets(&s3.ListBucketsInput{}))
```

# awslocal

`awslocal` is a wrapper for the Golang AWS SDK which makes interacting with [LocalStack](https://github.com/localstack/localstack) services easier. Simply replace your `aws.NewConfig()` with `awslocal.NewConfig()`.

## Example

```go
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
```

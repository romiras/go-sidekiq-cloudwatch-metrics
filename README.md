# go-sidekiq-cloudwatch-metrics

Send Sidekiq metrics to AWS CloudWatch. Executed as AWS Lambda function periodically by rate 1 minute.

## Building

`GOOS=linux go build -o main main.go`

## Prerequisites

Check your current user by running `aws sts get-caller-identity`.

### Ubuntu

Install package `golang-docker-credential-helpers`.

## AWS Serverless Application Model (SAM)

Generated with

```shell
sam init
```

### Running locally

Go to directory *lambda* and run `sam local invoke`.

If you encounter error `fork/exec /var/task/main: no such file or directory: PathError` then you need refer to Building section above.

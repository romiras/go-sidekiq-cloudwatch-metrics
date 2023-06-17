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

## References

* [Use Amazon CloudWatch metrics](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/working_with_metrics.html)
* [Invoking Lambda functions locally](https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/serverless-sam-cli-using-invoke.html)
* [Example: Using custom Amazon CloudWatch metrics](https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/customize-containers-cw.html)
* [Building Lambda functions with Go](https://docs.aws.amazon.com/lambda/latest/dg/lambda-golang.html)
* [Blank function in Go](https://github.com/awsdocs/aws-lambda-developer-guide/tree/main/sample-apps/blank-go)
* [AWS Lambda function logging in Go](https://docs.aws.amazon.com/lambda/latest/dg/golang-logging.html)
* [Using the AWS SAM CLI to upload local files at deployment](https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/deploy-upload-local-files.html)

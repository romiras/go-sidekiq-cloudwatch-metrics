package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
	"github.com/go-redis/redis/v8"
)

var (
	redisClient      *redis.Client
	redisUrl         = os.Getenv("REDIS_URL")  // redis://:qwerty@localhost:6379/0
	queueName        = os.Getenv("QUEUE_NAME") // "sidekiq_namespace:queue:foo"
	cloudwatchClient *cloudwatch.CloudWatch
)

const (
	MetricNamespace = "DemoNamespaceA"
)

func init() {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	opt, err := redis.ParseURL(redisUrl)
	if err != nil {
		panic(err)
	}

	redisClient = redis.NewClient(opt)
	cloudwatchClient = cloudwatch.New(sess)
}

func main() {
	lambda.Start(handler)
}

func handler(ctx context.Context) error {
	queueLength, err := redisClient.LLen(ctx, queueName).Result()
	if err != nil {
		return err
	}

	if err = redisClient.Close(); err != nil {
		return err
	}

	if err = sendMetric(ctx, MetricNamespace, queueLength); err != nil {
		return err
	}

	fmt.Printf("Queue length is %d\n", queueLength)

	return nil
}

func sendMetric(ctx context.Context, metricNamespace string, queueLength int64) error {
	_, err := cloudwatchClient.PutMetricDataWithContext(ctx, &cloudwatch.PutMetricDataInput{
		Namespace: aws.String(metricNamespace),
		MetricData: []*cloudwatch.MetricDatum{
			{
				MetricName: aws.String("queue_length"),
				Timestamp:  aws.Time(time.Now()),
				Value:      aws.Float64(float64(queueLength)),
				Unit:       aws.String(cloudwatch.StandardUnitCount),
			},
		},
	})
	return err
}

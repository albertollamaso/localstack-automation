package awslocalstack

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

func NewAWSSession() *session.Session {
	// Initialize a session
	sess, err := session.NewSession(&aws.Config{
		Region:           aws.String(os.Getenv("AWS_DEFAULT_REGION")),
		Credentials:      credentials.NewStaticCredentials(os.Getenv("AWS_ACCESS_KEY_ID"), os.Getenv("AWS_SECRET_ACCESS_KEY"), os.Getenv("AWS_TOKEN")),
		S3ForcePathStyle: aws.Bool(true),
		Endpoint:         aws.String(os.Getenv("LOCALSTACK_ENDPOINT")),
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	return (sess)
}

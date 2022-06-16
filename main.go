package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/albertollamaso/localstack-automation/awslocalstack"
	"github.com/albertollamaso/localstack-automation/common"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	"github.com/joho/godotenv"
)

var (
	TotalSecrets int
	sess         *session.Session
)

func init() {

	// load env vars
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	// Parse flags
	flag.IntVar(&TotalSecrets, "awssecretstotal", 0, "Total number of AWS secrets to generate on localstack")
	flag.Bool("awssecretslist", false, "List all secrets from AWS Secret Manager from localstack")
	flag.Parse()

	// create AWS session with localstack
	sess = awslocalstack.NewAWSSession()

}

func AWSSecretManager() {

	// Create service
	svc := secretsmanager.New(sess)

	// Create secrets
	if TotalSecrets > 0 {

		for i := 1; i <= TotalSecrets; i++ {
			awslocalstack.CreateSecrets(svc)
		}

		fmt.Println("Total AWS Secrets created: ", TotalSecrets)
	}

	// List secrets
	flagset := common.IsFlagPassed("awssecretslist")

	if flagset {
		total := awslocalstack.ListSecrets(svc)
		fmt.Println("#-----------------------------------------")
		fmt.Println("Total AWS Secrets in localstack: ", total)
		fmt.Println("#-----------------------------------------")
	}

}

func main() {
	// Run AWS Secret Manager
	AWSSecretManager()
}

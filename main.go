package main

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	"github.com/joho/godotenv"
)

func ListSecrets(svc *secretsmanager.SecretsManager) {

	input := &secretsmanager.ListSecretsInput{}
	result, err := svc.ListSecrets(input)

	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case secretsmanager.ErrCodeInvalidParameterException:
				fmt.Println(secretsmanager.ErrCodeInvalidParameterException, aerr.Error())
			case secretsmanager.ErrCodeInvalidNextTokenException:
				fmt.Println(secretsmanager.ErrCodeInvalidNextTokenException, aerr.Error())
			case secretsmanager.ErrCodeInternalServiceError:
				fmt.Println(secretsmanager.ErrCodeInternalServiceError, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
	}

	for _, secret := range result.SecretList {
		name := *secret.Name
		arn := *secret.ARN
		createdDate := *secret.CreatedDate
		lastChangedDate := *secret.LastChangedDate

		fmt.Println("#--------------------------------------------------------------------------")
		fmt.Println("Secret Name:", name)
		fmt.Println("ARN:", arn)
		fmt.Println("CreatedDate:", createdDate)
		fmt.Println("LastChangedDate:", lastChangedDate)

	}

}

func main() {

	// load env vars
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	// sess := awslocalstack.
	svc := secretsmanager.New(sess)
	ListSecrets(svc)
}

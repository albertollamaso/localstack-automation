package awslocalstack

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
)

func ListSecrets(svc *secretsmanager.SecretsManager) int {

	var max int64 = 9223372036854775807

	input := &secretsmanager.ListSecretsInput{
		MaxResults: &max,
	}

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

	total := 0

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

		total++

	}
	return (total)

}

func CreateSecrets(svc *secretsmanager.SecretsManager) {

	name := RandString(10)
	username := RandString(10)
	password := RandString(10)
	secretString := fmt.Sprintf("{\"username\":\"%s\",\"password\":\"%s\"}", username, password)

	input := &secretsmanager.CreateSecretInput{
		ClientRequestToken: aws.String("EXAMPLE1-90ab-cdef-fedc-ba987SECRET1"),
		Description:        aws.String("My secret"),
		Name:               aws.String(name),
		SecretString:       aws.String(secretString),
	}

	result, err := svc.CreateSecret(input)

	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case secretsmanager.ErrCodeInvalidParameterException:
				fmt.Println(secretsmanager.ErrCodeInvalidParameterException, aerr.Error())
			case secretsmanager.ErrCodeInvalidRequestException:
				fmt.Println(secretsmanager.ErrCodeInvalidRequestException, aerr.Error())
			case secretsmanager.ErrCodeLimitExceededException:
				fmt.Println(secretsmanager.ErrCodeLimitExceededException, aerr.Error())
			case secretsmanager.ErrCodeEncryptionFailure:
				fmt.Println(secretsmanager.ErrCodeEncryptionFailure, aerr.Error())
			case secretsmanager.ErrCodeResourceExistsException:
				fmt.Println(secretsmanager.ErrCodeResourceExistsException, aerr.Error())
			case secretsmanager.ErrCodeResourceNotFoundException:
				fmt.Println(secretsmanager.ErrCodeResourceNotFoundException, aerr.Error())
			case secretsmanager.ErrCodeMalformedPolicyDocumentException:
				fmt.Println(secretsmanager.ErrCodeMalformedPolicyDocumentException, aerr.Error())
			case secretsmanager.ErrCodeInternalServiceError:
				fmt.Println(secretsmanager.ErrCodeInternalServiceError, aerr.Error())
			case secretsmanager.ErrCodePreconditionNotMetException:
				fmt.Println(secretsmanager.ErrCodePreconditionNotMetException, aerr.Error())
			case secretsmanager.ErrCodeDecryptionFailure:
				fmt.Println(secretsmanager.ErrCodeDecryptionFailure, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}
	fmt.Println(result)
}

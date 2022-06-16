# LocalStack Automation

What is LocalStack?. LocalStack provides an easy-to-use test/mocking framework for developing Cloud applications. It spins up a testing environment on your local machine that provides the same functionality and APIs as the real AWS cloud environment.

Reference: https://localstack.cloud/

Here I am building an automation tool that does:

- Deploy localstack locally on your computer

```
# Deploy localstack using Docker:
make localstack-install

# Check localstack status deployment:
make localstack-status


# Simple validation:

export AWS_ACCESS_KEY_ID="test"
export AWS_SECRET_ACCESS_KEY="test"
export AWS_DEFAULT_REGION="us-east-1"
export AWS_PAGER=""

make localstack-test-awscli
```

- Create random resources mainly in AWS (Secrets in AWS Secret Manager)

## Requirements

```
Docker
python v3
golang >1.17
```


 ## Build the Golang Application

Build the Application with the following command
 
 ```
 make go-build
 ```

 ## AWS SECRET MANAGER

### Create random secrets

Especify the number of secrets you want to create with the `awssecretstotal` flag. Example:

 ```
 ./bin/localstack-automation -awssecretstotal=1000
 ```

 ### List all secrets

 List all existing secrets in localstack and the total:

 `./bin/localstack-automation --awssecretslist`
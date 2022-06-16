# LocalStack Automation
#TODO: explanation


## Requirements

```
Docker
python v3
golang >1.17
```

## Localstack installation and docker deploy
#TODO: explanation

```
make localstack-install
make localstack-status
```

## Export environment variables to connect AWS CLI locally to localstack
#TODO: explanation

```
export AWS_ACCESS_KEY_ID="test"
export AWS_SECRET_ACCESS_KEY="test"
export AWS_DEFAULT_REGION="us-east-1"
export AWS_PAGER=""
```

Then validate the AWS CLI connects to localstack with following command.

 `localstack-test-awscli`


## Create randoms AWS Secrets in localstack
#TODO: explanation

 `make aws-secrets.create --total 100`

 ## GOLANG
 ```
 make go-run
 ```
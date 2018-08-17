export GOOS=linux
AWS_REGION=us-east-2
serverlessl_FUNCTION=serverlesslGetCa
serverlessl_S3_BUCKET=serverlessl
serverlessl_S3_PREFIX=etcd/

deployment.zip:
	go build -o get_ca
	zip deployment.zip get_ca
	rm get_ca

clean:
	rm deployment.zip get_ca

.PHONY: deployment.zip
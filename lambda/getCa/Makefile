export GOOS=linux
AWS_REGION=us-east-2
serverlessl_FUNCTION=serverlesslGetCa
serverlessl_S3_BUCKET=serverlessl
serverlessl_S3_PREFIX=etcd/

deployment.zip: getCa.go
	go build -o getCa
	zip deployment.zip getCa
	rm getCa

clean:
	rm deployment.zip getCa

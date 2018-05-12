export GOOS=linux
AWS_REGION=us-east-2
serverlessl_FUNCTION=serverlesslInit
serverlessl_S3_BUCKET=serverlessl
serverlessl_S3_PREFIX=etcd/

deployment.zip: init.go
	go build -o init
	zip deployment.zip init
	rm init

clean:
	rm deployment.zip init

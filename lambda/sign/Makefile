export GOOS=linux
AWS_REGION=us-east-2
serverlessl_FUNCTION=serverlesslSign
serverlessl_S3_BUCKET=serverlessl
serverlessl_S3_PREFIX=etcd/

deployment.zip: sign.go
	go build -o sign
	zip deployment.zip sign
	rm sign

clean:
	rm deployment.zip sign

create: package
	aws lambda create-function \
		--region $(AWS_REGION) \
		--function-name $(serverlessl_FUNCTION) \
		--description "serverlessl signing function" \
		--environment "Variables={serverlessl_S3_BUCKET=$(serverlessl_S3_BUCKET),serverlessl_S3_PREFIX=$(serverlessl_S3_PREFIX)}" \
		--handler sign \
		--runtime go1.x \
		--role arn:aws:iam::057120273157:role/serverlessl_role \
		--zip-file fileb://./deployment.zip

upload: package
	aws lambda update-function-code \
		--region $(AWS_REGION) \
		--function-name $(serverlessl_FUNCTION) \
		--zip-file fileb://./deployment.zip
	$(MAKE) clean

all: package

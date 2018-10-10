build:
	cd lambda/getCa && make
	cd lambda/sign && make

test: 
	go test ./...

upload:
	aws s3 sync lambda/pkgs s3://serverlessl
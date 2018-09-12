build:
	cd lambda/getCa && make
	cd lambda/sign && make

test: 
	go test ./...
export GOOS=linux

build: 
	go build -o sign
	zip ../pkgs/sign.zip sign
	rm sign

.PHONY: build

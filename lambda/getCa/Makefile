export GOOS=linux

build:
	go build -o get_ca
	zip ../pkgs/ca.zip get_ca
	rm get_ca

.PHONY: build
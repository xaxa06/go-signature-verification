BINARY_NAME=signature

all: build

build: signaturefile main.go
	go build -o bin/${BINARY_NAME} main.go

run: build
	./bin/${BINARY_NAME}

.PHONY: clean
clean:
	go clean
	rm -rf resources/private.pem \
		resources/public.pem \
		resources/sha256.sign \
		bin

keyfiles := $(wildcard resources/*.pem)
signaturefile := resources/sha256.sign

keyfiles:
	openssl genrsa -out resources/private.pem 512
	openssl rsa -in resources/private.pem -pubout > resources/public.pem

signaturefile: keyfiles
	openssl dgst -sha256 -sign resources/private.pem -out resources/sha256.sign resources/message
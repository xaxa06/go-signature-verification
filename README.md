#  go-signature-verification
Quick lab for experimenting signature verification in Go.

Before the Go program is invoked, a key pair is generated and `resources/message` is signed using the private key (see Makefile for corresponding openssl commands).
The Go program output whether the previously generated signature verifies for `resources/message` content.

## Usage:
```sh
make all
./bin/signature
```
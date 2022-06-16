Quick lab for experimenting signature verification in Go.

Before the Go program is invoked, a key pair is generated and resources/message is signed using the private key (see Makefile for corresponding openssl commands).
The Go program verifies iffthe resources/message (more precisely its SHA 256 hash) matches the previously generated signature.

Usage:
```sh
make all
./bin/signature
```
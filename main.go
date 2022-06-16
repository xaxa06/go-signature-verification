package main

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

// https://pkg.go.dev/crypto/rsa#VerifyPKCS1v15
func main() {
	// For reference, the content of signature corresponds to:
	// hexdump -v -e '/1 "%02X"' resources/sha256.sign
	// signature, _ := hex.DecodeString("9F6CE56721ACCC394A3B487FC0C1EB60C2FEFB57802085C5D9A4CB7F6AC2167FB99FBAC3075A302213A1EA85D5E90AE1CEF988DBB036A1E201973FD86896608F")
	signature := FileToBytes("resources/sha256.sign")
	hash := FileToSHA256("resources/message")
	pubKey := FileToPublicKey("resources/public.pem")

	err := rsa.VerifyPKCS1v15(pubKey, crypto.SHA256, hash, signature)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error from verification: %s\n", err)
		return
	}
	fmt.Println("Verified OK!")
}

// FileToSignature returns the raw bytes of the file content
func FileToBytes(filePath string) []byte {
	signature, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(fmt.Sprintf("Cant open signature file: %v", err))
	}
	return signature
}

// FileToSHA256 computes and return the SHA 256 of file content at filePath
func FileToSHA256(filePath string) []byte {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		panic(err)
	}
	return h.Sum(nil)
}

// FileToPublicKey public key pem file to PublicKey object
func FileToPublicKey(pubKeyFilePath string) *rsa.PublicKey {
	var err error

	f, err := os.Open(pubKeyFilePath)
	if err != nil {
		panic(fmt.Sprintf("cant open pubkey file: %v", err))
	}

	data, err := ioutil.ReadAll(f)
	if err != nil {
		panic(fmt.Sprintf("couldnt read pubkey file: %v", err))
	}

	return BytesToPublicKey(data)
}

// BytesToPublicKey bytes to public key (https://gist.github.com/miguelmota/3ea9286bd1d3c2a985b67cac4ba2130a)
func BytesToPublicKey(pub []byte) *rsa.PublicKey {
	block, _ := pem.Decode(pub)
	enc := x509.IsEncryptedPEMBlock(block)
	b := block.Bytes
	var err error
	if enc {
		fmt.Println("is encrypted pem block")
		b, err = x509.DecryptPEMBlock(block, nil)
		if err != nil {
			panic(err)
		}
	}
	ifc, err := x509.ParsePKIXPublicKey(b)
	if err != nil {
		panic(err)
	}
	key, ok := ifc.(*rsa.PublicKey)
	if !ok {
		panic("the resulting private key read cant be cast to a rsa PublicKey type")
	}
	return key
}

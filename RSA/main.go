package main

import (
	"encoding/pem"
	"crypto/x509"
	"crypto/rsa"
	"crypto/rand"
	"fmt"
	"encoding/base64"
)

var prikey = []byte(`-----BEGIN RSA PRIVATE KEY-----
MIICXgIBAAKBgQCiiXBUieAiBEiFCvM/C+VQaj7HACwJDlZRYC9DO9iaBSZN3yfj
pxxGkBhIDILh8jR+rN/65juzUkyQh9yixfoIGmFI7OoHjvR4Nv5uu8iUPXespogN
ipGvSdaPlyVZdO+kLeMNCxidyTBm9HcUjRrHq1LJC4lZJEyqh7poYI1qKwIDAQAB
AoGATpvKCBjmnY0UxcOWAVfvgATid7xNj9OvS4wJT6sSBuuWqvLSh/e6ZSYrmCz0
3/IUEW/qo53WIdQb04rh2peFZMCFrTJrwWewj5l6EMclXcpzp7FxHYFYFkvt5EOk
+9zuaPg0j0w32Ry6McAELvxH1likdobUlKgIcsVdU+3NJoECQQDPsxSpGv4XNFtW
ArxaB7bgn/m0xe7mTUfXM5I9PfQ35wzzV8g0KX+p0+gMkMdcsyn8Hwx9hU/+cdxF
i7wBBXTXAkEAyFWznjIfqseMlGcAjgdiFLjL2BmdhfuWUzNV3Bc2s3PNUleEac7l
nPGjAdOSUnc7NYL1th2oxXlo9jIWYRK2zQJBAMWCtNkDFDIY1ep6+4ZZ46zENH+V
4lMbln+tSRn3+m/Wxlf6WCZSeFaVbwvT+eHLdteM5yHOTn3W5PM1qHkIKU0CQQDC
4aPWi3oEvxQBNWXzxDjkYZRixyBWuxRUIYcvkCF1Vnxw0pQL/emdz+47k54uvLAL
cTcnjxij6WRywuzER9ktAkEApwj2UDpbytOeEUmZd6Ryl8BiDeislLLl3rTcj2F1
6VhyHFg9QLr4oav3Nh91uqULUUfl/2IkdR2gM3CM6Czm8g==
-----END RSA PRIVATE KEY-----`)

var pubkey = []byte(`-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCiiXBUieAiBEiFCvM/C+VQaj7H
ACwJDlZRYC9DO9iaBSZN3yfjpxxGkBhIDILh8jR+rN/65juzUkyQh9yixfoIGmFI
7OoHjvR4Nv5uu8iUPXespogNipGvSdaPlyVZdO+kLeMNCxidyTBm9HcUjRrHq1LJ
C4lZJEyqh7poYI1qKwIDAQAB
-----END PUBLIC KEY-----`)

// 加密函数
func RSAEncrypt(origData []byte) []byte {

	// 公钥加密
	block, _ := pem.Decode(pubkey)
	// 解析公钥
	pubInterfacr, _ := x509.ParsePKIXPublicKey(block.Bytes)
	// 加载公钥
	pub := pubInterfacr.(*rsa.PublicKey)
	// 加密明文
	bits, _ := rsa.EncryptPKCS1v15(rand.Reader, pub, origData)
	// bits为加密的密文
	return bits
}

// 解密函数
func RSADecrypt(cipherTxt []byte) []byte {
	block , _ := pem.Decode(prikey)
	// 解析私钥
	priv, _ := x509.ParsePKCS1PrivateKey(block.Bytes)
	// 解密
	bts, _ := rsa.DecryptPKCS1v15(rand.Reader, priv, cipherTxt)
	return bts
}



func main()  {
	cipherTxt := RSAEncrypt([]byte("hello world"))
	fmt.Println(base64.StdEncoding.EncodeToString(cipherTxt))

	fmt.Println("解密的结果为：", string(RSADecrypt(cipherTxt)))
}

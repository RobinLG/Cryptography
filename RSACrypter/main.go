package main

import (
	"crypto/rsa"
	"crypto/rand"
	"crypto/md5"
	"fmt"
	"encoding/base64"
)

func main() {
	// 创建私钥
	prikey, _ := rsa.GenerateKey(rand.Reader, 1024)

	// 用私钥创建公钥
	pubkey := prikey.PublicKey

	org := []byte("hello")
	// 通过oaep函数实现公钥加密
	cipherTxt, _ := rsa.EncryptOAEP(md5.New(), rand.Reader, &pubkey, org, nil)

	fmt.Println(base64.StdEncoding.EncodeToString(cipherTxt))

	// 解密
	plainTxt, _ := rsa.DecryptOAEP(md5.New(), rand.Reader, prikey, cipherTxt, nil)
	fmt.Println(string(plainTxt))
}

package main

import (
	"crypto/rsa"
	"crypto/rand"
	"crypto/md5"
	"crypto"
	"fmt"
	"encoding/base64"
)

func main() {

	// 创建私钥
	prikey, _ := rsa.GenerateKey(rand.Reader, 1024)

	// 用私钥创建公钥
	pubkey := prikey.PublicKey

	// 准备做签名的Hash散列值
	plainTxt := []byte("hello")

	// 散列过程
	//hashed := md5.Sum(plainTxt)
	h := md5.New()
	h.Write(plainTxt)
	hashed := h.Sum(nil)

	// 用过Pss函数，实现对明文签名
	// pss函数添加杂质，使签名过程更安全
	opts := rsa.PSSOptions{rsa.PSSSaltLengthAuto, crypto.MD5}

	// 实现签名
	sign , _ := rsa.SignPSS(rand.Reader, prikey, crypto.MD5, hashed, &opts)

	fmt.Println(base64.StdEncoding.EncodeToString(sign))

	// 公钥验签
	err := rsa.VerifyPSS(&pubkey, crypto.MD5, hashed, sign, &opts)
	if err == nil {
		fmt.Println("验签成功")
	}
}

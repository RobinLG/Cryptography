package main

import (
	"crypto/dsa"
	"crypto/rand"
	"fmt"
)

func main() {

	// 通过dsa实现签名

	// 设置私钥使用的参数
	var param dsa.Parameters
	dsa.GenerateParameters(&param, rand.Reader, dsa.L1024N160)

	// 创建私钥
	var pri dsa.PrivateKey
	pri.Parameters = param

	// 生成私钥
	dsa.GenerateKey(&pri, rand.Reader)

	// 创建公钥
	pub := pri.PublicKey

	message := []byte("hello")

	// 签名
	r, s, _ := dsa.Sign(rand.Reader, &pri, message)

	//验证
	if dsa.Verify(&pub, message, r, s) {
		fmt.Println("验签成功")
	}
}

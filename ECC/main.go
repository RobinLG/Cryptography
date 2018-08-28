package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"math/big"
	"fmt"
)

func main() {

	message := []byte("hello")
	digest := sha256.Sum256(message)

	// 生成私钥
	// elliptic.P256()设置生成私钥为256位
	privatekey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	// 创建公钥
	publickey := privatekey.PublicKey

	// 私钥签名
	r, s, _ := ecdsa.Sign(rand.Reader, privatekey, digest[:])

	//======================以下代码是保存签名结果r,s=====================

	// 获取私钥的长度(字节)
	curveBytes := privatekey.Curve.Params().P.BitLen()/8

	// 获得签名返回的字节
	rByte, sByte := r.Bytes(), s.Bytes()

	// 创建数组，r与s字节长度为私钥长度的一致
	signature := make([]byte, curveBytes*2)
	copy(signature[:len(rByte)], rByte)
	copy(signature[len(rByte):], sByte)

	//=================================================================

	// tcp发送信息......

	// 验签
	digest = sha256.Sum256(message)
	// 获得公钥字节长度
	curveBytes = publickey.Curve.Params().P.BitLen()/8

	// 保存rbyte,sbyte
	r, s = new(big.Int), new(big.Int)

	r.SetBytes(signature[:curveBytes])
	s.SetBytes(signature[curveBytes:])

	//验签
	e := ecdsa.Verify(&publickey, digest[:], r, s)
	if e == true {
		fmt.Println("验签成功")
	}
}


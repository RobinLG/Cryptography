package main

import (
	"net"
	"fmt"
	"encoding/pem"
	"crypto/x509"
	"crypto/rsa"
	"crypto/md5"
	"crypto"
)

var pubkey = []byte(`-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCiiXBUieAiBEiFCvM/C+VQaj7H
ACwJDlZRYC9DO9iaBSZN3yfjpxxGkBhIDILh8jR+rN/65juzUkyQh9yixfoIGmFI
7OoHjvR4Nv5uu8iUPXespogNipGvSdaPlyVZdO+kLeMNCxidyTBm9HcUjRrHq1LJ
C4lZJEyqh7poYI1qKwIDAQAB
-----END PUBLIC KEY-----`)

func Recive() []byte {
	netListen, _ := net.Listen("tcp", "127.0.0.1:1234")

	for {
		conn, _ := netListen.Accept()

		// 保存接收到的数据
		data := make([]byte, 2048)
		for {
			n, _ := conn.Read(data)

			return data[:n]
		}
	}
}

func main() {

	data := Recive()
	//fmt.Println(string(data))

	// 拆分数组
	plaintxt := data[:5]
	fmt.Println("收到的数据为：", plaintxt)

	// 用公钥验签
	sig := data[5:]
	block, _ := pem.Decode(pubkey)
	// 解密公钥
	pubInterface, _ := x509.ParsePKIXPublicKey(block.Bytes)
	// 类型断言 类型为*publickey
	pub := pubInterface.(*rsa.PublicKey)

	h := md5.New()
	h.Write([]byte("Robin"))
	hashed := h.Sum(nil)

	e := rsa.VerifyPSS(pub, crypto.MD5, hashed, sig, nil)

	if e == nil {
		fmt.Println("验签成功，接收到的数据为：", string(plaintxt))
	}



}

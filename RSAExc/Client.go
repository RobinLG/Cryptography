package main

import (
	"net"
	"fmt"
	"crypto/md5"
	"encoding/pem"
	"crypto/x509"
	"crypto/rsa"
	"crypto"
	"crypto/rand"
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

// 签名
func Sign() []byte {

	// 散列明文
	plainTxt := []byte("Robin")
	//hashed := md5.Sum(plainTxt)
	h := md5.New()
	h.Write(plainTxt)
	hashed := h.Sum(nil)

	// 解析prikey，转换成*privtekey类型
	block, _ := pem.Decode(prikey)
	priv, _ := x509.ParsePKCS1PrivateKey(block.Bytes)

	// 签名
	opts := rsa.PSSOptions{rsa.PSSSaltLengthAuto, crypto.MD5}
	sign, _ := rsa.SignPSS(rand.Reader, priv, crypto.MD5, hashed, &opts)
	return sign
}

// 通过TCP发送数据
func Send(data []byte) {
	// 解析地址
	conn, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:1234")
	// 链接
	n, _ := net.DialTCP("tcp", nil, conn)
	// 发送数据
	n.Write(data)

	fmt.Println("发送结束")
}

func main() {
	// Send([]byte("hello"))

	// 发送的时候"hello"与sign拼接
	sign := Sign()
	data := make([]byte, len("hello") + len(sign))
	copy(data[:5], []byte("hello"))
	copy(data[5:], sign)

	fmt.Println("发送的总数据为：", data)
	fmt.Println("发送的内容为：", []byte("hello"))
	fmt.Println("发送的签名数据为：",sign)

	Send(data)
}


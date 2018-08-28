package main

import (
	"net"
	"fmt"
	"robin/MyCrypto/CryptedDic"
)

func Client(cipherTxt []byte) {
	// 解析address地址并返回
	netAddr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:1234")
	// 链接服务器
	conn, _ := net.DialTCP("tcp", nil, netAddr)
	conn.Write(cipherTxt)
	fmt.Println("发送数据")
}

func main() {
	data := []byte("It's a secret")
	cipherTxt := CryptedDic.EnCrypto("12345678", data)
	Client(cipherTxt)
}

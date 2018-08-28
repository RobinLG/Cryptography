package main

import (
	"net"
	"robin/MyCrypto/CryptedDic"
	"fmt"
)

func Server() []byte {
	//监听1234端口
	netListen, _ := net.Listen("tcp", "127.0.0.1:1234")
	defer netListen.Close()

	for {
		//等待链接
		conn, _ := netListen.Accept()

		//接收数据
		data := make([]byte, 2048)
		for {
			n, _ := conn.Read(data)
			// fmt.Println("接收到的数据为：", string(data[:n]))
			// break
			return data[:n]
		}
	}
}

func main() {
	cipherTxt := Server()
	orig := CryptedDic.DeCrypto("12345678", cipherTxt)
	fmt.Println("解密后的数据为：", string(orig))
}
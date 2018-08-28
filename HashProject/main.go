package main

import (
	"fmt"
	"crypto/md5"
	"encoding/hex"
	"crypto/sha256"
	"os"
	"io"
	"golang.org/x/crypto/ripemd160"
)

func MyMd5() {
	// 测试Md5的编写方法

	// Md5第一写法
	// 准备加密的明文
	data := []byte("hello world")
	// 用Md5加密
	s := fmt.Sprintf("%x", md5.Sum(data))
	// 十六进制，16字节，16*8=128位
	fmt.Println(s)

	// Md5第二写法
	data2 := []byte("hello world")
	m := md5.New()
	m.Write(data2)
	// 字节数组转换成字符串
	s2 := hex.EncodeToString(m.Sum(nil))
	fmt.Println(s2)
}

// 测试sha256加密算法的使用
func MySha256() {
	// 第一种
	data := []byte("hello world")
	s := fmt.Sprintf("%x", sha256.Sum256(data))
	fmt.Println(s)

	// 第二种
	data2 := []byte("hello world")
	m := sha256.New()
	m.Write(data2)
	fmt.Println(hex.EncodeToString(m.Sum(nil)))

	// 第三种
	// 将文件读入内存
	f, _ := os.Open("test")
	h := sha256.New()
	// 将数据拷入内存中
	io.Copy(h, f)
	s2 := h.Sum(nil)
	fmt.Println(hex.EncodeToString(s2))
}

func MyRipemd160() {
	hash := ripemd160.New()
	hash.Write([]byte("hello world"))
	fmt.Println(hex.EncodeToString(hash.Sum(nil)))
}

func main() {
	fmt.Println("---Md5---")
	MyMd5()
	fmt.Println("---Sha256---")
	MySha256()
	fmt.Println("---Ripemd160---")
	MyRipemd160()
}

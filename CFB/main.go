package main

import (
	"crypto/aes"
	"io"
	"crypto/rand"
	"crypto/cipher"
	"fmt"
	"encoding/hex"
	"encoding/base64"
)

// 加密
func AesCFBEncrypter(plainTxt []byte, key []byte) []byte {
	// 检查key合法
	block, _ := aes.NewCipher(key)
	cipherTxt := make([]byte, aes.BlockSize + len(plainTxt))
	iv := cipherTxt[:aes.BlockSize]

	// 向iv切片数组初始化为rand, Reader（随机内存流）
	io.ReadFull(rand.Reader, iv)

	// 设置加密模式为CFB
	stream := cipher.NewCFBEncrypter(block, iv)

	// 加密
	stream.XORKeyStream(cipherTxt[aes.BlockSize:], plainTxt)

	// cipherTxt 包含了key和明文两部分加密内容
	return cipherTxt
}

// 解密
func AesCFBDecrypter(cipherTxt []byte, key []byte) []byte {
	block, _ := aes.NewCipher(key)

	// 拆分iv和密文
	iv := cipherTxt[:aes.BlockSize]
	cipherTxt = cipherTxt[aes.BlockSize:]

	// 设置解密模式
	stream := cipher.NewCFBDecrypter(block, iv)

	des := make([]byte, len(cipherTxt))

	//解密
	stream.XORKeyStream(des, cipherTxt)

	return des
}

func main() {
	cipher := AesCFBEncrypter([]byte("hello"), []byte("1234567890123456"))
	fmt.Println(hex.EncodeToString(cipher))
	fmt.Println(base64.StdEncoding.EncodeToString(cipher))

	des := AesCFBDecrypter(cipher, []byte("1234567890123456"))
	fmt.Println(string(des))
}

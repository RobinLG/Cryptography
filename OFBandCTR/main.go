package main

import (
	"crypto/aes"
	"io"
	"crypto/rand"
	"crypto/cipher"
	"fmt"
	"encoding/hex"
)

// 加密
func AesOFBEncrypter(plainTxt []byte, key []byte) []byte {
	block, _ := aes.NewCipher(key)
	cipherTxt := make([]byte, block.BlockSize() + len(plainTxt))
	iv := cipherTxt[:block.BlockSize()]
	io.ReadFull(rand.Reader, iv)
	//设置加密模式
	//stream := cipher.NewOFB(block, iv)
	stream := cipher.NewCTR(block, iv)

	stream.XORKeyStream(cipherTxt[block.BlockSize():], plainTxt)
	return  cipherTxt
}

// 解密
func AesOFBDecrypter(cipherTxt []byte, key []byte) []byte {
	block, _ := aes.NewCipher(key)

	iv := cipherTxt[:block.BlockSize()]
	cipherTxt = cipherTxt[block.BlockSize():]

	//stream := cipher.NewOFB(block, iv)
	stream := cipher.NewCTR(block, iv)

	des := make([]byte, len(cipherTxt))
	stream.XORKeyStream(des, cipherTxt)
	return des

}

func main() {

	cipherTxt := AesOFBEncrypter([]byte("hello"), []byte("1234567890123456"))
	fmt.Println(hex.EncodeToString(cipherTxt))

	fmt.Println("解密后的结果: ", string(AesOFBDecrypter(cipherTxt, []byte("1234567890123456"))))

}

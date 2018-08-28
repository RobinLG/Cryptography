package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"fmt"
)

// PKCS5Padding 要求分组长度只能为8

// PKCS7Padding 要求分组的长度可以为[1-255]

// 补码
func PKCS7Padding(org []byte, blockSize int) []byte {
	padding := blockSize - len(org) % blockSize
	padAddr := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(org, padAddr ...)
}

// 去码
func PKCS7UnPadding(org []byte) []byte {
	// 如传"abc"，秘钥长度为8，则补码后abc55555，后面的数字大小与多少相等
	l := len(org)
	txt := int(org[l - 1])
	return org[:l - txt]
}

// 通过CBC分组模式，完成AES的密码过程
// AES也是对称加密
// AES秘钥长度，16/24/32
func AesCBCEnCrypto(org []byte, key []byte) []byte {
	// 校验秘钥
	block, _ := aes.NewCipher(key)
	// 按照公钥的长度进行分组补码
	org = PKCS7Padding(org, block.BlockSize())
	// 设置AES加密模式
	blockMode := cipher.NewCBCEncrypter(block, key)
	// 加密处理
	crypted := make([]byte, len(org))
	blockMode.CryptBlocks(crypted, org)

	return crypted

}

// 解密
func AesCBCDeCrypto(cipherTxt []byte, key []byte) []byte {
	block, _ := aes.NewCipher(key)
	blockMode := cipher.NewCBCDecrypter(block, key)
	org := make([]byte, len(cipherTxt))
	blockMode.CryptBlocks(org, cipherTxt)
	org = PKCS7UnPadding(org)
	return org

}

func main() {
	txt := AesCBCEnCrypto([]byte("hello"), []byte("1234567890123456"))
	fmt.Println("解密后的结果：", string(AesCBCDeCrypto(txt, []byte("1234567890123456"))))
}

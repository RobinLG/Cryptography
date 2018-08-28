package main

import (
	"bytes"
	"crypto/des"
	"crypto/cipher"
	"fmt"
	"encoding/hex"
)

func main()  {

	key := []byte("12345678")
	data := []byte("hello world")
	cipherTxt := DESEncrpto(data, key)
	fmt.Println("加密的结果：", hex.EncodeToString(cipherTxt))

	origData := DESDecrypt(cipherTxt, key)
	fmt.Println("解密的结果：",string(origData))
}

// 调用系统库中DES加密
func DESEncrpto(origData []byte, key []byte) []byte {
	// DES加密中Key长度必须为8
	// 3DES加密中Key长度必须为24

	// 校验秘钥
	block, _ := des.NewCipher(key)

	// 补码
	origData = PKC5Padding(origData, block.BlockSize())

	// 设置加密模式
	BlockMode := cipher.NewCBCEncrypter(block, key)

	// 加密明文
	crypted := make([]byte, len(origData))
	BlockMode.CryptBlocks(crypted, origData)
	return crypted
}

// 调用系统库中DES解密
func DESDecrypt(cryted []byte, key []byte) []byte {

	// 检验秘钥有效性
	block, _ := des.NewCipher(key)
	// 设置解密模式
	blockMode := cipher.NewCBCDecrypter(block, key)

	// 实现解密
	origData := make([]byte, len(cryted))
	blockMode.CryptBlocks(origData, cryted)

	origData = PKC5UnPadding(origData)
	return origData
}

// 去码
func PKC5UnPadding(cipherTxt []byte) []byte {
	l := len(cipherTxt)
	txt := int(cipherTxt[l -1])
	return cipherTxt[:l - txt]
}

// 补码
func PKC5Padding(cipherTxt []byte, blockSize int) []byte {

	padding := blockSize - len(cipherTxt) % blockSize

	padTxt := bytes.Repeat([]byte{byte(padding)}, padding)

	byteTxt := append(cipherTxt, padTxt ...)

	return byteTxt
}
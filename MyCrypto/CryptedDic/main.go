package CryptedDic

import "bytes"

/*
func main() {

	//fmt.Println(PKC5Padding([]byte("abc"), 8))
	//
	//rlt := PKC5Padding([]byte("abc"), 8)
	//fmt.Println(PKC5UnPadding(rlt))

	// 设置秘钥
	key := "12312312"
	// 设置明文
	// hello world55555
	data := []byte("hello world")

	cipherTxt := EnCrypto(key, data)
	fmt.Println("加密后的结果", hex.EncodeToString(cipherTxt))

	// 解密
	d := DeCrypto(key, cipherTxt)
	fmt.Println("解密后的结果：", string(d))
}*/

// 创建一个加密的方法
func EnCrypto(key string, data []byte) []byte {
	// 加密算法：首先计算key的总和，利用key的总和与明文参与运算

	sum := 0
	for i:=0; i<len(key); i++ {
		sum = sum + int(key[i])
	}

	// 首先对明文进行补码
	pad := PKC5Padding(data, len(key))

	// 通过加法运算，实现加密过程
	for i:=0; i<len(pad); i++ {
		pad[i] = pad[i] + byte(sum)
	}

	return pad

}

// 创建一个解密的方法
// 解密是加密的逆运算
func DeCrypto(key string, cipherTxt []byte) []byte {

	sum := 0
	for i:=0; i<len(key); i++ {
		sum = sum + int(key[i])
	}

	// 先做减法运算
	for i:=0; i<len(cipherTxt); i++ {
		cipherTxt[i] = cipherTxt[i] - byte(sum)
	}

	// 去码
	p := PKC5UnPadding(cipherTxt)

	return p
}

// 去码
func PKC5UnPadding(cipherTxt []byte) []byte {
	l := len(cipherTxt)
	// abc55555
	// a777777
	// 后面数字与数字所占位数相等
	txt := int(cipherTxt[l - 1])
	return cipherTxt[:l - txt]
}

// PKCS5Padding补码
func PKC5Padding(cipherTxt []byte, blockSize int) []byte {
	// 计算准备添加的数字
	padding := blockSize - len(cipherTxt) % blockSize

	// 55555
	padTxt := bytes.Repeat([]byte{byte(padding)},padding)

	byteTxt := append(cipherTxt, padTxt ...)

	return byteTxt
}

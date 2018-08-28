# Cryptography
## 代码许多细节处没有进行调整并且结构较为简单，意在分享在Go中使用各加密算法的方法，仅做参考或学习之用。需使用openssl工具创建生产公私钥，代码中也有直接使用算法自带的方法创建私钥。类Unix系统用openssl创建公私钥方法的命令行如下：   
#### 1. 生成私钥："genrsa -out rsa_private_key.pem 1024"
#### 2. 生成公钥："rsa -in rsa_private_key.pem -pubout -out rsa_public_key.pem"
#### 3. RSA私钥转换成 PKCS8 格式："pkcs8 -topk8 -inform PEM -in rsa_private_key.pem -outform PEM -nocrypt"(Java使用需要，PHP可略过此步骤)  

#### 所涉及的知识大致如下：  
####  链表 / Hash / MD5 / SHA256 / PKSC5Padding / DES / AES的ECB,CBC,CFB,OFB,CTR分组模式 / RSA / DSA / ECC

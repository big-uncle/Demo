package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"math/rand"
	"time"
)

const letterBytes = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const letterBytes1 = "0123456789abcdef"

func RandStringBytesRmndr1(n int) string {
	b := make([]byte, n)
	rand.Seed(time.Now().UnixNano())
	for i := range b {
		b[i] = letterBytes[rand.Int63()%int64(len(letterBytes1))]
	}
	return string(b)
}
func RandStringBytesRmndr(n int) string {
	b := make([]byte, n)
	rand.Seed(time.Now().UnixNano())
	for i := range b {
		b[i] = letterBytes[rand.Int63()%int64(len(letterBytes))]
	}
	return string(b)
}

func main() { ///生成随机32位秘钥
	str := RandStringBytesRmndr1(32)
	fmt.Println(str)
	hexkey, _ := hex.DecodeString(str) //10进制转为16进制
	fmt.Printf("%.15x\n", hexkey)      //  十六进制    将16进制转换为10进制来表示
	fmt.Printf("%x\n", hexkey)
	basekey := EncodeBase64(hexkey)
	fmt.Println(string(basekey))
	//随机生成秘钥
	secretKey := RandStringBytesRmndr(32)
	fmt.Printf("秘钥为%s\n", secretKey)
	fmt.Printf("秘钥为%.30s\n", secretKey)
	//字符串%20.30表示 字符串长度若小于20则，给前面补充20-len()个空格，使字符串整体长度为20个，若len()大于20则不变
	//小数点后面的则为，长度若大于30，则截取后面len()-30个字符串，若len()小于30则不变
	//%x   %s都一样都是一个道理
}

func EncodeBase64(origdata []byte) []byte {
	str := base64.StdEncoding.EncodeToString(origdata)
	return []byte(str)
}

func DecodeBase64(cipherdata []byte) []byte {
	orig, err := base64.StdEncoding.DecodeString(string(cipherdata))
	if err != nil {
		return nil
	}
	return orig
}

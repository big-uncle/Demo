package main

import (
	"./utils"
	"encoding/base64"
	"fmt"
)

func main() {
	origData := "我想静静" // 待加密的数据

	cipher, err := AesEncrypt(origData)
	if err != nil {
		fmt.Println("加密失败", err)
		return
	}
	fmt.Println("加密结果", cipher)

	encodebase64 := base64.StdEncoding.EncodeToString([]byte(cipher)) //再经base64编码
	fmt.Println("密文(base64)：", encodebase64)
	data, err := base64.StdEncoding.DecodeString(encodebase64) //再经base64解码
	if err != nil {
		panic(err)
	}

	fmt.Println("解密文(base64)：", string(data))
	clear, err := AesDecrypt(cipher)
	if err != nil {
		fmt.Println("解密失败", err)
		return
	}
	fmt.Println("解密结果", clear)

}

const secretKey = "ed4f8731b6ae7a19008fe896514a370b" // 加密的密钥随机生成32位秘钥

//aes加密
func AesEncrypt(origData string) (string, error) {
	if len(origData) == 0 {
		return "", fmt.Errorf("数据为空")
	}
	cipher := utils.AesEncryptCBC([]byte(origData), []byte(secretKey))
	if len(cipher) == 0 {
		return "", fmt.Errorf("数据解密失败")
	}
	//return string(security.EncodeBase64([]byte(hex.EncodeToString(cipher)))), nil //可以进行hex16进制转换，也可以不使用
	return string(utils.EncodeBase64(cipher)), nil
}

//aes解密    //如何使用base64，用于传输还是存储或是变量传输自己决定，但是怎样加密，就得怎样解密，俩者顺序必须一一对应
func AesDecrypt(encrypted string) (string, error) { //使不使用取决于加密和解密俩者的对应关系
	//plain, err := hex.DecodeString(string(security.DecodeBase64([]byte(encrypted))))//可以进行hex16进制转换，也可以不使用
	//if err != nil {
	//	return "", err
	//}
	if len(encrypted) == 0 {
		return "", fmt.Errorf("数据为空")
	}
	plain := utils.DecodeBase64([]byte(encrypted))
	data := utils.AesDecryptCBC(plain, []byte(secretKey))
	if len(data) == 0 {
		return "", fmt.Errorf("数据解密失败")
	}
	return string(data), nil
}

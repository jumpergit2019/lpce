package encrypt

import (
	"bytes"
)

func pKCS5Padding(ciphertext []byte, blockSize int) []byte {
	//注意此处是一定会进行填充的，填充个数为 1~blockSize 字节
	//因此在解密的时候才能从最后一个字节知道填充了几个字节，需要删除几个字节来获得原文
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)

}

func pKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	// 去掉最后一个字节 unpadding 次
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]

}

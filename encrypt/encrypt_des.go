package encrypt

import (
	"crypto/cipher"
	"crypto/des"
	"crypto/rand"
	"errors"
	"fmt"
	"io"
	"components/log"
)

type EncryptOpDes struct {
	desKey []byte
	direct bool
}


func (self *EncryptOpDes) Init(direct bool, params []interface{}) bool {

	if params == nil || len(params) != 1 {
		fmt.Printf("invalid param count.")
		return false
	}

	var ok bool
	self.desKey, ok = params[0].([]byte)
	if !ok {
		fmt.Printf("invalid param type ")
		return false
	}


	if len(self.desKey) != 64/8 {
		fmt.Println("err desKey length, need 128/192/256 bit.")
		return false
	}

	self.direct = direct
	return true
}


func (self *EncryptOpDes) Operate(input interface{}, output interface{}) (bool, error){

	if self.direct{
		tmpOutput, err := self.Encrypt(input.([]byte))
		if err != nil{
			fmt.Printf("pack failed. err: %s", err)
			return false, err
		}
		*(output.(*[]byte)) = tmpOutput
		return true, nil

	}else{
		tmpOutput, err := self.Decrypt(input.([]byte))
		if err != nil{
			fmt.Printf("unpack failed. err: %s", err)
			return false, err
		}
		*(output.(*[]byte)) = tmpOutput
		return true, nil
	}

	return true, nil
}

func (self *EncryptOpDes) Encrypt(data []byte) ([]byte, error) {

	defer log.TraceLog("EncryptOpDes.Encrypt")()
	if data == nil || self.desKey == nil {
		return nil, errors.New("invalid self.desKey or data")
	}

	block, err := des.NewCipher(self.desKey)
	if err != nil {
		panic(err)
	}

	//填充字节
	data = pKCS5Padding(data, block.BlockSize())

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	ciphertext := make([]byte, des.BlockSize+len(data))
	iv := ciphertext[:des.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[des.BlockSize:], data)

	// It's important to remember that ciphertexts must be authenticated
	// (i.e. by using crypto/hmac) as well as being encrypted in order to
	// be secure.

	//fmt.Printf("%x\n", ciphertext)
	return ciphertext, nil
}

func (self *EncryptOpDes) Decrypt(data []byte) ([]byte, error) {

	defer log.TraceLog("EncryptOpDes.Decrypt")()
	block, err := des.NewCipher(self.desKey)
	if err != nil {
		panic(err)
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	if len(data) < des.BlockSize {
		panic("ciphertext too short")

	}
	iv := data[:des.BlockSize]
	data = data[des.BlockSize:]

	// CBC mode always works in whole blocks.
	if len(data)%des.BlockSize != 0 {
		panic("ciphertext is not a multiple of the block size")

	}

	mode := cipher.NewCBCDecrypter(block, iv)

	// CryptBlocks can work in-place if the two arguments are the same.
	mode.CryptBlocks(data, data)

	data = pKCS5UnPadding(data)

	//fmt.Printf("%s\n", data)
	return data, nil
}

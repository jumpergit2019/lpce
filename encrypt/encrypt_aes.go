package encrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"fmt"
	"io"
	"components/log"
)

type EncryptOpAes struct {
	aesKey []byte
	direct bool
}



func (self *EncryptOpAes) Init(direct bool, params []interface{}) bool {

	if params == nil || len(params) != 1 {
		fmt.Printf("invalid param count.")
		return false
	}

	var ok bool
	self.aesKey, ok = params[0].([]byte)
	if !ok {
		fmt.Printf("invalid param type ")
		return false
	}

	if len(self.aesKey) != 128/8 && len(self.aesKey) != 192/8 && len(self.aesKey) != 256/8 {
		fmt.Println("err aesKey length, need 128/192/256 bit.")
		return false
	}

	self.direct = direct
	return true
}


func (self *EncryptOpAes) Operate(input interface{}, output interface{}) (bool, error){

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

func (self *EncryptOpAes) Encrypt(data []byte) ([]byte, error) {

	defer log.TraceLog("EncryptOpAes.Encrypt")()
	if data == nil || self.aesKey == nil {
		return nil, errors.New("invalid self.aesKey or data")
	}

	block, err := aes.NewCipher(self.aesKey)
	if err != nil {
		panic(err)
	}

	//填充字节
	data = pKCS5Padding(data, block.BlockSize())

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	ciphertext := make([]byte, aes.BlockSize+len(data))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[aes.BlockSize:], data)

	// It's important to remember that ciphertexts must be authenticated
	// (i.e. by using crypto/hmac) as well as being encrypted in order to
	// be secure.

	//fmt.Printf("%x\n", ciphertext)
	return ciphertext, nil
}

func (self *EncryptOpAes) Decrypt(data []byte) ([]byte, error) {

	defer log.TraceLog("EncryptOpAes.Decrypt")()
	block, err := aes.NewCipher(self.aesKey)
	if err != nil {
		panic(err)
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	if len(data) < aes.BlockSize {
		panic("ciphertext too short")

	}
	iv := data[:aes.BlockSize]
	data = data[aes.BlockSize:]

	// CBC mode always works in whole blocks.
	if len(data)%aes.BlockSize != 0 {
		panic("ciphertext is not a multiple of the block size")

	}

	mode := cipher.NewCBCDecrypter(block, iv)

	// CryptBlocks can work in-place if the two arguments are the same.
	mode.CryptBlocks(data, data)

	data = pKCS5UnPadding(data)

	// If the original plaintext lengths are not a multiple of the block
	// size, padding would have to be added when encrypting, which would be
	// removed at self point. For an example, see
	// https://tools.ietf.org/html/rfc5246#section-6.2.3.2. However, it's
	// critical to note that ciphertexts must be authenticated (i.e. by
	// using crypto/hmac) before being decrypted in order to avoid creating
	// a padding oracle.

	//fmt.Printf("%s\n", data)
	return data, nil
}

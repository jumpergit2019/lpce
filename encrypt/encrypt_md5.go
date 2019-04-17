package encrypt

import (
	"crypto/md5"
	"errors"
	"fmt"
	"components/log"
)

type EncryptOpMd5 struct {
	direct bool
}

func (self *EncryptOpMd5) Init(direct bool, params []interface{}) bool {
	self.direct = direct
	return true
}

func (self *EncryptOpMd5) Operate(input interface{}, output interface{}) (bool, error){

	if self.direct{
		tmpOutput, err := self.Encrypt(input.([]byte))
		if err != nil{
			fmt.Printf("pack failed. err: %s", err)
			return false, err
		}
		*(output.(*[]byte)) = tmpOutput
		return true, nil

	}else{
		fmt.Println("md5 couldn't decrypt.")
		return false, errors.New("md5 couldn't decrypt.")
	}

	return true, nil
}

func (*EncryptOpMd5) Encrypt(data []byte) ([]byte, error) {
	defer log.TraceLog("EncryptOpMd5.Encrypt")()
	r := md5.Sum(data)
	rst := r[:]
	return rst, nil

}

func (*EncryptOpMd5) Decrypt(data []byte) ([]byte, error) {
	defer log.TraceLog("EncryptOpMd5.Decrypt")()
	return nil, errors.New("md5 couldn't decrypt.")
}

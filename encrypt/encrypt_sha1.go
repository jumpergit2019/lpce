package encrypt

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"github.com/jumpergit2019/lpce/log"
)

type EncryptOpSha1 struct {
	direct bool
}

func (self *EncryptOpSha1) Init(direct bool, params []interface{}) bool{
	self.direct = direct
	return true
}

func (self *EncryptOpSha1) Operate(input interface{}, output interface{}) (bool, error){

	if self.direct{
		tmpOutput, err := self.Encrypt(input.([]byte))
		if err != nil{
			fmt.Printf("pack failed. err: %s", err)
			return false, err
		}
		*(output.(*[]byte)) = tmpOutput
		return true, nil

	}else{
		fmt.Println("sha1 couldn't decrypt.")
		return false, errors.New("sha1 couldn't decrypt.")
	}

	return true, nil
}

func (*EncryptOpSha1) Encrypt(data []byte) ([]byte, error) {
	defer log.TraceLog("EncryptOpSha1.Encrypt")()
	r := sha1.Sum(data)
	rst := r[:]
	return rst, nil

}

func (*EncryptOpSha1) Decrypt(data []byte) ([]byte, error) {
	defer log.TraceLog("EncryptOpSha1.Decrypt")()
	return nil, errors.New("sha1 couldn't decrypt.")
}

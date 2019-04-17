package package_op

import (
	"fmt"
	"../compress"
	"../encrypt"
	"../packet"
)

// 使用一系列　PackageOp 来构造一条操作链，　
//　对于封包过程来说，　输入为　interface{}, 输出为　[]byte
//　对于解包过程来说，　输入为 []byte, 输出为 interface{}


//!!!note:　需要明确一点,　一个单独的PackageOpLink是支持并发调用的，即可以在多个协程中互不干扰的使用
// 这是因为　所有涉及到的数据只有两种情况
// 1. 入参传入，因此不同协程不会具有共用数据　2. 初始化的成员如direct或者密钥，他们都是初始化之后不再变化，即只进行读取，不进行写入，因此可以并发访问

//因此在实际项目使用中可以考虑定义一个全局的 map[string]PackageOpLink　不同的需求定义不同的　PackageOpLink　放入其中．

type PackageOpLink struct{
	opLink []PackageOp
}

func (self *PackageOpLink) checkParam(opType PackageOpType, direct bool, params []interface{}) bool{
	if opType <= PackageOpMin || opType >=PackageOpMax{
		fmt.Println("invalid opType.")
		return false
	}

	if (opType == EncryptMd5 || opType == EncryptSha1) && !direct {
		fmt.Println("md5 | sha1 couldn't direct false.")
		return false
	}

	if (opType == EncryptAes || opType == EncryptDes) && len(params) != 1{
		fmt.Println("aes | des should set key.")
		return false
	}

	if (opType == EncryptRsa) && len(params) != 2{
		fmt.Println(" rsa should set key.")
		return false
	}

	return true
}

func (self *PackageOpLink) AddOp(opType PackageOpType, direct bool, params []interface{}) bool{
	if !self.checkParam(opType, direct, params){
		return false
	}

	var op PackageOp

	switch opType{
	//编码
	case PacketBase64:
		op = new(packet.PacketOpBase64)
		op.Init(direct, nil)
		break

	case PacketJson:
		op = new(packet.PacketOpJson)
		op.Init(direct, nil)
		break

	case PacketXml:
		op = new(packet.PacketOpXml)
		op.Init(direct, nil)
		break

	case PacketProtobuf:
		op = new(packet.PacketOpProtobuf)
		op.Init(direct, nil)
		break

		//压缩
	case CompressGzip:
		op = new(compress.CompressOpGzip)
		op.Init(direct, nil)
		break

	case CompressZlib:
		op = new(compress.CompressOpZlib)
		op.Init(direct, nil)
		break

		//加密
	case EncryptMd5:
		op = new(encrypt.EncryptOpMd5)
		op.Init(direct, nil)
		break

	case EncryptSha1:
		op = new(encrypt.EncryptOpSha1)
		op.Init(direct, nil)
		break

	case EncryptAes:
		op = new(encrypt.EncryptOpAes)
		op.Init(direct, params)
		break

	case EncryptDes:
		op = new(encrypt.EncryptOpDes)
		op.Init(direct, params)
		break

	case EncryptRsa:
		op = new(encrypt.EncryptOpRsa)
		op.Init(direct, params)
		break

	default:
		return false
		break
	}

	self.opLink = append(self.opLink, op)
	return true

}

func (self *PackageOpLink) Execute(input interface{}, output interface{}) error{

	//这里是一个链式反应，因此需要根据op类型来构建中间类型
	//中间过程的输出都是 []byte
	var tmpOutput []byte
	var tmpInput interface{}
	tmpInput = input

	for k := range self.opLink{
		if k == len(self.opLink)-1{
			rst,err := self.opLink[k].Operate(tmpInput, output)
			if !rst{
				fmt.Printf("op error: %s", err)
				return err
			}
			fmt.Printf("\ntmpInput: %v,\noutput: %v\n\n", tmpInput, output)

		}else{
			self.opLink[k].Operate(tmpInput, &tmpOutput)
			fmt.Printf("\ntmpInput: %v,\ntmpOutput: %v\n\n", tmpInput, tmpOutput)
			tmpInput = tmpOutput
		}
	}

	return nil
}


func (self *PackageOpLink) Reset() {
	if len(self.opLink)  != 0{
		self.opLink = self.opLink[:0]
	}
	return
}

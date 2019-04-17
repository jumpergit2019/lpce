package package_op


//package_op 表示一个操作，　包含三种类型：
//1. encode
//2. compress
//3. encrypt


//定义打包和加密类型
type PackageOpType int

const (
	PackageOpMin PackageOpType = 0 + iota
	//封包
	PacketBase64
	PacketJson
	PacketXml
	PacketProtobuf

	//压缩
	CompressGzip
	CompressZlib

	//加密
	EncryptMd5
	EncryptSha1
	EncryptAes
	EncryptDes
	EncryptRsa

	PackageOpMax
)



//操作接口
type PackageOp interface{
	Init(direct bool, params []interface{}) bool //direct 表示操作方向, true　表示　编码/压缩/加密，　false 表示　解码/解压/解密
	Operate(input interface{}, output interface{}) (bool,error)

}

type PacketOp interface{
	PackageOp
	Pack(originData interface{}) ([]byte, error)
	Unpack(packData []byte, obj interface{}) error
}

type CompressOp interface{
	PackageOp
	Compress(data []byte) ([]byte, error)
	Decompress(data []byte) ([]byte, error)
}

type EncryptOp interface{
	PackageOp
	Encrypt(data []byte) ([]byte, error)
	Decrypt(data []byte) ([]byte, error)
}



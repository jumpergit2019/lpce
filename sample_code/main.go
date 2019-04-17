package main

import (
	//"github.com/golang/protobuf/proto"
	//"components/package/encrypt"
	//"fmt"
	//"components/package/packet"
	"fmt"
	"components/package_v_1_0/package_op"
)

func main() {
	////////////////base64////////////////

	//var op packet.PacketOpBase64
	//op.Init(true, nil)
	//originData := []byte("this is a test.")
	//var rstData []byte
	//op.Operate(originData, &rstData)
	//fmt.Printf("rst: %s", rstData)
	//
	//var opr packet.PacketOpBase64
	//opr.Init(false, nil)
	//
	//var rstrData []byte
	//opr.Operate(rstData, &rstrData)
	//fmt.Printf("rstr: %s", rstrData)

	////////////////json////////////////

	//var op packet.PacketOpJson
	//op.Init(true, nil)
	//test1 := struct{
	//	Name string
	//	Age int
	//}{
	//	Name: "wang",
	//	Age: 1,
	//}
	//
	//var rst1 []byte
	//op.Operate(test1, &rst1)
	//fmt.Printf("rst1: %s", rst1)
	//
	//var test2 struct{
	//	Name string
	//	Age int
	//}
	//
	//var op2 packet.PacketOpJson
	//op2.Init(false, nil)
	//op2.Operate(rst1, &test2)
	//fmt.Printf("rst2: %v", test2)

	////////////////protobuf////////////////

	//bodyData := "guangzhou/fangcun/vip/company"
	//
	//p := &zqpacket.StringMessage{
	//	Body: proto.String(bodyData),
	//	Header: &zqpacket.Header{
	//		MessageId: proto.String("20-05"),
	//		Topic:     proto.String("golang"),
	//	},
	//}
	//
	//var op packet.PacketOpProtobuf
	//op.Init(true, nil)
	//
	//var rst1 []byte
	//op.Operate(p, &rst1)
	//fmt.Printf("rst1: %v", rst1)
	//
	//
	//var test2 zqpacket.StringMessage
	//var op2 packet.PacketOpProtobuf
	//op2.Init(false, nil)
	//op2.Operate(rst1, &test2)
	////注意此处　打印　&test2 和　test2　是不一样的，因为 *(zqpacket.StringMessage) 是定义了 String() 函数的
	//fmt.Println("rst2: ", &test2)

	////////////////xml////////////////

	//type info struct {
	//	Name string
	//	Age  int
	//	Male bool
	//}
	//
	////注意：xml　比较受限, 不能使用map 并且如下匿名结构体类型也是不能支持编码的
	////test1 := struct{
	////	Name string
	////	Age int
	////}{
	////	Name: "wang",
	////	Age: 1,
	////}
	//
	//var op packet.PacketOpXml
	//op.Init(true, nil)
	//var test1 info
	//test1.Name = "wang"
	//test1.Age = 1
	//test1.Male = true
	//
	//var rst1 []byte
	//op.Operate(&test1, &rst1)
	//fmt.Printf("rst1: %s", rst1)
	//
	//
	//var test2 info
	//var op2 packet.PacketOpXml
	//op2.Init(false, nil)
	//op2.Operate(rst1, &test2)
	//fmt.Printf("rst2: %v", test2)

	////////////////aes////////////////

	//var op encrypt.EncryptOpAes
	//params := make([]interface{}, 0)
	//params = append(params, []byte("abcdefghijklmnop"))
	//op.Init(true, params)
	//
	//param1 := []byte("this is a aes test")
	//var rst1 []byte
	//op.Operate(param1, &rst1)
	//fmt.Printf("rst1: %x\n", rst1)
	//
	//var test2 []byte
	//var op2 encrypt.EncryptOpAes
	//op2.Init(false, params)
	//op2.Operate(rst1, &test2)
	//fmt.Printf("rst2: %s", test2)

	////////////////des////////////////

	//var op encrypt.EncryptOpDes
	//params := make([]interface{}, 0)
	//params = append(params, []byte("ijklmnop"))
	//op.Init(true, params)
	//
	//param1 := []byte("this is a des test")
	//var rst1 []byte
	//op.Operate(param1, &rst1)
	//fmt.Printf("rst1: %x\n", rst1)
	//
	//var test2 []byte
	//var op2 encrypt.EncryptOpDes
	//op2.Init(false, params)
	//op2.Operate(rst1, &test2)
	//fmt.Printf("rst2: %s", test2)

	////////////////rsa////////////////


	// !!!此处注意，不要在密钥公钥中的某行前面留空格，并且应该使用``，不是""

//	PrivKeyLocal := []byte(`
//-----BEGIN RSA PRIVATE KEY-----
//MIICWwIBAAKBgQCnCnuWcNacRnqwDfSNLx7bbJLJM+foyxqSzp/M0fYqjhMp8voe
//51PUEGetCvM2kAakmRue6MXQ3TKrV7L6d3XTYGabBPzwDd0KoucklVVOS2vi1E7U
//V1bZhB60YdayCb9dcnEdA0uyA+qQgk2VhMtP1fER8lll5EiUUT+T0vnq9wIDAQAB
//AoGAfZo9Seb5CLNaR42GyK6Y1kdyrEYSaJJoHeGueTWbk24XbOCeQKSS/Q+E1bI5
//JVrxE81o3nmLXT0mf35HP1yaRCrofCV7a4QBlD9CNkMfy68fJEA6gMFuVVAES6Fa
//Zt1ENZ81NeENURUC+lLFSlUWm2Xbf+MZtCFIRE5Tj1HxvQkCQQDPTDZKpyqZ/1yg
//PO1/Quu0iisDYROJMm4sHQowIYXkHA/pUQMEveomBGRLavWrN9t4oEotFAPi0qYW
//847m7TmDAkEAzkkNyoz08+Dg4+SfwbjEyglyX7OkmOOGnCvEJldQm0wLZvrpJS6i
//n24UiYx2Cg93BZrvD9Ce7oNEnwbnHG3yfQJAJtOce6ER3qQwwiaHSUXMhhU29zwQ
//f6r9ba/Gv7sXq+EBre6phRLZL2O1MVcISph8t/w1yHmuPKa9yyC1TFV0ZwJADbeh
//6SQybb04dy8OyI0G2QCD0IVbnqcSnnPymTIZNBp8b56jvks5mSxyxSrH9qdMnNzO
//pNiUmPu1pnWJDMTq6QJAHIToUuuAN2z3pLpUJsM40T6sEwgbxiFPZ3iT4/T2Tgpy
//BKLqQxR7jXKdl0iWYteC96pQ0bqytFse4lnmPMUCew==
//-----END RSA PRIVATE KEY-----
//`)
//	//fmt.Printf("%s", PrivKeyLocal)
//
//	PubKeyRemote := []byte(`
//-----BEGIN PUBLIC KEY-----
//MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCnCnuWcNacRnqwDfSNLx7bbJLJ
//M+foyxqSzp/M0fYqjhMp8voe51PUEGetCvM2kAakmRue6MXQ3TKrV7L6d3XTYGab
//BPzwDd0KoucklVVOS2vi1E7UV1bZhB60YdayCb9dcnEdA0uyA+qQgk2VhMtP1fER
//8lll5EiUUT+T0vnq9wIDAQAB
//-----END PUBLIC KEY-----
//`)
//
//		var op encrypt.EncryptOpRsa
//		params := make([]interface{}, 0)
//		params = append(params, PubKeyRemote)
//		params = append(params, PrivKeyLocal)
//		op.Init(true, params)
//
//		param1 := []byte("this is a rsa test")
//		var rst1 []byte
//		op.Operate(param1, &rst1)
//		fmt.Printf("rst1: %x\n", rst1)
//
//		var test2 []byte
//		var op2 encrypt.EncryptOpRsa
//		op2.Init(false, params)
//		op2.Operate(rst1, &test2)
//		fmt.Printf("rst2: %s", test2)

	////////////////md5////////////////

	//var op encrypt.EncryptOpMd5
	//op.Init(true, nil)
	//
	//param1 := []byte("this is a md5 test")
	//var rst1 []byte
	//op.Operate(param1, &rst1)
	//fmt.Printf("rst1: %x\n", rst1)



	////////////////sha1////////////////

	//var op encrypt.EncryptOpSha1
	//op.Init(true, nil)
	//
	//param1 := []byte("this is a sha1 test")
	//var rst1 []byte
	//op.Operate(param1, &rst1)
	//fmt.Printf("rst1: %x\n", rst1)



	////////////////gzip////////////////

	//var op compress.CompressOpGzip
	//op.Init(true, nil)
	//
	//param1 := []byte("this is a gzip test")
	//var rst1 []byte
	//op.Operate(param1, &rst1)
	//fmt.Printf("rst1: %x\n", rst1)
	//
	//param2 := make([]byte, len(rst1))
	//copy(param2, rst1)
	//
	////fmt.Printf("param2: %s", param2)
	//var op2 compress.CompressOpGzip
	//op2.Init(false, nil)
	//
	//var test2 []byte
	//op2.Operate(param2, &test2)
	//fmt.Printf("rst2: %s\n", test2)

	////////////////zlib////////////////

	//var op compress.CompressOpZlib
	//op.Init(true, nil)
	//
	//param1 := []byte("this is a zlib test")
	//var rst1 []byte
	//op.Operate(param1, &rst1)
	//fmt.Printf("rst1: %x\n", rst1)
	//
	//param2 := make([]byte, len(rst1))
	//copy(param2, rst1)
	//
	//var op2 compress.CompressOpZlib
	//op2.Init(false, nil)
	//
	//var test2 []byte
	//op2.Operate(param2, &test2)
	//fmt.Printf("rst2: %s\n", test2)


	///////////////////////////////////////// 链接测试 /////////////////////////////////////////
	/////////////////////////// json/xml ///////////////////////////
	PrivKeyLocal := []byte(`
-----BEGIN RSA PRIVATE KEY-----
MIICWwIBAAKBgQCnCnuWcNacRnqwDfSNLx7bbJLJM+foyxqSzp/M0fYqjhMp8voe
51PUEGetCvM2kAakmRue6MXQ3TKrV7L6d3XTYGabBPzwDd0KoucklVVOS2vi1E7U
V1bZhB60YdayCb9dcnEdA0uyA+qQgk2VhMtP1fER8lll5EiUUT+T0vnq9wIDAQAB
AoGAfZo9Seb5CLNaR42GyK6Y1kdyrEYSaJJoHeGueTWbk24XbOCeQKSS/Q+E1bI5
JVrxE81o3nmLXT0mf35HP1yaRCrofCV7a4QBlD9CNkMfy68fJEA6gMFuVVAES6Fa
Zt1ENZ81NeENURUC+lLFSlUWm2Xbf+MZtCFIRE5Tj1HxvQkCQQDPTDZKpyqZ/1yg
PO1/Quu0iisDYROJMm4sHQowIYXkHA/pUQMEveomBGRLavWrN9t4oEotFAPi0qYW
847m7TmDAkEAzkkNyoz08+Dg4+SfwbjEyglyX7OkmOOGnCvEJldQm0wLZvrpJS6i
n24UiYx2Cg93BZrvD9Ce7oNEnwbnHG3yfQJAJtOce6ER3qQwwiaHSUXMhhU29zwQ
f6r9ba/Gv7sXq+EBre6phRLZL2O1MVcISph8t/w1yHmuPKa9yyC1TFV0ZwJADbeh
6SQybb04dy8OyI0G2QCD0IVbnqcSnnPymTIZNBp8b56jvks5mSxyxSrH9qdMnNzO
pNiUmPu1pnWJDMTq6QJAHIToUuuAN2z3pLpUJsM40T6sEwgbxiFPZ3iT4/T2Tgpy
BKLqQxR7jXKdl0iWYteC96pQ0bqytFse4lnmPMUCew==
-----END RSA PRIVATE KEY-----
`)
	//fmt.Printf("%s", PrivKeyLocal)

	PubKeyRemote := []byte(`
-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCnCnuWcNacRnqwDfSNLx7bbJLJ
M+foyxqSzp/M0fYqjhMp8voe51PUEGetCvM2kAakmRue6MXQ3TKrV7L6d3XTYGab
BPzwDd0KoucklVVOS2vi1E7UV1bZhB60YdayCb9dcnEdA0uyA+qQgk2VhMtP1fER
8lll5EiUUT+T0vnq9wIDAQAB
-----END PUBLIC KEY-----
`)

	type s1 struct{
		Name string
		Age int
		Male bool
	}

	test1 := s1{
		Name: "wang",
		Age: 1,
		Male: true,
	}


	var polink package_op.PackageOpLink
	polink.AddOp(package_op.PacketJson, true, nil)
	//polink.AddOp(package_op.PacketXml, true, nil)

	polink.AddOp(package_op.CompressGzip, true, nil)
	//polink.AddOp(package_op.CompressZlib, true, nil)

	//polink.AddOp(package_op.EncryptAes, true, []interface{}{[]byte("abcdefghijklmnop")})
	//polink.AddOp(package_op.EncryptDes, true, []interface{}{[]byte("ijklmnop")})
	polink.AddOp(package_op.EncryptRsa, true, []interface{}{PubKeyRemote, PrivKeyLocal})
	var rst1 []byte
	err := polink.Execute(test1, &rst1)
	if err != nil{
		fmt.Printf("err: %s", err)
		return
	}



	polink.Reset()
	//polink.AddOp(package_op.EncryptAes, false, []interface{}{[]byte("abcdefghijklmnop")})
	//polink.AddOp(package_op.EncryptDes, false, []interface{}{[]byte("ijklmnop")})
	polink.AddOp(package_op.EncryptRsa, false, []interface{}{PubKeyRemote, PrivKeyLocal})

	polink.AddOp(package_op.CompressGzip, false, nil)
	//polink.AddOp(package_op.CompressZlib, false, nil)

	//polink.AddOp(package_op.PacketXml, false, nil)
	polink.AddOp(package_op.PacketJson, false, nil)
	var test2 s1
	err = polink.Execute(rst1, &test2)
	if err != nil{
		fmt.Printf("err: %s", err)
		return
	}
	fmt.Printf("origin data: %v\n", test2)





	fmt.Println("--------------------------------------------------------")
	/////////////////////////// protobuf ///////////////////////////

	//polink.Reset()
	////bodyData := "guangzhou/fangcun/vip/company"
	//
	////p := &zqpacket.StringMessage{
	////	Body: proto.String(bodyData),
	////	Header: &zqpacket.Header{
	////		MessageId: proto.String("20-05"),
	////		Topic:     proto.String("golang"),
	////	},
	////}
	//
	//polink.AddOp(package_op.PacketJson, true, nil)
	////polink.AddOp(package_op.PacketXml, true, nil)
	////polink.AddOp(package_op.PacketProtobuf, true, nil)
	//
	//polink.AddOp(package_op.CompressGzip, true, nil)
	////polink.AddOp(package_op.CompressZlib, true, nil)
	//
	////polink.AddOp(package_op.EncryptAes, true, []interface{}{[]byte("abcdefghijklmnop")})
	////polink.AddOp(package_op.EncryptDes, true, []interface{}{[]byte("ijklmnop")})
	//polink.AddOp(package_op.EncryptRsa, true, []interface{}{PubKeyRemote, PrivKeyLocal})
	//var rst11 []byte
	//err = polink.Execute(p, &rst11)
	//if err != nil{
	//	fmt.Printf("err: %s", err)
	//	return
	//}
	//
	//
	//
	////var protobufrst zqpacket.StringMessage
	//polink.Reset()
	////polink.AddOp(package_op.EncryptAes, false, []interface{}{[]byte("abcdefghijklmnop")})
	////polink.AddOp(package_op.EncryptDes, false, []interface{}{[]byte("ijklmnop")})
	//polink.AddOp(package_op.EncryptRsa, false, []interface{}{PubKeyRemote, PrivKeyLocal})
	//
	//polink.AddOp(package_op.CompressGzip, false, nil)
	////polink.AddOp(package_op.CompressZlib, false, nil)
	//
	////polink.AddOp(package_op.PacketXml, false, nil)
	//polink.AddOp(package_op.PacketJson, false, nil)
	////polink.AddOp(package_op.PacketProtobuf, false, nil)
	//err = polink.Execute(rst11, &protobufrst)
	//if err != nil{
	//	fmt.Printf("err: %s", err)
	//	return
	//}
	//fmt.Printf("origin data: %v\n", &protobufrst)
}

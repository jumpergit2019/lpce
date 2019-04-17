package packet

import (
	"github.com/jumpergit2019/lpce/log"
	"encoding/xml"
	"fmt"
)

type PacketOpXml struct{
	direct bool
}



func (self *PacketOpXml) Init(direct bool, params []interface{}) bool{
	self.direct = direct
	return true
}

func (self *PacketOpXml) Operate(input interface{}, output interface{}) (bool,error){

	if self.direct{
		tmpOutput, err := self.Pack(input)
		if err != nil{
			fmt.Printf("pack failed. err: %s", err)
			return false, err
		}
		*(output.(*[]byte)) = tmpOutput
		return true, nil

	}else{
		err := self.Unpack(input.([]byte), output)
		if err != nil{
			fmt.Printf("unpack failed. err: %s", err)
			return false, err
		}
		return true, nil
	}

	return true, nil
}
//todo: xml是不能对 map 编码的, 这里需要添加检查
//https://stackoverflow.com/questions/30928770/marshall-map-to-xml-in-go?utm_medium=organic&utm_source=google_rich_qa&utm_campaign=google_rich_qa
func (*PacketOpXml) Pack(originData interface{}) ([]byte, error) {
	defer log.TraceLog("PacketOpXml.Pack")()
	return xml.Marshal(originData)
}

func (*PacketOpXml) Unpack(packData []byte, obj interface{}) error {
	defer log.TraceLog("PacketOpXml.Unpack")()

	//fmt.Println("type: ", reflect.ValueOf(obj).Type())
	err := xml.Unmarshal(packData, obj)
	//fmt.Println("value: ", reflect.ValueOf(obj).Interface())
	return err
}

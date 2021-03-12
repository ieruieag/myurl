package main

import (
	"encoding/binary"
	"encoding/hex"
	"time"
)

//simple "2dd01b3d0f246c178d1d58b328226483"
func GetSession(token string) (session string){
	seed := uint32(time.Now().Unix())
	seed = uint32(712773359)
	bytestr,ok := hex.DecodeString(token)
	uint32Len := 4	//4byte
	if ok != nil{
		return
	}

	begin := 0
	for i:=0; i<4; i++{
		bytetemp := bytestr[begin:begin+uint32Len]	//"2dd01b3" "0f246c17" next...
		begin = begin + uint32Len

		println(string(hex.EncodeToString(bytetemp)))
		uinttemp := binary.LittleEndian.Uint32(bytetemp)	//1025232941
		blur := uinttemp ^ seed	//seed = 712773359  blur = 392683202
		println(blur)
		blurByte := make([]byte,4)
		binary.LittleEndian.PutUint32(blurByte,blur)
		println(string(hex.EncodeToString(blurByte)))
		session += hex.EncodeToString(blurByte)
	}
	return
}

func main(){
	str := "2dd01b3d0f246c178d1d58b328226483"
	//str := "0f246c17"
	strRet := ""
	//str := "3d1bd02d"
	//str := "42101953"
	var seed uint32
	seed = 712773359
	//defe := time.Now().Unix()
	//println(defe)
	//seed = uint32(time.Now().UnixNano() )
	//bytestr := []byte(str)
	bytestr,_ := hex.DecodeString(str)
	begin := 0
	//32位的字符串,相当于4个uint32
	for i:=0; i<4; i++{
		bytetemp := bytestr[begin:begin+4]
		println(string(hex.EncodeToString(bytetemp)))
		uinttemp := binary.LittleEndian.Uint32(bytetemp)
		ret := uinttemp ^ seed
		println(ret)
		by11 := make([]byte,4)
		binary.LittleEndian.PutUint32(by11,ret)
		hex.EncodeToString(by11)
		//binary.BigEndian.PutUint32(by11,ret)
		println(string(hex.EncodeToString(by11)))
		begin = begin + 4

		strRet += string(hex.EncodeToString(by11))
	}
	println(strRet)
	println(GetSession(str))
}

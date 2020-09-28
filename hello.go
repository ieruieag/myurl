package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/eoscanada/eos-go"
)

func test() {
	d := "abcde"
	print(d)
}

type Card struct {
	Id 					uint64 		`json:"id"` //
	Req_id 				uint64 		`json:"req_id"` //
	Vid_shoe 			string		`json:"vid_shoe"` //
	Vid_shoe_index 		uint64 		`json:"vid_shoe_index"` //
	Round 				string 		`json:"round"` //
	Poker_index 		uint32		`json:"poker_index"` //
	Poker_code 			string 		`json:"poker_code"` //
	Reveal_index 		uint32 		`json:"reveal_index"` //
	Tx_id				string		`json:"tx_id"` //
}

func main() {
	print("hello world\n")
	test()
	api:=eos.New("https://jungle3.cryptolions.io:443")
	keyBag := &eos.KeyBag{}
	err := keyBag.ImportPrivateKey(context.Background(),"5JU4GwLhSXof6ZtKhYwFsSrjTugkAqRgvmqfBMEXr5sPmt7cgYE")
	if err != nil {
		panic(fmt.Errorf("get info: %s", err))
	}

	info, err := api.GetInfo(context.Background())
	if err != nil {
		panic(fmt.Errorf("get info: %s", err))
	}
	print(info)

	params := &eos.GetTableRowsRequest{}
	params.Code = "liontestea11"
	params.Scope = "liontestea11"
	params.Table = "shoes"
	params.JSON = true
	params.Limit = 416
	params.KeyType = "i64"
	params.Index = "secondary"
	params.LowerBound = "4"
	params.UpperBound = "4"
	rowresp,err := api.GetTableRows(context.Background(),*params)
	//var data interface{}
	//bytes := json.Unmarshal(rowresp.Rows,&data)

	var data []Card
	bytes := json.Unmarshal(rowresp.Rows,&data)
	print(bytes)
}

//t1 := `{"type":"a", id:"aaa"}`t2 := `{"type":"b", id:22222}`

//type Data struct {
//	Type string      `json:"type"`
//	Id   interface{} `json:"id"`}
//
//func decode(t string) {
//	var x Data
//	err := json.Unmarshal([]byte(t), &x)
//	if err != nil {        panic(err)
//	}
//	if x.Type == "a" {
//		fmt.Println(x.Id.(string))
//	} else {
//		fmt.Println(x.Id.(float64)) //json解析中number默认作为float64解析
//	}
//}
//
//func main() {
//	t1 := `{"type":"a", "id":"aaa"}`
//	t2 := `{"type":"b", "id":22222}`
//
//	decode(t1)
//	decode(t2)
//}

//type Resp struct {
//	Type string          `json:"type"`
//	Data json.RawMessage `json:"data"`
//}
//type Data struct {
//	Id json.Number `json:"id"` //处理大数
//}
//
//func main() {
//	t := `{"type": "a", "data":{"id": 1234567890123456789012345}}`
//
//	var x Resp
//	//var y Data
//	var y interface{}
//	json.Unmarshal([]byte(t), &x)
//
//	//进一步解组
//	if "a" == x.Type {
//		json.Unmarshal(x.Data, &y)
//	}
//
//	fmt.Println(y)
//
//	r, _ := json.Marshal(x)
//	fmt.Println(string(r))
//}
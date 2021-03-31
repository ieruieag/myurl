package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	createPtr := flag.String("create", "", "create channel vid")
	typePtr := flag.Int("type", 1, "channel type")
	deletePtr := flag.String("delete", "", "create channel vid")

	// 进行flag解析
	flag.Parse()
	println(*createPtr,*typePtr,*deletePtr)
}

func doArgs(){
	vidPtr := flag.String("vid", "", "create channel vid")
	typePtr := flag.Int("type", 0, "channel type")
	flag.Parse()
	vid := *vidPtr
	typevid := *typePtr

	fn := func(typev int){
		switch os.Args[1] {
		case "create":
			fallthrough
		case "c":

			break
		case "delete":
			fallthrough
		case "d":

			break

		}
	}
	switch typevid {
	case 0:
		fn(1)
		fn(2)
	case 1:
		fallthrough
	case 2:
		fn(typevid)
		break
	}
}

func main() {

}
package main

import (
	"CoHvs/client"
	"CoHvs/server"
	"CoHvs/utils"
	"fmt"
)

func main() {

	utils.PrintBanner()

	if utils.FindArg("server") {
		if utils.FindArg("cson") {
			server.SwitchOnCompatiSymmetric()
		}
		fmt.Println("【start up as server mode】")
		server.GetMultiServer().Work()
	} else if utils.FindArg("mock") {
		fmt.Println("【start up as mock mode】")
		client.Mock()
	} else {
		fmt.Println("【start up as client mode】")
		client.Work()
	}

}

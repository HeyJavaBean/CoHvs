package main

import (
	"CoHvs/client"
	"CoHvs/server"
	"fmt"
	"os"
)

func main() {

	fmt.Println(
		"   _____        _    _             \n" +
		"  / ____|      | |  | |            \n" +
			" | |      ___  | |__| |__   __ ___ \n" +
			" | |     / _ \\ |  __  |\\ \\ / // __|\n" +
			" | |____| (_) || |  | | \\ V / \\__ \\\n" +
			"  \\_____|\\___/ |_|  |_|  \\_/  |___/  v0.3.0  \n" +
			"                                  --By HeyJavaBean")

	serverMode := false
	for _, arg := range os.Args {
		if arg=="server"{
			serverMode = true
		}
		if arg=="cson"{
			server.SwitchOnCompatiSymmetric()
		}
		if arg=="mock"{
			fmt.Println("start up as mock mode")
			client.Mock()
			return
		}
	}



	if serverMode{
		fmt.Println("【start up as server mode】")
		server.GetMultiServer().Work()
	}else{
		fmt.Println("【start up as client mode】")
		client.Work()
	}



}

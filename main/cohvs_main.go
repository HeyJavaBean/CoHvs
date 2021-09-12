package main

import (
	"CoHvs/client"
	"CoHvs/server"
	"fmt"
	"os"
)

func main() {

	serverMode := false
	for _, arg := range os.Args {
		if arg=="server"{
			serverMode = true
			break
		}
	}

	if serverMode{
		fmt.Println("start up as server mode")
		server.GetBumpServer().Work()
	}else{
		fmt.Println("start up as client mode")
		client.Work()
	}
}

package server

import (
	"CoHvs/utils"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

type Shell struct {
	cmdMap map[string]func()int
	server *MultiServer
}

func GetShell(server *MultiServer) *Shell {
	sh := new(Shell)
	sh.server = server
	sh.cmdMap = map[string]func()int{}
	sh.cmdMap["quit"] = sh.quit
	sh.cmdMap["ping"] = sh.ping
	sh.cmdMap["list"] = sh.listPlayer
	sh.cmdMap["log"] = sh.log
	return sh
}

func (sh *Shell) Work(){
	fmt.Println("Interactive Mode is on!")
	for {
		var input string
		fmt.Print("【CoHvs-Cli】>")
		fmt.Scanln(&input)
		if len(strings.TrimSpace(input))<1{
			continue
		}
		cmd := sh.cmdMap[input]
		if cmd==nil{
			fmt.Println("【wrong cmd!】")
		}else{
			ret := cmd()
			if ret==1{
				break
			}
		}
	}
}

func (sh *Shell)ping() int{
	fmt.Println("pong")
	return 0
}

func (sh *Shell)quit() int{
	fmt.Println("bye!")
	return 1
}

func (sh *Shell)log() int{
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	utils.UseLog()
	//设置要接收的信号
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs
		done <- true
	}()

	<-done
	signal.Stop(sigs)
	utils.UnuseLog()

	return 0
}

func (sh *Shell)listPlayer() int{
	list := sh.server.controller.listPlayers()
	if len(list)<1{

		utils.PrintLog("No players is online now.")
	}else{
		utils.PrintLog("============")
		for _, p := range list {
			utils.PrintLog(p.id)
			utils.PrintLog(p.playerAddr)
			utils.PrintLog(p.netter.port)
			utils.PrintLog("============")
		}
	}
	return 0
}


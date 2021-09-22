package server

import (
	"CoHvs/constant"
	"CoHvs/utils"
	"fmt"
	"net"
	"strconv"
)

//MultiServer
//can only handle port restricted cone nat model
//https://zhuanlan.zhihu.com/p/86759357
type MultiServer struct {
	controller *VPlayerController
	mapConn    *net.UDPConn
}


//GetMultiServer
//Init the server, bind ports...
func GetMultiServer() *MultiServer {

	//listen map info
	mapConn := utils.GetUdpListenConn(utils.GetMapPort())

	controller := new(VPlayerController)
	//bind n VPlayer to the port
	vplayers := []VPlayer{}
	for i:=0;i<constant.MaxNum;i++{
		port:=utils.GetPlayerPort(i)
		conn := utils.GetUdpListenConn(port)
		vplayer := VPlayer{netter: Netter{conn: conn,port:port},controller: controller,id: "vp-"+strconv.Itoa(i)}
		vplayers=append(vplayers, vplayer)
	}

	controller.players=vplayers

	server := MultiServer{
		controller: controller,
		mapConn:    mapConn,

	}
	return &server
}

func (server *MultiServer) Work() {

	fmt.Println("Server start......")
	fmt.Println("compatiSymmetric switch is ",compatiSymmetric)

	server.controller.work()


	server.syncMap()

	//do shell job
	//todo： check player status, quit
	shell := GetShell(server)
	shell.Work()


}

//syncMap
//Get big packets from map port and broadcast them to all vplayers.
//1. host player create a room and then send udp broadcast packets about 700 bytes
//2. other players get this and reply a join packets about 200 bytes right to the host player's ip
// so, the map packets must be sent from the right port, otherwise the reply steps will fail.
func (server *MultiServer) syncMap() {
	go func() {
		for {
			msg,rAddr,err := utils.ParseMsg(server.mapConn)
			if err!=nil{
				fmt.Println(err)
			}else{
				utils.PrintLog("【Map Info From ",rAddr.String(),"】")
				server.controller.broadcast(msg)
			}
		}
	}()
}




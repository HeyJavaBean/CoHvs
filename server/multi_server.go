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
		conn := utils.GetUdpListenConn(utils.GetPlayerPort(i))
		vplayer := VPlayer{netter: Netter{conn: conn},controller: controller,id: "vp-"+strconv.Itoa(i)}
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

	server.controller.work()

	//go xxx
	server.syncMap()

	//do shell job
	//todo： check player status, quit


}

//syncMap
//Get big packets from map port and broadcast them to all vplayers.
func (server *MultiServer) syncMap() {
	for {
		msg,rAddr,err := utils.ParseMsg(server.mapConn)
		if err!=nil{
			fmt.Println(err)
		}else{
			fmt.Println("【Map Info From ",rAddr.String(),"】")
			server.controller.broadcast(msg)
		}
	}
}




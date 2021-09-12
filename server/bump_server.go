package server

import (
	"CoHvs/constant"
	"CoHvs/utils"
	"fmt"
	"net"
	"strconv"
)

type BumpServer struct {
	players []*net.UDPAddr
	conn    *net.UDPConn
}



func GetBumpServer() *BumpServer {
	udpAddr, _ := net.ResolveUDPAddr("udp", ":"+strconv.Itoa(constant.ServerUdpPort))
	conn, _ := net.ListenUDP("udp", udpAddr)
	server := BumpServer{
		players: []*net.UDPAddr{},
		conn:    conn,
	}
	go server.startRpc()
	return &server
}

func (server *BumpServer) startRpc() {

	udpAddr, _ := net.ResolveUDPAddr("udp", ":"+strconv.Itoa(constant.ServerRpcPort))
	conn, _ := net.ListenUDP("udp", udpAddr)
	for {

		msg,_,err := utils.ParseMsg(conn)
		if err!=nil{
			fmt.Println(err)
			continue
		}

		fmt.Println("bump server do broadcast!")
		for _, player := range server.players {
			server.conn.WriteToUDP(msg, player)
		}

	}
}






func (server *BumpServer) Work() {

	fmt.Println("now bump server start to work!")
	for {

		msg,raddr,_ := utils.ParseMsg(server.conn)
		fmt.Println("Get:",msg," from ",raddr.IP,":",raddr.Port)
		notExist := true
		for _, player := range server.players {
			if player.String()==raddr.String(){
				notExist = false
			}else{
				server.conn.WriteToUDP(msg, player)
			}
		}
		if notExist{
			fmt.Println("do registry")
			server.players = append(server.players, raddr)
			go utils.ConnectGame(server.conn,raddr)
		}
	}

}



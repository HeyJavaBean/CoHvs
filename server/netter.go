package server

import (
	"CoHvs/utils"
	"net"
)

type Netter struct{
	conn    *net.UDPConn
}

func (netter *Netter) send(msg []byte,addr *net.UDPAddr){
	if addr!=nil{
		netter.conn.WriteToUDP(msg,addr)
	}
}

func (netter *Netter) deal(handlerFunction func([]byte, *net.UDPAddr)) {

	for {
		msg,raddr,_ := utils.ParseMsg(netter.conn)
		handlerFunction(msg,raddr)
	}

}

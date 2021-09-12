package utils

import (
	"CoHvs/constant"
	"net"
	"time"
)

func ParseMsg(conn *net.UDPConn) ([]byte, *net.UDPAddr,error) {
	var buf [2048]byte
	n, raddr, err := conn.ReadFromUDP(buf[0:])
	msg := buf[0:n]
	return msg,raddr,err
}

func ConnectGame(conn *net.UDPConn,raddr *net.UDPAddr) {
	for i := 0; i < 20; i++ {
		conn.WriteToUDP(constant.GameRequestPacket, raddr)
		time.Sleep(5 * time.Second)
	}
}

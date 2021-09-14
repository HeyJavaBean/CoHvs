package utils

import (
	"CoHvs/constant"
	"net"
	"strconv"
	"time"
)

func ParseMsg(conn *net.UDPConn) ([]byte, *net.UDPAddr,error) {
	var buf [2048]byte
	n, raddr, err := conn.ReadFromUDP(buf[0:])
	msg := buf[0:n]
	return msg,raddr,err
}

func GetUdpListenConn(port int) *net.UDPConn{
	udpAddr, _ := net.ResolveUDPAddr("udp", ":"+strconv.Itoa(port))
	conn, _ := net.ListenUDP("udp", udpAddr)
	return conn
}

func ConnectGame(conn *net.UDPConn,raddr *net.UDPAddr) {
	for i := 0; i < 20; i++ {
		conn.WriteToUDP(constant.GameRequestPacket, raddr)
		time.Sleep(5 * time.Second)
	}
}

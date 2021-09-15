package client

import (
	"CoHvs/constant"
	"fmt"
	"net"
	"time"
)

// ping
// Send udp packet from local port 6112 to the game server so the game server will continuously send back CoH game packets, then start your game to receive these packets.
func ping(ip string,port int) {

	fmt.Println("\ntrying to ping ",ip,":",port)
	// the default CoH game port is 6112
	lAddr := &net.UDPAddr{IP: net.ParseIP("localhost"), Port: constant.GamePort}
	rAddr := &net.UDPAddr{IP: net.ParseIP(ip), Port: port}
	conn, err := net.DialUDP("udp", lAddr, rAddr)
	defer conn.Close()
	if err != nil {
		fmt.Println(err)
	}

	var buf [2048]byte

	// loop trying to ping the server
	conn.SetReadDeadline(time.Now().Add(time.Second*5))
	for {
		fmt.Println("ping....")
		conn.Write([]byte("Hey!"))

		n,timeoutErr := conn.Read(buf[0:])
		if n>0 {
			break
		}
		if timeoutErr != nil{
			conn.SetReadDeadline(time.Now().Add(time.Second*5))
		}
	}


}


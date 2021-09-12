package client

import (
	"fmt"
	"net"
	"time"
)

// ping
// Send udp packet from local port 6112 to the game server so the game server will continuously send back CoH game packets, then start your game to receive these packets.
func ping(serverIp string,serverUdpPort int) {

	// the default CoH game port is 6112
	lAddr := &net.UDPAddr{IP: net.ParseIP("localhost"), Port: 6112}
	rAddr := &net.UDPAddr{IP: net.ParseIP(serverIp), Port: serverUdpPort}
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
		conn.Write([]byte("Let's Play CoH!"))

		n,timeoutErr := conn.Read(buf[0:])
		if n>0 {
			break
		}
		if timeoutErr != nil{
			conn.SetReadDeadline(time.Now().Add(time.Second*5))
		}
	}


}


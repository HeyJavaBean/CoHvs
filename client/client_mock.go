package client

import (
	"CoHvs/constant"
	"CoHvs/utils"
	"fmt"
	"net"
	"time"
)

func Mock() {
	// bind port 6112 and receive udp packets
	conn := utils.GetUdpListenConn(constant.GamePort)
	go func() {
		for {
			time.Sleep(65 * time.Second)
			rAddr := &net.UDPAddr{IP: net.ParseIP("255.255.255.255"), Port: 6112}
			conn.WriteToUDP([]byte("thisisaverybigmapinfothisisaverybigmapinfothisisaverybigmapinfothisisaverybigmapinfothisisaverybigmapinfothisisaverybigmapinfothisisaverybigmapinfothisisaverybigmapinfothisisaverybigmapinfothisisaverybigmapinfothisisaverybigmapinfothisisaverybigmapinfothisisaverybigmapinfothisisaverybigmapinfothisisaverybigmapinfothisisaverybigmapinfothisisaverybigmapinfothisisaverybigmapinfothisisaverybigmapinfothisisaverybigmapinfothisisaverybigmapinfothisisaverybigmapinfothisisaverybigmapinfothisisaverybigmapinfothisisaverybigmapinfothisisaverybigmapinfothisisaverybigmapinfothisisaverybigmapinfothisisaverybigmapinfothisisaverybigmapinfothisisaverybigmapinfothisisaverybigmapinfothisisaverybigmapinfothisisaverybigmapinfothisisaverybigmapinfothisisaverybigmapinfothisisaverybigmapinfothisisaverybigmapinfothisisaverybigmapinfothisisaverybigmapinfothisisaverybigmapinfothisisaverybigmapinfothisisaverybigmapinfothisisaverybigmapinfothisisaverybigmapinfothisisaverybigmapinfothisisaverybigmapinfothisisaverybigmapinfothisisaverybigmapinfothisisaverybigmapinfothisisaverybigmapinfo"), rAddr)
			fmt.Println("send map!")
		}
	}()

	fmt.Println("mock game start!")

	var raddrs []*net.UDPAddr
	go func() {
		for {
			msg, raddr, _ := utils.ParseMsg(conn)
			raddrs = append(raddrs, raddr)
			if len(msg) > 300 {
				fmt.Println("【Get A Map Info】")
			} else {
				fmt.Println(raddr.String()+" => :", string(msg))
			}
		}
	}()

	i := 1
	for {
		if len(raddrs) == 0 {
			continue
		}
		raddr := raddrs[(len(raddrs)%i)]
		i++
		fmt.Print("【Input】:")
		var line string
		fmt.Scanln(&line)
		conn.WriteToUDP([]byte("Company Of Heroes MOCK_DATA:"+line), raddr)
		fmt.Println("send 【", line, "】 to ", raddr.String())
	}

}

package client

import (
	"CoHvs/constant"
	"CoHvs/utils"
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"net"
	"strings"
)

func sniff() {

	//todo: A more accessible way to pick nic is needed.
	//roughly pick the first nic to sniff the game packets
	devices, err := pcap.FindAllDevs()
	deviceName := devices[0].Name
	if err!=nil{
		fmt.Println(err)
		return
	}

	//sniff
	handle, err := pcap.OpenLive(deviceName, int32(65535), true, pcap.BlockForever)
	if err!=nil{
		fmt.Println(err)
		return
	}
	port := uint16(6112)
	filter := getFilter(port)
	if err := handle.SetBPFFilter(filter); err != nil {
		fmt.Printf("set bpf filter failed: %v", err)
		return
	}

	defer handle.Close()

	fmt.Println("start tracing map packets on 6112 of ",deviceName)
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	packetSource.NoCopy = true
	for packet := range packetSource.Packets() {
		checkPacket(packet)
	}
}

//定义过滤器
func getFilter(port uint16) string {
	filter := fmt.Sprintf("udp and ((src port %v) or (dst port %v))",  port, port)
	return filter
}

//checkPacket
//In CoH 2601&2602, when a player create a new room, packets with map info will be broadcast
//And these packets always have such features:
//1. udp packets
//2. is broadcast and src port is 6112
//3. bigger than 400 bytes
//so we sniff on the nic to catch these packets and forward them to the server.
func checkPacket(packet gopacket.Packet) {

	if packet.NetworkLayer() != nil && packet.TransportLayer() != nil &&
		packet.TransportLayer().LayerType() == layers.LayerTypeUDP &&
		packet.NetworkLayer().LayerType() == layers.LayerTypeIPv4{
		udp := packet.TransportLayer().(*layers.UDP)
		ip := packet.NetworkLayer().(*layers.IPv4)
		if strings.Contains(udp.SrcPort.String(),"6112") && udp.Length >= constant.BigPacketLen {
			//todo: only catch packets from localhost => check ip
			fmt.Println("get ",udp.Length," ",udp.SrcPort, " from ",ip.SrcIP)

			go call(udp.Payload)
		}
	}
}



//call
//send packets by udp
//todo： udp connection can be reused.
func call(msg []byte) {

	lAddr := &net.UDPAddr{IP: net.ParseIP("localhost"), Port: constant.MapSendPort}
	rAddr := &net.UDPAddr{IP: net.ParseIP(constant.ServerIp), Port: utils.GetMapPort()}
	conn, err := net.DialUDP("udp", lAddr, rAddr)
	defer conn.Close()
	if err != nil {
		fmt.Println(err)
	}

	conn.Write(msg)
}






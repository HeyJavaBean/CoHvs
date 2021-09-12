package client

import (
	"CoHvs/constant"
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"log"
	"net"
	"strings"
)
func newsniff() {

	deviceName := getDevice()

	//deviceName := "\\Device\\NPF_{4E68440F-E437-4FD2-8114-330404948C5D}"
	snapLen := int32(65535)
	port := uint16(6112)
	filter := getFilter(port)
	fmt.Printf("device:%v, snapLen:%v, port:%v\n", deviceName, snapLen, port)
	fmt.Println("filter:", filter)

	//打开网络接口，抓取在线数据
	handle, err := pcap.OpenLive(deviceName, snapLen, true, pcap.BlockForever)
	if err != nil {
		fmt.Printf("pcap open live failed: %v", err)
		return
	}

	// 设置过滤器
	if err := handle.SetBPFFilter(filter); err != nil {
		fmt.Printf("set bpf filter failed: %v", err)
		return
	}
	defer handle.Close()

	// 抓包
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




//检测监听到的包是否是需要转发的广播包
func checkPacket(packet gopacket.Packet) {


	if packet.NetworkLayer() != nil && packet.TransportLayer() != nil && packet.TransportLayer().LayerType() == layers.LayerTypeUDP {

		udp := packet.TransportLayer().(*layers.UDP)

		if strings.Contains(udp.SrcPort.String(),"6112") && udp.Length >= constant.BigPacketLen {
			fmt.Println("get ",udp.Length," ",udp.SrcPort)
			fmt.Println("find and send")
			go call(udp.Payload)
		}
	}
}

//通过rpc给服务器发送广播内容
func call(msg []byte) bool {

	// the default CoH game port is 6112
	lAddr := &net.UDPAddr{IP: net.ParseIP("localhost"), Port: 6113}
	rAddr := &net.UDPAddr{IP: net.ParseIP(constant.ServerIp), Port: constant.ServerRpcPort}
	conn, err := net.DialUDP("udp", lAddr, rAddr)
	defer conn.Close()
	if err != nil {
		fmt.Println(err)
		return false
	}

	conn.Write(msg)

	return true
}

//选择获取合适的网卡设备的名称
func getDevice() string {
	// 得到所有的(网络)设备
	devices, err := pcap.FindAllDevs()
	if err != nil {
		log.Fatal(err)
	}
	// 打印设备信息
	//fmt.Println("Devices found:")
	//for _, d := range devices {
	//	fmt.Println(d.Name)
	//	fmt.Println(d.Flags)
	//	fmt.Println(d.Addresses)
	//	fmt.Println()
	//}
	//
	dev := devices[0]
	fmt.Println("dev use:",dev.Name)
	//for _, device := range devices {
	//
	//	if len(device.Addresses) < 1 {
	//		continue
	//	}
	//	flag:=false
	//	for _, address := range device.Addresses {
	//		if address.IP.String()=="127.0.0.1"{
	//			break
	//		}
	//		flag=true
	//	}
	//
	//	if flag{
	//		fmt.Println(device.Name)
	//		fmt.Println(device.Flags)
	//		fmt.Println(device.Addresses)
	//		fmt.Println()
	//		dev = device
	//		break
	//	}
	//
	//}
	return dev.Name
}




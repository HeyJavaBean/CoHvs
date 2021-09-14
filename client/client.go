package client

import (
	"CoHvs/constant"
	"CoHvs/utils"
	"fmt"
)

//Work
//to force CoH to get connection to the server, the following two steps is necessary:
//1. Send udp packet from local port 6112 to the game server so the game server will continuously send back CoH game packets, then start your game to receive these packets.
//2. Sniff your local network and catch those broadcast udp packets from 6112 which means the game is sending map info to other players, then forward them to the game server so the game server to let every player.
func Work() {

	//to play with n people, the client should send packets to n ports in order to mock n different players.
	for offset:=0;offset<constant.MaxNum;offset++{
		ping(constant.ServerIp,utils.GetPlayerPort(offset))
	}

	fmt.Print("\n\n\n")
	fmt.Println("=====================================")
	fmt.Println("ok then you can start the game now!")
	fmt.Println("=====================================")
	fmt.Print("\n\n\n")

	//then sniff map info packets. You can still play games if this step failed.
	sniff()

	//todo check and fix this!
	//block the program until user press the key.
	fmt.Scanln()
}

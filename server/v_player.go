package server

import (
	"CoHvs/utils"
	"fmt"
	"net"
	"time"
)

type VPlayerController struct {
	players []VPlayer
}

type VPlayer struct {
	id         string
	netter     Netter
	controller *VPlayerController
	playerAddr *net.UDPAddr
}

//handle
//1.get msg from conn
//2.check raddr to locate user (only in cone model)
//3.forward packet to other vplayer by user map
func (p *VPlayer) handle(msg []byte, raddr *net.UDPAddr) {
	ctx := p.controller

	fmt.Println(p.id, "and ", p.playerAddr == nil," :len ", len(msg), " from ", raddr.String(), " at ", time.Now().Format("2006-01-02 15:04:05"))

	// if the msg is ping msg, then send back connection packages
	if len(msg) < 10 {
		fmt.Println("【Game Connection】")
		go utils.ConnectGame(p.netter.conn, raddr)
		return
	}

	sender := ctx.findUser(raddr)

	if sender == nil {
		//register: find an empty vplayer to present as the user
		fmt.Println("register user on vplayer ...")
		//todo thread safe
		id := ctx.register(raddr)
		fmt.Println(raddr.String(),"is registered on ",id)
		//fixme: not good
		p.handle(msg,raddr)
		return
	}
	// if current vplayer has no player to present, just ignore the msg.
	// if remote player send msg to itself vplayer, also ignore it.
	fmt.Println(sender.id+" and "+p.id)
	if p.playerAddr == nil || p==sender{
		fmt.Println("ignore!")
		return
	}

	//forward
	// current vplayer send the package to target vplayer's real raddr
	fmt.Println("simply forward packets ...")
	sender.netter.send(msg, p.playerAddr)

}

func (controller *VPlayerController) findUser(addr *net.UDPAddr) *VPlayer {
	for i := range controller.players {
		p := &controller.players[i]
		if p.playerAddr != nil && p.playerAddr.String() == addr.String() {
			return p
		}
	}
	return nil
}

func (controller *VPlayerController) register(addr *net.UDPAddr) string{
	for i := range controller.players {
		p := &controller.players[i]
		if p.playerAddr == nil {
			p.playerAddr = addr
			return p.id
		}
	}
	return "no free position"
}

func (controller *VPlayerController) work() {
	for i := range controller.players {
		p := &controller.players[i]
		go p.netter.deal(p.handle)
	}
}

func (controller *VPlayerController) broadcast(msg []byte) {
	for i := range controller.players {
		p := &controller.players[i]
		p.netter.send(msg, p.playerAddr)
	}
}
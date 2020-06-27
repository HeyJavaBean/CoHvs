package com.imlehr.coh1vs.client.game;

import lombok.SneakyThrows;

import java.net.DatagramSocket;
import java.net.InetSocketAddress;
import java.net.SocketAddress;

import static com.imlehr.coh1vs.utils.PacketUtils.sendPk;
import static com.imlehr.coh1vs.utils.RequestUtils.tryAndSleep;

/**
 * @author Lehr
 * @create: 2020-06-26
 */
public class GameConnector {

    @SneakyThrows
    public static void connnect() {
        //重入6112端口
        DatagramSocket server = new DatagramSocket(6112);
        System.out.println("[正在连接服务器....]");
        //向服务器汇报自己
        SocketAddress serverSA = new InetSocketAddress("120.79.225.195", 8888);
        tryAndSleep(20,3000,server,sendPk(serverSA));
        server.close();
        System.out.println("[连接成功，请赶快启动游戏!]");
    }


}

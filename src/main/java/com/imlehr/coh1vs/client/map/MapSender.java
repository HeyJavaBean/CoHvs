package com.imlehr.coh1vs.client.map;

import java.net.DatagramPacket;
import java.net.DatagramSocket;
import java.net.InetSocketAddress;
import java.net.SocketAddress;

/**
 * @author Lehr
 * @create: 2020-06-25
 */
public class MapSender {



    private DatagramSocket server;
    private SocketAddress serverSA;

    public MapSender()
    {
        try {
            server = new DatagramSocket(8887);
            serverSA = new InetSocketAddress("120.79.225.195",8887);
            System.out.println("[地图发送服务初始化成功]");
        } catch (Exception e) {
            e.printStackTrace();
        }
    }

    public void sendMap(byte[] bytes)
    {
        DatagramPacket mapInfo = new DatagramPacket(bytes,bytes.length,serverSA);
        try{
            server.send(mapInfo);
            System.out.println("[地图信息同步成功！]");
        }catch (Exception e)
        {
            System.out.println("[地图信息同步失败...]");
        }
    }

}

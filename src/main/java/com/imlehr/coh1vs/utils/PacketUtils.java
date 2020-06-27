package com.imlehr.coh1vs.utils;

import java.net.DatagramPacket;
import java.net.SocketAddress;

/**
 * @author Lehr
 * @create: 2020-06-19
 */
public class PacketUtils {

    public static DatagramPacket sendPk(SocketAddress sa) {
        byte[] data = new byte[1024];
        DatagramPacket pack = new DatagramPacket(data, data.length,sa);
        return pack;
    }

    public static DatagramPacket createPk(byte[] data) {
        DatagramPacket pack = new DatagramPacket(data, data.length);
        return pack;
    }


    public static DatagramPacket getPk() {
        byte[] data = new byte[1024];
        DatagramPacket pack = new DatagramPacket(data, data.length);
        return pack;
    }



}

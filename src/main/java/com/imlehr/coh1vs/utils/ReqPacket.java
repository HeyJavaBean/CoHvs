package com.imlehr.coh1vs.utils;

import java.io.File;
import java.io.FileInputStream;
import java.io.InputStream;
import java.net.DatagramPacket;
import java.net.SocketAddress;

/**
 * @author Lehr
 * @create: 2020-06-26
 */
public class ReqPacket {

    private static byte[] bytes = {84,104,81,84,8,0,0,0,9,0,1,0,0,0,0,-61,61,-29,116,0};

    public static DatagramPacket getReqPacket(SocketAddress sa)
    {
        return new DatagramPacket(bytes,bytes.length,sa);
    }

}

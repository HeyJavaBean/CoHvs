package com.imlehr.coh1vs.client.map;

import org.jnetpcap.Pcap;
import org.jnetpcap.PcapIf;

import java.util.ArrayList;
import java.util.List;

/**
 * @author Lehr
 * @create: 2020-06-26
 */
public class MapListener {

    public static void listen() {
        System.out.println("地图监听器启动中...");
        List<PcapIf> devs = new ArrayList<>();
        StringBuilder errsb = new StringBuilder();
        int r = Pcap.findAllDevs(devs, errsb);
        if (r == Pcap.NOT_OK || devs.isEmpty()) {
            System.out.println("[地图监听器初始化失败...未能找到网卡]");
            return;
        }
        PcapIf device = devs.get(2);
        int snaplen = Pcap.DEFAULT_SNAPLEN;
        int flags = Pcap.MODE_PROMISCUOUS;
        int timeout = 10 * 1000;
        Pcap pcap = Pcap.openLive(device.getName(), snaplen, flags, timeout, errsb);
        if (pcap == null) {
            System.out.println("[地图监听器初始化失败...网卡绑定失败]");
            return;
        }
        System.out.println("[开始监听网卡获取地图中....]");
        new Thread(()->
        {
            pcap.loop(0, new CoHPacketHandler<>(), "jnetpcap");
            pcap.close();
        }).start();


    }

}

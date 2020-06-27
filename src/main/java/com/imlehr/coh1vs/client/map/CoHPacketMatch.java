package com.imlehr.coh1vs.client.map;

import org.jnetpcap.packet.PcapPacket;
import org.jnetpcap.protocol.tcpip.Udp;

public class CoHPacketMatch {

    private static MapSender sender = new MapSender();

    private static CoHPacketMatch pm;

    private Udp udp = new Udp();

    public static CoHPacketMatch getInstance() {
        if (pm == null) {
            pm = new CoHPacketMatch();
        }
        return pm;
    }

    public void handlePacket(PcapPacket packet) {
        if (packet.hasHeader(udp)) {
            handleUdp(packet);
        }
    }

    private void handleUdp(PcapPacket packet) {
        packet.getHeader(udp);
        String srcPort = String.valueOf(udp.source());
        if (srcPort.equals("6112")) {
            if (udp.getPayloadLength() > 400 && udp.getPayloadLength() < 1024)
            {
                sender.sendMap(udp.getPayload());
            }
        }
    }
}
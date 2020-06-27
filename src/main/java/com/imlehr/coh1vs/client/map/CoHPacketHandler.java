package com.imlehr.coh1vs.client.map;
 
import org.jnetpcap.packet.PcapPacket;
import org.jnetpcap.packet.PcapPacketHandler;

/**
 * @author lehr
 */
public class CoHPacketHandler<Object> implements PcapPacketHandler<Object>  {
     
    @Override
    public void nextPacket(PcapPacket packet, Object obj) {
        CoHPacketMatch packetMatch = CoHPacketMatch.getInstance();
        packetMatch.handlePacket(packet);
    }
}
package com.imlehr.coh1vs.server.game;

import com.imlehr.coh1vs.utils.ReqPacket;
import lombok.SneakyThrows;

import java.net.DatagramPacket;
import java.net.DatagramSocket;
import java.net.SocketAddress;
import java.util.concurrent.ExecutorService;
import java.util.concurrent.Executors;

import static com.imlehr.coh1vs.utils.PacketUtils.getPk;
import static com.imlehr.coh1vs.utils.ReqPacket.getReqPacket;
import static com.imlehr.coh1vs.utils.RequestUtils.loopSend;

/**
 * @author Lehr
 * @create: 2020-06-26
 */
public class GameServer {

    private static boolean flag = true;

    private SocketAddress player1;
    private SocketAddress player2;

    private DatagramSocket server;

    private static ExecutorService pool;

    @SneakyThrows
    public GameServer() {
        server = new DatagramSocket(8888);
        pool = Executors.newFixedThreadPool(10);
        System.out.println("[服务器初始化完毕]");
    }

    @SneakyThrows
    public boolean repost(DatagramPacket pk) {
        SocketAddress sa = pk.getSocketAddress();
        if (player1 != null && player2 != null) {
            //转发
            if (sa.equals(player2)) {
                System.out.println("[转发报文!" + player2 + "->" + player1 + "]");
                pk.setSocketAddress(player1);
                server.send(pk);
            }
            if (sa.equals(player1)) {
                System.out.println("[转发报文!" + player1 + "->" + player2 + "]");
                pk.setSocketAddress(player2);
                server.send(pk);
            }
            return true;
        }
        return false;
    }


    @SneakyThrows
    public void work() {

        System.out.println("[正在等待玩家...]");

        while (flag) {
            if (player1 != null && player2 != null) {
                break;
            }

            DatagramPacket pk = getPk();
            server.receive(pk);
            SocketAddress sourceSA = pk.getSocketAddress();

            if (player1 == null && !sourceSA.equals(player2)) {
                player1 = sourceSA;
                System.out.println("[发现玩家1：" + sourceSA + "]");
                DatagramPacket pack = getReqPacket(sourceSA);
                //保持心跳
                loopSend(server, pack);
                continue;
            }
            if (player2 == null && !sourceSA.equals(player1)) {
                player2 = sourceSA;
                System.out.println("[发现玩家2：" + sourceSA + "]");
                DatagramPacket pack = getReqPacket(sourceSA);
                loopSend(server, pack);
                continue;
            }
            System.out.println("[等待另一位玩家上线....]");
        }

        System.out.println("玩家上线，开始数据同步...");

        while (flag) {
            DatagramPacket pk = getPk();
            server.receive(pk);
            pool.execute(()->repost(pk));
        }

    }

    @SneakyThrows
    public void sendMap(DatagramPacket packet) {
        if (player1 != null && player2 != null) {
            packet.setSocketAddress(player1);
            server.send(packet);
            packet.setSocketAddress(player2);
            server.send(packet);
        }
    }
}

package com.imlehr.coh1vs.server.map;

import com.imlehr.coh1vs.server.game.GameServer;
import lombok.SneakyThrows;

import java.net.DatagramPacket;
import java.net.DatagramSocket;
import java.util.concurrent.ExecutorService;
import java.util.concurrent.Executors;

import static com.imlehr.coh1vs.utils.PacketUtils.getPk;

/**
 * @author Lehr
 * @create: 2020-06-26
 */
public class MapServer {

    private static boolean flag = true;

    private GameServer gameServer;

    private DatagramSocket server;

    private static ExecutorService pool;

    @SneakyThrows
    public MapServer(GameServer gameServer)
    {
        this.gameServer = gameServer;
        server = new DatagramSocket(8887);
        pool = Executors.newFixedThreadPool(1);
        System.out.println("[地图收发模块启动成功]");
    }


    @SneakyThrows
    private void sendMap()
    {
        DatagramPacket pk = getPk();
        server.receive(pk);
        System.out.println("收到地图！");
        //让GameServer去发送
        gameServer.sendMap(pk);
    }

    @SneakyThrows
    public void work()
    {
        System.out.println("地图监听器开始工作");
        pool.execute(()->
        {
            while(flag)
            {
                sendMap();
            }
        });
    }
}

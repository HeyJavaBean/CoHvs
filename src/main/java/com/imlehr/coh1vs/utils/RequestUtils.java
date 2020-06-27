package com.imlehr.coh1vs.utils;

import lombok.SneakyThrows;

import java.net.DatagramPacket;
import java.net.DatagramSocket;
import java.util.concurrent.ExecutorService;
import java.util.concurrent.Executors;
import java.util.concurrent.atomic.AtomicBoolean;
import java.util.concurrent.atomic.AtomicInteger;

import static com.imlehr.coh1vs.utils.PacketUtils.getPk;

/**
 * @author Lehr
 * @create: 2020-06-19
 */
public class RequestUtils {

    private static ExecutorService pool = Executors.newFixedThreadPool(10);

    /**
     *
     * @param count   尝试的次数，每次间隔1秒钟
     * @param server    哪个socket
     * @param sleepTime   每次尝试的间隔时间，单位是1000分之一秒
     * @param packet    尝试发送的包
     * @return
     */
    @SneakyThrows
    public static DatagramPacket tryAndSleep(Integer count, Integer sleepTime, DatagramSocket server,DatagramPacket packet)
    {
        AtomicBoolean success = new AtomicBoolean(false);
        AtomicInteger times = new AtomicInteger(count);

        pool.execute(()->
        {
            while(!success.get())
            {
                if(times.decrementAndGet()>0)
                {
                    System.out.println("[尝试中....]");
                    try{
                        server.send(packet);
                        Thread.sleep(sleepTime);
                    }
                    catch (Exception e)
                    {
                        System.out.println("[出现错误，忽略....]");
                    }
                }
                else
                {
                    System.out.println("[执行失败....]");
                    System.exit(-1);
                }
            }
        });
        Thread.sleep(2000);
        DatagramPacket pk = getPk();
        server.receive(pk);
        success.set(true);
        return pk;
    }


    @SneakyThrows
    public static void loopSend(DatagramSocket server, DatagramPacket pack) {
        System.out.println("开始同步");
        Integer count = 3;
        //连续发送3个建立同步
        while (count-- > 0) {
            server.send(pack);
        }

        //保持心跳5分钟
        pool.execute(()->
        {
            Integer county = 30;
            while (county-- > 0) {
                try {
                    server.send(pack);
                    Thread.sleep(10000);
                } catch (Exception e) {

                }
            }
        });
    }


}

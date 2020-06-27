package com.imlehr.coh1vs.client;

import com.imlehr.coh1vs.client.game.GameConnector;
import com.imlehr.coh1vs.client.map.MapListener;

/**
 * @author Lehr
 * @create: 2020-06-26
 */
public class ClientMain {

    public static void main(String[] args) {

        //本地和服务器进行连线
        GameConnector.connnect();
        //启动地图监听器
        MapListener.listen();


    }

}

package com.imlehr.coh1vs.server;

import com.imlehr.coh1vs.server.game.GameServer;
import com.imlehr.coh1vs.server.map.MapServer;

/**
 * @author Lehr
 * @create: 2020-06-26
 */
public class ServerMain {


    public static void main(String[] args) {

        GameServer gameServer = new GameServer();

        MapServer mapServer = new MapServer(gameServer);

        mapServer.work();

        gameServer.work();




    }

}

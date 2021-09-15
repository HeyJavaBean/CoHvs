# CoHvs
> Company of Heroes 1 Over-LANs PVP Program

Beta Version, with CoHvs, you and your friends can play CoH in different LANs.
Now I'm still working on it, so in complicated network enviroments (like hotspots) it may not works.

To play CoH with your friends in different LANs, do as the following steps:
- Get CoHvs
  - Download CoHvs in "releases" and chose the right version with your OS.
  - The CoHvs can be both client and server.
- Server Side:
  - Get a server with a public ip (in constant.go and the default value is my server ip)
  - Execute “**./cohvs server** ” in order to run the program in server mode.
  - Wait the client side to start.
  - Ps:  if your friends is using mobile phone hotspots or the NAT server is Symmetric NAT, trying to use "**./cohvs server cson**" to start the server.
- Client Side:
  - Execute "**./cohvs.exe**" by simply click the exe icon or using cmd.
  - When client successfully get connection with the server, then start CoH.
  - If the client shell says "sniffer is starting", you can create a new game room, otherwise you can only join other players' room (but it's still ok to play the game).
  - Only the first guy connecting to the server can create room. Other players then just join in the game. Cuz I screwed up the map sending logic :(
    - ↑ I'll fix it asap... 
  - Enjoy the game!

--------------------


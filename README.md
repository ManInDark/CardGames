# Watten

This program allows for the game "[Watten](https://de.wikipedia.org/wiki/Watten)" to be played digitally (english [explanation](https://www.pagat.com/trumps/watten.html)). For it to be played all players need to connect via a websocket to the server.

# Setup

You only need Go to run the server, so even a GitHub Codespace is enough. To run the server execute the following commands:

    cd WattenServer
    go run .

This should automatically install all necessary packages and serve the game on port 2000, but be aware that GitHub Codespaces already use that port, so you'll have to change it to something else (like 4000). This can be done by editing the following line in WattenServer/WattenServer.go
 
     http.ListenAndServe(":2000", nil)
     
Afterwards you can open the game by navigating to :[port]/home.html wherever you installed it. The game starts automatically after enough connections have been established.

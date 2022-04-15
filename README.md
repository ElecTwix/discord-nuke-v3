# discord-nuke-v3

Proof-of-Concept

## Steps

Pre Step 1

[Page link](https://discord.com/developers/applications)

![](https://i.imgur.com/ixUWxs6.png)

Pre Step 2

![](https://i.imgur.com/hek5Yer.png)

Pre Step 3

![](https://i.imgur.com/oL6Zxd6.png)

*Save your token somewhere*



Step 1


![](https://i.imgur.com/tE5OptO.png)

Step 1

![](https://i.imgur.com/HDDfzsM.png)

![](https://i.imgur.com/aaBuM8J.png)

invite your bot 

put your client id in your link like down below

https://discordapp.com/api/oauth2/authorize?client_id=123456789&permissions=139586816064&scope=bot

Step 2 

Changing Configs in config.json

| varabile       | default           | description |
| ------------- |:-------------:| :---------------:|
| token      | "toke-here" | put your token here |
| ban      | true      | will bot gonna ban people |
| removeperm | true      | will bot remove people roles |
| deletechannels | true      | will bot delete all channels |
| command | "nuke"      | command that will trigger nuke |

Step 3

Executing The bot

| Type       | Command           |
| ------------- |:-------------:|
| Source      | go run main.go |
| Linux Build      | ./nuke-bot      |
| Windows Build | .\nuke-bot      |

Step 4

send your nuke message in any channel bot can see.

Done.



Source 
go run main.go

# Streamdeck OP-FW integration

1. Download and install the UDP streamdeck plugin [here](https://github.com/Zayik/CommandSender) ([direct-download](https://github.com/Zayik/CommandSender/raw/master/Release/com.biffmasterzay.commandsender.streamDeckPlugin))
2. Download the precompiled integration binary [here](dist/OP-FW Streamdeck.exe) or compile it yourself
3. Run the `opfw_streamdeck.exe` (you will have to run this everytime before you open FiveM)
4. It should open a command prompt saying something similar to this  
![screenshot](https://i.twoot.org/Cufo7/RImeYAcO29.png)
5. You can minimize the command prompt but keep it open while you use FiveM (it has to be running for the integration to work)
6. If you open FiveM and connect to any OP-FW server you should now see something like this in the top right  
![screenshot](https://i.twoot.org/Cufo7/FeNEbEQo25.png)
7. If you loose connection you can try to reconnect using `/reconnect_command_socket`  
![screenshot](https://i.twoot.org/Cufo7/XOrirOcE89.png)
8. Now in the streamdeck software you need to add the `CommandSender` action  
![screenshot](https://i.twoot.org/Cufo7/CutaSURO98.png)
9. Leave all settings how they are by default, but change the "Port" to `42069`  
![screenshot](https://i.twoot.org/Cufo7/zUpoBova02.png)
10. In the `Command Pressed` field you enter what ever FiveM command you want to run when you press the button on your streamdeck

You're all set, you will have to run `OP-FW Streamdeck.exe` every time you load into FiveM. If you open it after you have already loaded into FiveM you will have to run the `/reconnect_command_socket` command.
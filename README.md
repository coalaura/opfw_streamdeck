# Streamdeck OP-FW integration

1. Download and install the UDP streamdeck plugin [here](https://github.com/Zayik/CommandSender) ([direct-download](https://github.com/Zayik/CommandSender/raw/master/Release/com.biffmasterzay.commandsender.streamDeckPlugin))
2. Download the precompiled integration binary [here](dist/OP-FW Streamdeck.exe) or compile it yourself
3. Run the `OP-FW Streamdeck.exe` (you will have to run this everytime before you open FiveM)
4. It will not open a window or command prompt but add an icon to your taskbar    
![screenshot](https://i.twoot.org/Cufo7/liqEguFu74.png)
5. You can right click the icon to quit the app or open the latest logfile  
![screenshot](https://i.twoot.org/Cufo7/biyoRAYE74.png)
6. You can minimize the command prompt but keep it open while you use FiveM (it has to be running for the integration to work)
7. If you open FiveM and connect to any OP-FW server you should now see something like this in the top right  
![screenshot](https://i.twoot.org/Cufo7/FeNEbEQo25.png)
8. If you loose connection you can try to reconnect using `/reconnect_command_socket`  
![screenshot](https://i.twoot.org/Cufo7/XOrirOcE89.png)
9. Now in the streamdeck software you need to add the `CommandSender` action  
![screenshot](https://i.twoot.org/Cufo7/CutaSURO98.png)
10. Leave all settings how they are by default, but change the "Port" to `42069`  
![screenshot](https://i.twoot.org/Cufo7/zUpoBova02.png)
11. In the `Command Pressed` field you enter what ever FiveM command you want to run when you press the button on your streamdeck

You're all set, you will have to run `OP-FW Streamdeck.exe` every time you load into FiveM. If you open it after you have already loaded into FiveM you will have to run the `/reconnect_command_socket` command.

### Troubleshooting

If the streamdeck integration is not working, try following these steps:

1. Fully quit streamdeck  
![screenshot](https://i.twoot.org/Cufo7/YIBUwEXE26.png)
2. Fully quit the integration  
![screenshot](https://i.twoot.org/Cufo7/WAjISIwa67.png)
3. Start the integration back up
4. Start the streamdeck software back up
5. Run `/reconnect_command_socket` in the legacy server  
![screenshot](https://i.twoot.org/Cufo7/mEZEPEME36.png)

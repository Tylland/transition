

Run in Terminal on Windows:

ssh andreas-tyllsmark@nuclear


nuclear: 192.168.2.36
nats: 4222



### Start NATS Server
```
 nats-server -js -DV -sd ~/Temp/jetstream_storage/ -l /var/log/nats/nats-server.log
```

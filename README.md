`make build`
`scp ocpp-service pi@192.168.0.105:/home/pi/ocpp-service`
`scp service-config.env pi@192.168.0.105:/home/pi/service-config.env`

`sudo nano /lib/systemd/system/ocpp.service` >
```
[Unit]
 Description=OCPP Service
 After=multi-user.target

 [Service]
 Type=idle
 ExecStart=/home/pi/ocpp-service

 [Install]
 WantedBy=multi-user.target
```

`sudo chmod 644 /lib/systemd/system/ocpp.service`
`sudo systemctl daemon-reload`
`sudo systemctl enable ocpp.service`
`sudo chmod 755 ocpp-service`
`sudo systemctl start ocpp.service`

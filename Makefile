
.PHONY: clean build

build:
	GO111MODULE=on go get -v -d . && \
	CC=arm-linux-gnueabihf-gcc CGO_ENABLED=1 GOARCH=arm GOARM=7 GOOS=linux go build -a -o ocpp-service

build-mac-os:
	GO111MODULE=on go get -v -d . && \
	GOOS=linux GOARCH=arm GOARM=7 go build -a -o ocpp-service

build-and-push-mac:
	GO111MODULE=on go get -v -d . && \
	GOOS=linux GOARCH=arm GOARM=7 go build -a -o ocpp-service && \
	scp ocpp-service pi@192.168.0.105:/home/pi/ocpp-service

clean:
	rm -rf ocpp-service

.PHONY: clean build

build:
	GO111MODULE=on go get -v -d . && \
	CC=arm-linux-gnueabihf-gcc CGO_ENABLED=1 GOARCH=arm GOARM=7 GOOS=linux go build -a -o ocpp-service

clean:
	rm -rf ocpp-service
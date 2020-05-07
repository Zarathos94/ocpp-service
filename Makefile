
.PHONY: clean build

build:
	GO111MODULE=on go get -v -d . && \
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ocpp-service

clean:
	rm -rf ocpp-service
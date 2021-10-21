.PHONY: default
default: help ;

include .env

help:				## Show this help.
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'

build-client:			## Build client
	cd client && npm i && npm run pretty && npm run build

build-server:			## Build server
	cd server && CGO_ENABLED=0 go build -ldflags "-w" -a -o ./build/uart-mp3-player ./cmd/webserver

build-server-rpi:		## Build server for RaspberryPi
	cd server && env GOOS=linux GOARCH=arm GOARM=5 CGO_ENABLED=0 go build -ldflags "-w" -a -o ./build/uart-mp3-player-arm ./cmd/webserver

bundle-spa-server:		## Create server bundle
	cd server/cmd/webserver && rice append -i . --exec ../../build/uart-mp3-player

bundle-spa-server-rpi:		## Create server bundle for RaspberryPi
	cd server/cmd/webserver && rice append -i . --exec ../../build/uart-mp3-player-arm

start-client-dev:		## Start Client in development mode
	cd client && npm start

start-server-dev:		## Start Server in development mode
	cd server/cmd/webserver && go run main.go

start-server-prod:		## Start Server in production mode
	./server/build/uart-mp3-player

test-server:			## Test Server
	cd server && ENV=TEST go test ./... -v

build: 				## Build Application
	make build-client 
	make build-server
	make bundle-spa-server 
	make test-server 	

start-dev: 			## Start Application in development mode
	make build 
	make start-server-dev		

start-prod: 			## Start Application in production mode
	make build 
	make start-server-prod

build-rpi:			## Build application on RaspberryPi
	make build-client
	make build-server-rpi
	make bundle-spa-server-rpi

deploy-rpi:			## Deploy application on RaspberryPi
	make build-rpi
	sshpass -p "${RPI_PASSWORD}" scp server/build/uart-mp3-player-arm ${RPI_USERNAME}@${RPI_HOST}:/home/pi/go/bin
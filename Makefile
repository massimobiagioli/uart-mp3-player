.PHONY: default
default: help ;

help:				## Show this help.
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'

build-client:			## Build client
	cd client && npm i && npm run build

build-server:			## Build server
	cd server && CGO_ENABLED=0 go build -ldflags "-w" -a -o ./build/uart-mp3-player ./cmd/webserver

bundle-spa-server:		## Create server bundle
	cd server/cmd/webserver && rice append -i . --exec ../../build/uart-mp3-player

start-client-dev:		## Start Client in development mode
	cd client && npm start

start-server-dev:		## Start Server in development mode
	cd server/cmd/webserver && go run main.go

start-server-prod:		## Start Server in production mode
	./server/build/uart-mp3-player

test-server:			## Test Server
	cd server && go test ./... -v

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
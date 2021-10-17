help:           	## Show this help.
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'

build-client:       ## Build client
	cd client && npm i && npm run build

build-server:       ## Build server
	cd server && CGO_ENABLED=0 go build -ldflags "-w" -a -o ./build/uart-mp3-player ./cmd/webserver
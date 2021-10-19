# uart-mp3-player

## Customize .env
Duplicate file <code>.env.sample</code> in <code>.env</code> and customize environment variables

## Install Rice Tool

<code>
go get <a href="github.com/GeertJohan/go.rice">github.com/GeertJohan/go.rice</a>

go get <a href="github.com/GeertJohan/go.rice/rice">github.com/GeertJohan/go.rice/rice</a>
</code>

*After install Rice package, add rice binary to PATH*


## Configure NodeJs for client application

From project root/client, launch command: <code>nvm use</code>

## Start application in *development* mode

From project root, launch command : <code>make start-dev</code> 

## Start application in *production* mode

From project root, launch command : <code>make start-prod</code>  

## Open application from local machine

Open url: http://localhost/8080

## Deploy application to RaspberryPi

From project root, launch command : <code>make deploy-rpi</code>

## Open application from RaspberryPi

Once connected into RaspberryPi, launch binary: <code>/home/pi/go/bin/uart-mp3-player-arm</code>,
then open url: http://raspberry-host/8080
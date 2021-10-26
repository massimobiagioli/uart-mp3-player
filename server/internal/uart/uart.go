package uart

import (
	"log"
	"os"
	"strconv"

	"github.com/jacobsa/go-serial/serial"
)

var serialOptions = serial.OpenOptions{
	PortName:        "/dev/ttyAMA0",
	BaudRate:        9600,
	DataBits:        8,
	StopBits:        1,
	MinimumReadSize: 4,
}

var serialCommands = initSerialCommands()

func Reset() {
	command := serialCommands["RESET"]
	log.Printf("Command sent: %v\n", command)
	if os.Getenv("ENV") == "TEST" {
		return
	}
	sendCommand(command)
}

func Play(folderId string, songId string) {
	command := serialCommands["PLAY"]
	folderIdAsNum, err := strconv.Atoi(folderId)
	if err == nil {
		command[5] = byte(folderIdAsNum)
	}
	songIdAsNum, err := strconv.Atoi(songId)
	if err == nil {
		command[6] = byte(songIdAsNum)
	}
	log.Printf("Command sent: %v\n", command)
	if os.Getenv("ENV") == "TEST" {
		return
	}
	sendCommand(command)
}

func Stop(folderId string, songId string) {
	command := serialCommands["STOP"]
	log.Printf("Command sent: %v\n", command)
	if os.Getenv("ENV") == "TEST" {
		return
	}
	sendCommand(command)
}

func initSerialCommands() map[string][]byte {
	serialCommands := make(map[string][]byte)
	serialCommands["RESET"] = []byte{0x7E, 0xFF, 0x06, 0x0C, 0x00, 0x00, 0x00, 0xEF}
	serialCommands["PLAY"] = []byte{0x7E, 0xFF, 0x06, 0x0F, 0x00, 0x01, 0x01, 0xEF}
	serialCommands["STOP"] = []byte{0x7E, 0xFF, 0x06, 0x16, 0x00, 0x00, 0x00, 0xEF}

	return serialCommands
}

func sendCommand(cmd []byte) {
	port, err := serial.Open(serialOptions)
	if err != nil {
		log.Fatalf("serial.Open: %v", err)
	}

	defer port.Close()

	n, err := port.Write(cmd)
	if err != nil {
		log.Fatalf("port.Write: %v", err)
	}
	log.Println("Wrote", n, "bytes.")
}

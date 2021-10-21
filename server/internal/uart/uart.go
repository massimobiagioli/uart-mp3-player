package uart

import (
	"log"
	"os"

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
	if os.Getenv("ENV") == "TEST" {
		log.Printf("Command sent: %v\n", command)
		return
	}
	sendCommand(command)
}

func Play(folderId string, songId string) {
	log.Printf("Command PLAY (FolderId: %v, SongId: %v) not implemented yet\n", folderId, songId)
}

func Stop(folderId string, songId string) {
	log.Printf("Command STOP (FolderId: %v, SongId: %v) not implemented yet\n", folderId, songId)
}

func initSerialCommands() map[string][]byte {
	serialCommands := make(map[string][]byte)
	serialCommands["RESET"] = []byte{0x7E, 0xFF, 0x06, 0x0C, 0x00, 0x00, 0x00, 0xEF}

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

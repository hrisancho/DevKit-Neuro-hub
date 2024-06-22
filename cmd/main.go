package main

import (
	controller "DevKit-Neuro-hub/proto"
	utils "DevKit-Neuro-hub/utils"
	"fmt"
	"google.golang.org/protobuf/proto"
	"log"
	"net"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Please provide host:port")
		os.Exit(1)
	}

	// Resolve the string address to a UDP address
	udpAddr, err := net.ResolveUDPAddr("udp", os.Args[1])

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Start listening for UDP packages on the given address
	conn, err := net.ListenUDP("udp", udpAddr)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Read from UDP listener in endless loop
	var buf [1024]byte
	// Read from UDP listener in endless loop
	for {
		n, addr, err := conn.ReadFromUDP(buf[:])
		if err != nil {
			log.Fatalln(err)
		}
		data := &controller.MessageWitchCRC{}
		err = proto.Unmarshal(buf[0:n], data)
		if err != nil {
			log.Fatalln(err)
		}
		// Функция проверки сообщения на подлинность
		err = utils.AuthenticationCRC(data)
		if err != nil {
			log.Println(err)
		} else {
			// Продолжаем работу с пакетом

			//TODO в случае чего поменять сообщение из protobuf
			msg := &controller.ChannelsDataSet{}

			err = proto.Unmarshal(data.Msg[0:], msg)
			if err != nil {
				log.Fatalln(err)
			}
			log.Print("> ", data)
			log.Println(msg)

			// Write back the message over UPD
			conn.WriteToUDP([]byte("Hello UDP Client\n"), addr)
		}
	}
}

package main

import (
	controller "DevKit-Neuro-hub/proto"
	"bufio"
	"fmt"
	"github.com/snksoft/crc"
	"google.golang.org/protobuf/proto"
	"log"
	"net"
	"os"
)

func main() {

	if len(os.Args) == 1 {
		fmt.Println("Please provide host:port to connect to")
		os.Exit(1)
	}

	// Resolve the string address to a UDP address
	udpAddr, err := net.ResolveUDPAddr("udp", os.Args[1])

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Dial to the address with UDP
	conn, err := net.DialUDP("udp", nil, udpAddr)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	//Сообщение внутри дурного сообщения
	data1 := &controller.ChannelsDataSet{
		Channel1: 123.4,
		Channel2: 223.5,
		Channel3: 323.6,
		Id:       34,
	}
	msg1, err := proto.Marshal(data1)
	if err != nil {
		log.Fatalln(err)
	}

	data := &controller.MessageWitchCRC{
		Msg:  msg1,
		Kind: controller.CRCType_crc32,
		Crc:  int64(crc.CalculateCRC(crc.CRC32, msg1)),
	}

	msg, err := proto.Marshal(data)
	if err != nil {
		log.Fatalln(err)
	}

	// Send a message to the server
	_, err = conn.Write(msg)
	fmt.Println("send...")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Read from the connection untill a new line is send
	dat, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}

	// Print the data read from the connection to the terminal
	fmt.Print("> ", string(dat))
}

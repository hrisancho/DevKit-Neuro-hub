package server

import (
	"DevKit-Neuro-hub/internal/config"
	utils "DevKit-Neuro-hub/internal/crc"
	main_logger "DevKit-Neuro-hub/internal/logger"
	"DevKit-Neuro-hub/internal/validator"
	controller "DevKit-Neuro-hub/proto"
	"context"
	"fmt"
	"google.golang.org/protobuf/proto"
	"log"
	"net"
)

type Server struct {
	Conn      *net.UDPConn
	Addr      *net.UDPAddr
	Config    config.Config
	Logger    *main_logger.Logger
	Validator *validator.AppValidator
}

func NewServer(config config.Config, logger *main_logger.Logger, validator *validator.AppValidator) (server *Server, err error) {

	udpAddr, err := net.ResolveUDPAddr("udp", config.Addr)

	if err != nil {
		log.Fatalln(err)
	}

	conn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		log.Fatalln(err)
	}

	server = &Server{
		Conn:      conn,
		Addr:      udpAddr,
		Config:    config,
		Logger:    logger,
		Validator: validator,
	}

	return
}

func (server Server) Run(ctx context.Context) {
	var buf [1024]byte
	// Read from UDP listener in endless loop
	for {
		n, addr, err := server.Conn.ReadFromUDP(buf[:])
		if err != nil {
			log.Fatalln(err)
		}
		requestData := &controller.MessageWitchCRC{}
		err = proto.Unmarshal(buf[0:n], requestData)
		if err != nil {
			log.Fatalln(err)
		}
		// Функция проверки сообщения на подлинность
		err = utils.AuthenticationCRC(requestData)
		if err != nil {
			log.Println(err)
		} else {
			var msg proto.Message
			// В зависимости от отправляемого пакета правильно его распарсить
			if requestData.Type == controller.MessageType_channelsDataSet {
				msg = &controller.ChannelsDataSet{}
			}
			if requestData.Type == controller.MessageType_rawDataPack {
				msg = &controller.RawDataPack{}
			}
			if requestData.Type == controller.MessageType_emgMetric {
				msg = &controller.EmgMetric{}
			}
			if requestData.Type == controller.MessageType_system {
				msg = &controller.System{}
			}

			err = proto.Unmarshal(buf[0:n], msg)
			if err != nil {
				log.Fatalln(err)
			}
			fmt.Println(msg)
			// Write back the message over UPD
			server.Conn.WriteToUDP([]byte("Hello UDP Client\n"), addr)
		}
	}
}

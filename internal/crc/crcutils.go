package utils

import (
	controller "DevKit-Neuro-hub/proto"
	"errors"
	"github.com/snksoft/crc"
)

func AuthenticationCRC(msg *controller.MessageWitchCRC) error {
	var ccittCrc uint64
	if msg.Kind == controller.CRCType_crc32 {
		ccittCrc = crc.CalculateCRC(crc.CRC32, msg.Msg)
	}
	if msg.Kind == controller.CRCType_crc16 {
		ccittCrc = crc.CalculateCRC(crc.CRC16, msg.Msg)
	} else {
		return nil
	}

	if int64(ccittCrc) == msg.Crc {
		return nil
	}
	return errors.New("AuthenticationCRC calculation failed")
}

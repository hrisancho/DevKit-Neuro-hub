package utils

import (
	controller "DevKit-Neuro-plotter/proto"
	"errors"
	"github.com/snksoft/crc"
)

func AuthenticationCRC(msg *controller.MessageWitchCRC) error {
	if msg.Kind == controller.CRCType_nothizng {
		return nil
	}
	ccittCrc := crc.CalculateCRC(crc.CRC32, msg.Msg)
	if int64(ccittCrc) == msg.Crc {
		return nil
	}
	return errors.New("AuthenticationCRC calculation failed")
}

package gomksservo42c

import (
	"errors"
	"fmt"
	"log"

	"github.com/tarm/serial"
)

type Service struct {
	serial          *serial.Port
	responseChannel <-chan []uint8
}

// portname = /dev/COM45
// opens a new serial connection for the service
func CreateService(portName string, baud int) (*Service, error) {
	return CreateServiceByConfig(&serial.Config{Name: portName, Baud: baud})
}

func CreateServiceByConfig(config *serial.Config) (*Service, error) {
	s, err := serial.OpenPort(config)
	if err != nil {
		return nil, err
	}

	responseChannel := make(chan []uint8)
	started := make(chan int)

	go func() {
		buf := make([]uint8, 128)
		started <- 1

		for {
			n, err := s.Read(buf)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("sending response to chan %v\n", []byte(buf[:n]))
			responseChannel <- buf[:n]
		}
	}()

	<-started

	return &Service{
		serial:          s,
		responseChannel: responseChannel,
	}, nil
}

// returns if command was successfull, error is set if a error occured
func (s *Service) SendCommand(slave Slave, action Command, payload ...uint8) (bool, error) {
	dataToSend := AppendChecksum(append([]uint8{uint8(slave), uint8(action)}, payload...))

	_, err := s.serial.Write(dataToSend)
	if err != nil {
		return false, err
	}

	response := <-s.responseChannel

	if !CheckChecksum(response) {
		return false, errors.New("checksum from response not valid")
	}

	if len(response) > 1 && response[0] != uint8(slave) {
		return false, errors.New("response was for another slave")
	}

	if action == RunMotorSteps && len(response) > 2 {
		//another response is expected
		if response[2] == 0x01 {
			response = <-s.responseChannel
		}

		if len(response) > 2 {
			return response[2] == 0x02, nil
		}
	}

	if len(response) > 2 {
		return response[2] == 0x01, nil
	}

	return false, errors.New("cannot read response (wrong length)")
}

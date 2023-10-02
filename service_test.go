package gomksservo42c

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestServiec(t *testing.T) {
	service, err := CreateService("/dev/cu.usbserial-10", 115200)

	assert.Nil(t, err)

	// result, err := service.SendCommand(Slave0, SetZeroMode, 0x02)
	// result, err = service.SendCommand(Slave0, SetZero)

	result, err := service.SendCommand(Slave_0, Command_RunMotor, 0x0F)

	time.Sleep(time.Duration(2) * time.Second)

	result, err = service.SendCommand(Slave_0, Command_StopMotor)

	time.Sleep(time.Duration(2) * time.Second)

	result, err = service.SendCommand(Slave_0, Command_RunMotorSteps, 0xBF, 0x00, 0x00, 0xFF, 0xFF)

	assert.Nil(t, err)
	assert.False(t, result)
}

package gomksservo42c

func (s *Service) Action_RunSteps(slave Slave, clockwise bool, speed uint8, pulses uint32) (bool, error) {
	var dirspeedByte uint8

	//first bit direction, other bits from speed
	if clockwise {
		dirspeedByte = (speed & 0x7F) | 0x80
	} else {
		dirspeedByte = speed & 0x7F
	}

	result_0 := uint8((pulses & 0x000000FF))
	result_1 := uint8((pulses & 0x0000FF00) >> 8)
	result_2 := uint8((pulses & 0x00FF0000) >> 16)
	result_3 := uint8((pulses & 0xFF000000) >> 24)

	return s.SendCommand(slave, Command_RunMotorSteps, dirspeedByte, result_3, result_2, result_1, result_0)
}

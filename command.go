package gomksservo42c

// use 1 byte
type Command uint8

const (
	ReadEncoderValue            Command = 0x30
	ReadPulsesRecieved          Command = 0x33
	ReadMotorShaftAngleError    Command = 0x39
	ReadEnPinStatus             Command = 0x3A
	ReleaseLockedShaftState     Command = 0x3D
	ReadLockedShaftState        Command = 0x3E
	RunCalibration              Command = 0x80
	SetMotorType                Command = 0x81
	SetWorkMode                 Command = 0x82
	SetCurrent                  Command = 0x83
	SetSubdivision              Command = 0x84
	SetEnPinActive              Command = 0x85
	SetDir                      Command = 0x86
	SetAutomaticTurnOfScreen    Command = 0x87
	SetLockedShaftState         Command = 0x88
	SetSubdivisionInterpolation Command = 0x89
	SetBaudRate                 Command = 0x8A
	SetSlaveAddress             Command = 0x8B
	RestoreDefaults             Command = 0x3F
	SetZeroMode                 Command = 0x90
	SetZero                     Command = 0x91
	SetZeroSpeed                Command = 0x92
	SetZeroModeDir              Command = 0x93
	GoToZero                    Command = 0x94
	SetKpPosition               Command = 0xA1
	SetKiPosition               Command = 0xA2
	SetKdPosition               Command = 0xA3
	SetAcceleration             Command = 0xA4
	SetMaximumTorque            Command = 0xA5
	SetEnPinActiveUart          Command = 0xF3
	RunMotor                    Command = 0xF6
	StopMotor                   Command = 0xF7
	UpdateStatusInF6Function    Command = 0xFF
	RunMotorSteps               Command = 0xFD
)

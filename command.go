package gomksservo42c

// use 1 byte
type Command uint8

const (
	Command_ReadEncoderValue            Command = 0x30
	Command_ReadPulsesRecieved          Command = 0x33
	Command_ReadMotorShaftAngleError    Command = 0x39
	Command_ReadEnPinStatus             Command = 0x3A
	Command_ReleaseLockedShaftState     Command = 0x3D
	Command_ReadLockedShaftState        Command = 0x3E
	Command_RunCalibration              Command = 0x80
	Command_SetMotorType                Command = 0x81
	Command_SetWorkMode                 Command = 0x82
	Command_SetCurrent                  Command = 0x83
	Command_SetSubdivision              Command = 0x84
	Command_SetEnPinActive              Command = 0x85
	Command_SetDir                      Command = 0x86
	Command_SetAutomaticTurnOfScreen    Command = 0x87
	Command_SetLockedShaftState         Command = 0x88
	Command_SetSubdivisionInterpolation Command = 0x89
	Command_SetBaudRate                 Command = 0x8A
	Command_SetSlaveAddress             Command = 0x8B
	Command_RestoreDefaults             Command = 0x3F
	Command_SetZeroMode                 Command = 0x90
	Command_SetZero                     Command = 0x91
	Command_SetZeroSpeed                Command = 0x92
	Command_SetZeroModeDir              Command = 0x93
	Command_GoToZero                    Command = 0x94
	Command_SetKpPosition               Command = 0xA1
	Command_SetKiPosition               Command = 0xA2
	Command_SetKdPosition               Command = 0xA3
	Command_SetAcceleration             Command = 0xA4
	Command_SetMaximumTorque            Command = 0xA5
	Command_SetEnPinActiveUart          Command = 0xF3
	Command_RunMotor                    Command = 0xF6
	Command_StopMotor                   Command = 0xF7
	Command_UpdateStatusInF6Function    Command = 0xFF
	Command_RunMotorSteps               Command = 0xFD
)

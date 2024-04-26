package f12021

type PacketMotionData struct {
	Header        PacketHeader      `json:"m_header"`
	CarMotionData [22]CarMotionData `json:"m_carMotionData"`

	// RL, RR, FL, FR
	SuspensionPosition     [4]float32 `json:"m_suspensionPosition"`
	SuspensionVelocity     [4]float32 `json:"m_suspensionVelocity"`
	SuspensionAcceleration [4]float32 `json:"m_suspensionAcceleration"`
	WheelSpeed             [4]float32 `json:"m_wheelSpeed"`
	WheelSlip              [4]float32 `json:"m_wheelSlip"`

	LocalVelocityX float32 `json:"m_localVelocityX"`
	LocalVelocityY float32 `json:"m_localVelocityY"`
	LocalVelocityZ float32 `json:"m_localVelocityZ"`

	AngularVelocityX float32 `json:"m_angularVelocityX"`
	AngularVelocityY float32 `json:"m_angularVelocityY"`
	AngularVelocityZ float32 `json:"m_angularVelocityZ"`

	AngularAccelerationX float32 `json:"m_angularAccelerationX"`
	AngularAccelerationY float32 `json:"m_angularAccelerationY"`
	AngularAccelerationZ float32 `json:"m_angularAccelerationZ"`

	FrontWheelsAngle float32 `json:"m_frontWheelsAngle"`
}

func (p PacketMotionData) PacketHeader() Header {
	return &p.Header
}

type CarMotionData struct {
	WorldPositionX float32 `json:"m_worldPositionX"`
	WorldPositionY float32 `json:"m_worldPositionY"`
	WorldPositionZ float32 `json:"m_worldPositionZ"`

	WorldVelocityX float32 `json:"m_worldVelocityX"`
	WorldVelocityY float32 `json:"m_worldVelocityY"`
	WorldVelocityZ float32 `json:"m_worldVelocityZ"`

	WorldForwardDirX int16 `json:"m_worldForwardDirX"`
	WorldForwardDirY int16 `json:"m_worldForwardDirY"`
	WorldForwardDirZ int16 `json:"m_worldForwardDirZ"`

	WorldRightDirX int16 `json:"m_worldRightDirX"`
	WorldRightDirY int16 `json:"m_worldRightDirY"`
	WorldRightDirZ int16 `json:"m_worldRightDirZ"`

	GForceLateral      float32 `json:"m_gForceLateral"`
	GForceLongitudinal float32 `json:"m_gForceLongitudinal"`
	GForceVertical     float32 `json:"m_gForceVertical"`

	Yaw   float32 `json:"m_yaw"`
	Pitch float32 `json:"m_pitch"`
	Roll  float32 `json:"m_roll"`
}

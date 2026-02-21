package model

type FH4DataStructure struct {
	IsRaceOn int32 `json:"isRaceOn"`

	TimestampMS uint32 `json:"timestampMS"`

	EngineMaxRpm     float32 `json:"engineMaxRpm"`
	EngineIdleRpm    float32 `json:"engineIdleRpm"`
	CurrentEngineRpm float32 `json:"currentEngineRpm"`

	AccelerationX float32 `json:"accelerationX"`
	AccelerationY float32 `json:"accelerationY"`
	AccelerationZ float32 `json:"accelerationZ"`

	VelocityX float32 `json:"velocityX"`
	VelocityY float32 `json:"velocityY"`
	VelocityZ float32 `json:"velocityZ"`

	AngularVelocityX float32 `json:"angularVelocityX"`
	AngularVelocityY float32 `json:"angularVelocityY"`
	AngularVelocityZ float32 `json:"angularVelocityZ"`

	Yaw   float32 `json:"yaw"`
	Pitch float32 `json:"pitch"`
	Roll  float32 `json:"roll"`

	NormalizedSuspensionTravelFrontLeft  float32 `json:"normalizedSuspensionTravelFrontLeft"`
	NormalizedSuspensionTravelFrontRight float32 `json:"normalizedSuspensionTravelFrontRight"`
	NormalizedSuspensionTravelRearLeft   float32 `json:"normalizedSuspensionTravelRearLeft"`
	NormalizedSuspensionTravelRearRight  float32 `json:"normalizedSuspensionTravelRearRight"`

	TireSlipRatioFrontLeft  float32 `json:"tireSlipRatioFrontLeft"`
	TireSlipRatioFrontRight float32 `json:"tireSlipRatioFrontRight"`
	TireSlipRatioRearLeft   float32 `json:"tireSlipRatioRearLeft"`
	TireSlipRatioRearRight  float32 `json:"tireSlipRatioRearRight"`

	WheelRotationSpeedFrontLeft  float32 `json:"wheelRotationSpeedFrontLeft"`
	WheelRotationSpeedFrontRight float32 `json:"wheelRotationSpeedFrontRight"`
	WheelRotationSpeedRearLeft   float32 `json:"wheelRotationSpeedRearLeft"`
	WheelRotationSpeedRearRight  float32 `json:"wheelRotationSpeedRearRight"`

	WheelOnRumbleStripFrontLeft  int32 `json:"wheelOnRumbleStripFrontLeft"`
	WheelOnRumbleStripFrontRight int32 `json:"wheelOnRumbleStripFrontRight"`
	WheelOnRumbleStripRearLeft   int32 `json:"wheelOnRumbleStripRearLeft"`
	WheelOnRumbleStripRearRight  int32 `json:"wheelOnRumbleStripRearRight"`

	WheelInPuddleDepthFrontLeft  uint32 `json:"wheelInPuddleDepthFrontLeft"`
	WheelInPuddleDepthFrontRight uint32 `json:"wheelInPuddleDepthFrontRight"`
	WheelInPuddleDepthRearLeft   uint32 `json:"wheelInPuddleDepthRearLeft"`
	WheelInPuddleDepthRearRight  uint32 `json:"wheelInPuddleDepthRearRight"`

	SurfaceRumbleFrontLeft  float32 `json:"surfaceRumbleFrontLeft"`
	SurfaceRumbleFrontRight float32 `json:"surfaceRumbleFrontRight"`
	SurfaceRumbleRearLeft   float32 `json:"surfaceRumbleRearLeft"`
	SurfaceRumbleRearRight  float32 `json:"surfaceRumbleRearRight"`

	TireSlipAngleFrontLeft  float32 `json:"tireSlipAngleFrontLeft"`
	TireSlipAngleFrontRight float32 `json:"tireSlipAngleFrontRight"`
	TireSlipAngleRearLeft   float32 `json:"tireSlipAngleRearLeft"`
	TireSlipAngleRearRight  float32 `json:"tireSlipAngleRearRight"`

	TireCombinedSlipFrontLeft  float32 `json:"tireCombinedSlipFrontLeft"`
	TireCombinedSlipFrontRight float32 `json:"tireCombinedSlipFrontRight"`
	TireCombinedSlipRearLeft   float32 `json:"tireCombinedSlipRearLeft"`
	TireCombinedSlipRearRight  float32 `json:"tireCombinedSlipRearRight"`

	SuspensionTravelMetersFrontLeft  float32 `json:"suspensionTravelMetersFrontLeft"`
	SuspensionTravelMetersFrontRight float32 `json:"suspensionTravelMetersFrontRight"`
	SuspensionTravelMetersRearLeft   float32 `json:"suspensionTravelMetersRearLeft"`
	SuspensionTravelMetersRearRight  float32 `json:"suspensionTravelMetersRearRight"`

	CarOrdinal int32 `json:"carOrdinal"`

	CarClass int32 `json:"carClass"`

	CarPerformanceIndex int32 `json:"carPerformanceIndex"`

	DrivetrainType int32 `json:"drivetrainType"`

	NumCylinders int32 `json:"numCylinders"`

	Unknown1 int32 `json:"unknown1"`
	Unknown2 int32 `json:"unknown2"`
	Unknown3 int32 `json:"unknown3"`

	PositionX float32 `json:"positionX"`
	PositionY float32 `json:"positionY"`
	PositionZ float32 `json:"positionZ"`

	Speed  float32 `json:"speed"`
	Power  float32 `json:"power"`
	Torque float32 `json:"torque"`

	TireTempFrontLeft  float32 `json:"tireTempFrontLeft"`
	TireTempFrontRight float32 `json:"tireTempFrontRight"`
	TireTempRearLeft   float32 `json:"tireTempRearLeft"`
	TireTempRearRight  float32 `json:"tireTempRearRight"`

	Boost            float32 `json:"boost"`
	Fuel             float32 `json:"fuel"`
	DistanceTraveled float32 `json:"distanceTraveled"`

	BestLap         float32 `json:"bestLap"`
	LastLap         float32 `json:"lastLap"`
	CurrentLap      float32 `json:"currentLap"`
	CurrentRaceTime float32 `json:"currentRaceTime"`
	LapNumber       uint16  `json:"lapNumber"`
	RacePosition    uint8   `json:"racePosition"`

	Accel     uint8 `json:"accel"`
	Brake     uint8 `json:"brake"`
	Clutch    uint8 `json:"clutch"`
	HandBrake uint8 `json:"handBrake"`
	Gear      uint8 `json:"gear"`
	Steer     int8  `json:"steer"`

	NormalizedDrivingLine       int8 `json:"normalizedDrivingLine"`
	NormalizedAIBrakeDifference int8 `json:"normalizedAIBrakeDifference"`
	Unknown4                    int8 `json:"unknown4"`
}

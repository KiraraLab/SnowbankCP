package repository

type ChannelInfo struct {
	Name        string
	Description string
	Enabled     bool
	Ratio       float64
	SensorRatio float64
}

type DeviceConfig struct {
	DeviceID    string
	Name        string
	Description string
	Enabled     bool
	Channels    map[int8]ChannelInfo
}

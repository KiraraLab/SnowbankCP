package datachannel

type CollectorInfo struct {
	Name        string
	Description string
	ID          string
}

type Collector interface {
	Channels() []DataChannel
}

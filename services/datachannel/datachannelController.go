package datachannel

type DataChannelController struct {
	datachannel DataChannel
	collector   Collector
	storer      Storer
}

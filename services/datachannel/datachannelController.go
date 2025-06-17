package datachannel

import (
	"github.com/kiraralab/snowbankcp/services/collectors"
	"github.com/kiraralab/snowbankcp/services/storers"
	// "collectors"
	// "storers"
)

type DataChannelController struct {
	datachannel DataChannel
	collector   collectors.Collector
	storer      storers.Storer
}

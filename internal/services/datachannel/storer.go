package datachannel

type StorerInfo struct {
	Name        string
	Description string
	ID          string
}

type Storer interface {
	// Store stores the data to the storer.
	Store(Record) error
}

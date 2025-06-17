package datachannel

type DataChannel struct {
	Name        string
	Description string
	ID          string
	Topics      []string
}

func (dataChannel *DataChannel) Store(record Record) error {
	return nil
}
func (dataChannel *DataChannel) Collect() (Record, error) {
	return nil, nil
}

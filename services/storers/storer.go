package storers

type Storer interface {
	// Store stores the data to the storer.
	Store(data interface{}) error
}

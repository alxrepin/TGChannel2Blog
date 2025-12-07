package domain

type Bus interface {
	Dispatch(event EventType, data []byte) error
}

package kafka

type Consumer interface {
	Process(record []byte) error
}

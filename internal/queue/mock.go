package queue

type MockQueue struct {
	q []*AppQueueDto
}

func (mq *MockQueue) Publish(msg []byte) error {
	dto := new(AppQueueDto)
	dto.Unmarshal(msg)

	mq.q = append(mq.q, dto)

	return nil
}

func (mq *MockQueue) Consume(c chan<- AppQueueDto) error {
	return nil
}

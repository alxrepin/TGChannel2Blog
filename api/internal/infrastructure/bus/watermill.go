package bus

import (
	"app/internal/domain"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
)

type WatermillBus struct {
	publisher message.Publisher
}

func NewWatermillBus(publisher message.Publisher) *WatermillBus {
	return &WatermillBus{publisher: publisher}
}

func (w *WatermillBus) Dispatch(event domain.EventType, data []byte) error {
	msg := message.NewMessage(watermill.NewUUID(), data)

	return w.publisher.Publish(string(event), msg)
}

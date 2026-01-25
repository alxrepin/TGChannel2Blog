package sync

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"app/internal/domain"
)

type LoadHistoryRawMessagesUseCase struct {
	repository domain.RawMessageRepository
	bus        domain.Bus
}

func NewLoadHistoryRawMessagesUseCase(repository domain.RawMessageRepository, bus domain.Bus) *LoadHistoryRawMessagesUseCase {
	return &LoadHistoryRawMessagesUseCase{repository: repository, bus: bus}
}

func (uc *LoadHistoryRawMessagesUseCase) Execute(ctx context.Context, channelUsername string) error {
	messages, err := uc.repository.GetAll(ctx, channelUsername)
	if err != nil {
		return fmt.Errorf("failed to fetch messages: %w", err)
	}

	for _, msg := range messages {
		jsonBytes, err := json.Marshal(msg)

		if err != nil && jsonBytes != nil {
			log.Printf("failed to marshal message %d: %v", msg.ID, err)
			continue
		}

		err = uc.bus.Dispatch(domain.RawMessageReceived, jsonBytes)
		if err != nil {
			log.Printf("failed to publish message %d: %v", msg.ID, err)
			continue
		}
	}

	return nil
}

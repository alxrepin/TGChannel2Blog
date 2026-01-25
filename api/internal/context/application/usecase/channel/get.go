package channel

import (
	"app/internal/context/domain"
	"context"
	"fmt"
)

type GetChannelUseCase struct {
	repository domain.ChannelRepository
}

type ChannelResponse struct {
	ID            int64   `json:"id"`
	Name          string  `json:"name"`
	Title         string  `json:"title"`
	Description   *string `json:"description,omitempty"`
	Avatar        *string `json:"avatar,omitempty"`
	Subscriptions int64   `json:"subscriptions"`
}

func NewGetChannelUseCase(channelRepository domain.ChannelRepository) *GetChannelUseCase {
	return &GetChannelUseCase{
		repository: channelRepository,
	}
}

func (uc *GetChannelUseCase) Execute(ctx context.Context) (*ChannelResponse, error) {
	channel, err := uc.repository.Get(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get channel: %w", err)
	}

	response := &ChannelResponse{
		ID:            channel.ID,
		Name:          channel.Name,
		Title:         channel.Title,
		Description:   channel.Description,
		Avatar:        channel.Avatar,
		Subscriptions: channel.Subscriptions,
	}

	return response, nil
}

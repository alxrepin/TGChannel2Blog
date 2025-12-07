package telegram

import (
	"context"
	"fmt"

	"github.com/gotd/td/session"
	"github.com/gotd/td/telegram"
	"github.com/gotd/td/telegram/auth"
	"github.com/gotd/td/tg"
)

type Client struct {
	apiID   int
	apiHash string
	phone   string
}

func NewClient(apiID int, apiHash, phone string) *Client {
	return &Client{
		apiID:   apiID,
		apiHash: apiHash,
		phone:   phone,
	}
}

func (c *Client) Run(ctx context.Context, fn func(ctx context.Context, client *telegram.Client) error) error {
	client := telegram.NewClient(c.apiID, c.apiHash, telegram.Options{
		SessionStorage: &session.FileStorage{Path: "./var/telegram/session.json"},
	})

	flow := auth.NewFlow(
		auth.CodeOnly(c.phone, Authenticator{}),
		auth.SendCodeOptions{},
	)

	return client.Run(ctx, func(ctx context.Context) error {
		if err := client.Auth().IfNecessary(ctx, flow); err != nil {
			return fmt.Errorf("auth failed: %w", err)
		}

		return fn(ctx, client)
	})
}

func (c *Client) FetchMessages(ctx context.Context, client *telegram.Client, channelUsername string) ([]tg.MessageClass, error) {
	resolve, err := client.API().ContactsResolveUsername(
		ctx,
		&tg.ContactsResolveUsernameRequest{
			Username: channelUsername,
		},
	)

	if err != nil {
		return nil, fmt.Errorf("resolve username failed: %w", err)
	}

	if len(resolve.Chats) == 0 {
		return nil, fmt.Errorf("channel not found")
	}

	channel, ok := resolve.Chats[0].(*tg.Channel)
	if !ok {
		return nil, fmt.Errorf("not a channel")
	}

	var allMessages []tg.MessageClass
	offsetID := 0

	for {
		req := &tg.MessagesGetHistoryRequest{
			Peer: &tg.InputPeerChannel{
				ChannelID:  channel.ID,
				AccessHash: channel.AccessHash,
			},
			OffsetID: offsetID,
			Limit:    100,
		}

		resp, err := client.API().MessagesGetHistory(ctx, req)
		if err != nil {
			return nil, fmt.Errorf("get history failed: %w", err)
		}

		var msgs []tg.MessageClass

		switch m := resp.(type) {
		case *tg.MessagesMessages:
			msgs = m.Messages
		case *tg.MessagesChannelMessages:
			msgs = m.Messages
		default:
			return nil, fmt.Errorf("unexpected history type")
		}

		if len(msgs) == 0 {
			break
		}

		allMessages = append(allMessages, msgs...)

		if lastMsg, ok := msgs[len(msgs)-1].(*tg.Message); ok {
			offsetID = lastMsg.ID
		} else {
			break
		}
	}

	return allMessages, nil
}

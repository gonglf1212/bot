package botsdk

import (
	"context"

	"github.com/bot/dto"
	"github.com/bot/log"
	"go.uber.org/zap"
)

func (c *Client) PostMessage(channelID string, msg *dto.MessageToCreate) (*dto.Message, error) {
	resp, err := c.restyClient.R().SetContext(context.Background()).
		SetResult(dto.Message{}).
		SetPathParam("channel_id", channelID).
		SetBody(msg).
		Post(messagesURI)
	if err != nil {
		log.Error("--post message err", zap.Error(err))
		return nil, err
	}
	return resp.Result().(*dto.Message), nil
}

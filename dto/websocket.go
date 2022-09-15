/*
 * @Author: gonglf
 * @Date: 2022-09-15 09:35:25
 * @LastEditors: gonglf
 * @LastEditTime: 2022-09-15 10:33:14
 * @Description:
 *
 */
package dto

import (
	"github.com/bot/internal/bot/token"
)

// WebsocketAP wss 接入点信息
type WebsocketAP struct {
	URL               string            `json:"url"`
	Shards            uint32            `json:"shards"`
	SessionStartLimit SessionStartLimit `json:"session_start_limit"`
}

// SessionStartLimit 链接频控信息
type SessionStartLimit struct {
	Total          uint32 `json:"total"`
	Remaining      uint32 `json:"remaining"`
	ResetAfter     uint32 `json:"reset_after"`
	MaxConcurrency uint32 `json:"max_concurrency"`
}

// Session 连接的 session 结构，包括链接的所有必要字段
type Session struct {
	ID      string
	URL     string
	Token   token.Token
	Intent  Intent
	LastSeq uint32
	// Shards  ShardConfig
}

// WSPayload websocket 消息结构
type WSPayload struct {
	// WSPayloadBase
	Data       interface{} `json:"d,omitempty"`
	RawMessage []byte      `json:"-"` // 原始的 message 数据
}

// Intent 类型
type Intent int

const (
	IntentGuilds Intent = 1 << iota

	// IntentGuildMembers 包含
	// - GUILD_MEMBER_ADD
	// - GUILD_MEMBER_UPDATE
	// - GUILD_MEMBER_REMOVE
	IntentGuildMembers

	IntentGuildBans
	IntentGuildEmojis
	IntentGuildIntegrations
	IntentGuildWebhooks
	IntentGuildInvites
	IntentGuildVoiceStates
	IntentGuildPresences
	IntentGuildMessages
)

// WSIdentityData 鉴权数据
type WSIdentityData struct {
	Token      string   `json:"token"`
	Intents    Intent   `json:"intents"`
	Shard      []uint32 `json:"shard"` // array of two integers (shard_id, num_shards)
	Properties struct {
		Os      string `json:"$os,omitempty"`
		Browser string `json:"$browser,omitempty"`
		Device  string `json:"$device,omitempty"`
	} `json:"properties,omitempty"`
}

/*
 * @Author: gonglf
 * @Date: 2022-09-15 12:25:01
 * @LastEditors: gonglf
 * @LastEditTime: 2022-09-15 18:26:24
 * @Description:
 *
 */
package bot

import (
	"strings"

	"github.com/bot/dto"
	"github.com/bot/dto/message"
	"github.com/bot/internal/bot/botsdk"
	"github.com/bot/internal/bot/event"
	"github.com/bot/internal/bot/token"
	"github.com/bot/internal/bot/websocket"
	"github.com/bot/internal/bot/websocket/client"
	"github.com/gocpp/log"
	"go.uber.org/zap"
)

type Server struct {
	apiClient *botsdk.Client
}

func NewServer() {

}

func (s *Server) Connect(token *token.Token) {
	s.apiClient = botsdk.NewBotSdk(token)
	apInfo, err := s.apiClient.Gateway()

	if err != nil {
		panic("connet error")
		return
	}
	var intents = s.registerHandlers()
	session := dto.Session{
		URL:     apInfo.URL,
		Token:   *token,
		Intent:  intents,
		LastSeq: 0,
		// Shards: dto.ShardConfig{
		// 	ShardID:    i,
		// 	ShardCount: apInfo.Shards,
		// },
	}

	// processor = Processor{api: api}
	client.Setup()
	wsClient := websocket.ClientImpl.New(session)
	if err := wsClient.Connect(); err != nil {
		log.Error("Connect error", zap.Error(err))
		// s.sessionChan <- session // 连接失败，丢回去队列排队重连
		return
	}
	// 初次鉴权
	err = wsClient.Identify()
	if err != nil {
		log.Error("[ws/session] Identify/Resume err ", zap.Error(err))
		return
	}
	if err := wsClient.Listening(); err != nil { //会阻塞
		log.Error("[ws/session] Listening err ", zap.Error(err))
	}
}

func (s *Server) registerHandlers() (intents dto.Intent) {
	intents = websocket.RegisterHandlers(
		// at 机器人事件，目前是在这个事件处理中有逻辑，会回消息，其他的回调处理都只把数据打印出来，不做任何处理
		// event.DefaultHandlers.ATMessage, //用斜杠指令的方式
		s.ATMessageEventHandler(),
		// // 如果想要捕获到连接成功的事件，可以实现这个回调
		// ReadyHandler(),
		event.DefaultHandlers.Ready,
		// // 连接关闭回调
		// ErrorNotifyHandler(),
		event.DefaultHandlers.ErrorNotify,
		// // 频道事件
		// GuildEventHandler(),
		// event.DefaultHandlers.Guild,
		// // 成员事件
		// MemberEventHandler(),
		// event.DefaultHandlers.GuildMember,
		// // 子频道事件
		// ChannelEventHandler(),
		// event.DefaultHandlers.Channel,
		// // 私信，目前只有私域才能够收到这个，如果你的机器人不是私域机器人，会导致连接报错，那么启动 example 就需要注释掉这个回调
		// DirectMessageHandler(),
		event.DefaultHandlers.DirectMessage,
		// // 频道消息，只有私域才能够收到这个，如果你的机器人不是私域机器人，会导致连接报错，那么启动 example 就需要注释掉这个回调
		// CreateMessageHandler(),
		// event.DefaultHandlers.Message, //当前的聊天内容
		// // 互动事件
		// InteractionHandler(),
		event.DefaultHandlers.Interaction,
		// // 发帖事件
		// ThreadEventHandler(),
		event.DefaultHandlers.Thread,
	)
	return
}

func (s *Server) ATMessageEventHandler() event.ATMessageEventHandler {
	return func(event *dto.WSPayload, data *dto.WSATMessageData) error {
		input := strings.ToLower(message.ETLInput(data.Content))
		cmd := message.ParseCommand(input) //指令
		toCreate := &dto.MessageToCreate{
			Content: "打卡成功",
			MessageReference: &dto.MessageReference{
				// 引用这条消息
				MessageID:             data.ID,
				IgnoreGetMessageError: true,
			},
			MsgID: data.ID,
		}

		switch cmd.Cmd {

		case "/打卡":
			s.apiClient.PostMessage(data.ChannelID, toCreate)
		}
		return nil
	}
}

/*
 * @Author: gonglf
 * @Date: 2022-09-15 09:35:25
 * @LastEditors: gonglf
 * @LastEditTime: 2022-09-15 10:22:30
 * @Description:
 *
 */
package websocket

// // type messageChan chan *dto.WSPayload
// type closeErrorChan chan error

// // Client websocket 连接客户端
// type Client struct {
// 	version int
// 	conn    *wss.Conn
// 	// messageQueue    messageChan
// 	session *dto.Session
// 	// user            *dto.WSUser
// 	// closeChan       closeErrorChan
// 	heartBeatTicker *time.Ticker // 用于维持定时心跳
// }

// // New 新建一个连接对象
// func (c *Client) New(session dto.Session) WebSocket {
// 	return &Client{
// 		messageQueue:    make(messageChan, DefaultQueueSize),
// 		session:         &session,
// 		closeChan:       make(closeErrorChan, 10),
// 		heartBeatTicker: time.NewTicker(60 * time.Second), // 先给一个默认 ticker，在收到 hello 包之后，会 reset
// 	}
// }

// // Connect 连接到 websocket
// func (c *Client) Connect() error {
// 	if c.session.URL == "" {
// 		return errors.New("err url")
// 	}

// 	var err error
// 	c.conn, _, err = wss.DefaultDialer.Dial(c.session.URL, nil)
// 	if err != nil {
// 		log.Errorf("%s, connect err: %v", c.session, err)
// 		return err
// 	}
// 	log.Infof("%s, url %s, connected", c.session, c.session.URL)

// 	return nil
// }

// // Identify 对一个连接进行鉴权，并声明监听的 shard 信息
// func (c *Client) Identify() error {
// 	// 避免传错 intent
// 	if c.session.Intent == 0 {
// 		c.session.Intent = dto.IntentGuilds
// 	}
// 	payload := &dto.WSPayload{
// 		Data: &dto.WSIdentityData{
// 			Token:   c.session.Token.GetString(),
// 			Intents: c.session.Intent,
// 			Shard: []uint32{
// 				c.session.Shards.ShardID,
// 				c.session.Shards.ShardCount,
// 			},
// 		},
// 	}
// 	payload.OPCode = dto.WSIdentity
// 	return c.Write(payload)
// }

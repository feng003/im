package client

import (
	"context"
	"time"
	"im/common/libnet"
	"im/rpc_im/client/imservice"

	"github.com/golang/protobuf/proto"
	"github.com/zeromicro/go-zero/core/logx"
)

type Client struct {
	Session   *libnet.Session
	Manager   *libnet.Manager
	ImRpc     imservice.ImService
	heartbeat chan *libnet.Message
}

func NewClient(manager *libnet.Manager, session *libnet.Session, imrpc imservice.ImService) *Client {
	return &Client{
		Session:   session,
		Manager:   manager,
		ImRpc:     imrpc,
		heartbeat: make(chan *libnet.Message),
	}
}

func (c *Client) Login(msg *libnet.Message) error {
	loginReq, err := makeLoginMessage(msg)
	if err != nil {
		return err
	}

	c.Session.SetToken(loginReq.Token)
	c.Manager.AddSession(c.Session)

	_, err = c.ImRpc.Login(context.Background(), &imservice.LoginRequest{
		Token:         loginReq.Token,
		Authorization: loginReq.Authorization,
		SessionId:     c.Session.Session().String(),
	})
	if err != nil {
		msg.Status = 1
		msg.Body = []byte(err.Error())
		err = c.Send(*msg)
		if err != nil {
			logx.Errorf("[Login] client.Send error: %v", err)
		}
		return err
	}

	msg.Status = 0
	msg.Body = []byte("login success")
	err = c.Send(*msg)
	if err != nil {
		logx.Errorf("[Login] client.Send error: %v", err)
	}

	return err
}

func (c *Client) Receive() (*libnet.Message, error) {
	return c.Session.Receive()
}

func (c *Client) Send(msg libnet.Message) error {
	return c.Session.Send(msg)
}

func (c *Client) Close() error {
	return c.Session.Close()
}

func (c *Client) HandlePackage(msg *libnet.Message) error {
	// 消息转发
	req := makePostMessage(c.Session.Session().String(), msg)
	logx.Infof("[HandlePackage] session: %s client.PostMessage: %v", c.Session.Session().String(), req)
	if req == nil {
		return nil
	}
	_, err := c.ImRpc.PostMessage(context.Background(), req)
	if err != nil {
		logx.Errorf("[HandlePackage] client.PostMessage error: %v", err)
	}

	return err
}

const heartBeatTimeout = time.Second * 60

func (c *Client) HeartBeat() error {
	timer := time.NewTimer(heartBeatTimeout)
	for {
		select {
		case heaetbeat := <-c.heartbeat:
			c.Session.SetReadDeadline(time.Now().Add(heartBeatTimeout * 5))
			c.Send(*heaetbeat)
			break
		case <-timer.C:
		}
	}
}

func makeLoginMessage(msg *libnet.Message) (*imservice.LoginRequest, error) {
	// 登录功能还没做
	// 这里临时处理，先把PostMsg中的Msg转换成LoginRequest中的Token和Authorization用于登录处理
	var postMsg imservice.PostMsg
	err := proto.Unmarshal(msg.Body, &postMsg)
	if err != nil {
		return nil, err
	}
	loginReq := imservice.LoginRequest{
		Token:         postMsg.Msg,
		Authorization: postMsg.Msg,
	}

	return &loginReq, nil
}

func makePostMessage(sessionId string, msg *libnet.Message) *imservice.PostMsg {
	var postMessageReq imservice.PostMsg
	err := proto.Unmarshal(msg.Body, &postMessageReq)
	if err != nil {
		logx.Errorf("[makePostMessage] proto.Unmarshal msg: %v error: %v", msg, err)
		return nil
	}
	postMessageReq.Version = uint32(msg.Version)
	postMessageReq.Status = uint32(msg.Status)
	postMessageReq.ServiceId = uint32(msg.ServiceId)
	postMessageReq.Cmd = uint32(msg.Cmd)
	postMessageReq.Seq = msg.Seq
	postMessageReq.SessionId = sessionId

	return &postMessageReq
}

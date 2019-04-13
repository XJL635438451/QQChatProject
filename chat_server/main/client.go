package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"MyGitHubProject/Wechat/QQChatProject/proto"
	"MyGitHubProject/Wechat/QQChatProject/common"
)

type Client struct {
	conn   net.Conn
	userId int
	buf    [8192]byte
}

func (p *Client) Process() (err error) {
	for {
		var msg proto.Message
		msg, err = common.RecvMsg(p.conn)
		if err != nil {
			clientMgr.DelClient(p.userId)
			//TODO:通知所有在线用户，该用户已经下线
			return err
		}

		err = p.processMsg(msg)
		if err != nil {
			fmt.Println("process msg failed, err:", err)
			continue
			//return
		}
	}
}

func (p *Client) processMsg(msg proto.Message) (err error) {
	switch msg.Cmd {
	case proto.UserLoginFlag:
		err = p.login(msg)
	case proto.UserRegisterFlag:
		err = p.register(msg)
	case proto.UserSendMessageFlag:
		err = p.proccessUserSendMessage(msg)
	default:
		err = errors.New("unsupport message")
		return
	}
	return
}

func (p *Client) SendMessageToUser(userId int, text string) {
	var respMsg proto.Message
	respMsg.Cmd = proto.UserRecvMessageFlag

	var recvMsg proto.UserRecvMessageReqData
	recvMsg.UserId = userId
	recvMsg.Data = text

	data, err := json.Marshal(recvMsg)
	if err != nil {
		fmt.Println("marshal failed, ", err)
		return
	}

	respMsg.Data = string(data)
	err = common.SendMsg(p.conn, respMsg)
	if err != nil {
		return
	}
}

func (p *Client) proccessUserSendMessage(msg proto.Message) (err error) {
	var userReq proto.UserSendMessageReqData
	err = json.Unmarshal([]byte(msg.Data), &userReq)
	if err != nil {
		fmt.Println("unmarshal failed, err:", err)
		return
	}

	users := clientMgr.GetAllUsers()
	for id, client := range users {
		if id == userReq.UserId {
			continue
		}

		client.SendMessageToUser(userReq.UserId, userReq.Data)
	}
	return
}

func (p *Client) loginResp(err error) {
	var respMsg proto.Message
	respMsg.Cmd = proto.UserLoginResFlag

	var loginRes proto.LoginResData
	loginRes.Code = 200

	userMap := clientMgr.GetAllUsers()
	for userId, _ := range userMap {
		loginRes.User = append(loginRes.User, userId)
	}

	if err != nil {
		loginRes.Code = 500
		loginRes.Error = fmt.Sprintf("%v", err)
	}

	data, err := json.Marshal(loginRes)
	if err != nil {
		fmt.Println("marshal failed, ", err)
		return
	}

	respMsg.Data = string(data)
	err = common.SendMsg(p.conn, respMsg)
	if err != nil {
		return
	}
}

func (p *Client) login(msg proto.Message) (err error) {
	defer func() {
		p.loginResp(err)
	}()

	fmt.Printf("recv user login request, data:%v", msg)
	var cmd proto.LoginData
	err = json.Unmarshal([]byte(msg.Data), &cmd)
	if err != nil {
		fmt.Println("unmarshal failed, err:", err)
		return
	}

	_, err = mgr.Login(cmd.Id, cmd.Passwd)
	if err != nil {
		return
	}

	clientMgr.AddClient(cmd.Id, p)
	p.userId = cmd.Id

	p.NotifyOthersUserOnline(cmd.Id)
	return
}

func (p *Client) NotifyOthersUserOnline(userId int) {
	users := clientMgr.GetAllUsers()
	for id, client := range users {
		if id == userId {
			continue
		}

		client.NotifyUserOnline(userId)
	}
}

func (p *Client) NotifyUserOnline(userId int) {

	var respMsg proto.Message
	respMsg.Cmd = proto.UserStatusNotifyResFlag

	var noitfyRes proto.UserStatusNotifyData
	noitfyRes.UserId = userId
	noitfyRes.Status = common.UserOnline

	data, err := json.Marshal(noitfyRes)
	if err != nil {
		fmt.Println("marshal failed, ", err)
		return
	}

	respMsg.Data = string(data)
	err = common.SendMsg(p.conn, respMsg)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func (p *Client) register(msg proto.Message) (err error) {
	var cmd proto.RegisterData
	err = json.Unmarshal([]byte(msg.Data), &cmd)
	if err != nil {
		return
	}

	err = mgr.Register(&cmd.User)
	if err != nil {
		return
	}
    //notify client that register success
	p.loginResp(nil) 

	return
}

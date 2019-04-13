package main

import (
	"net"
	"fmt"
	"encoding/json"
	"MyGitHubProject/Wechat/QQChatProject/proto"
	"MyGitHubProject/Wechat/QQChatProject/common"
)

func register(conn net.Conn, userId int, passwd string) (err error) {
	var msg proto.Message
	msg.Cmd = proto.UserRegisterFlag

	var registerCmd proto.RegisterData
	registerCmd.User.UserId = userId
	registerCmd.User.Passwd = passwd
	err = userRegisterMessage(&registerCmd)
	if err != nil {
		return
	}
	data, err := json.Marshal(registerCmd)
	if err != nil {
		return
	}

	msg.Data = string(data)
	err = common.SendMsg(conn, msg)
	if err != nil {
		return
	}

	msg, err = common.RecvMsg(conn)
	if err != nil {
		return
	}

	var registerRes proto.LoginResData
	err = json.Unmarshal([]byte(msg.Data), &registerRes)
    if err != nil {
        err = fmt.Errorf("Failed to unmarshal data[%v].", msg.Data)
        return
    }

    //need to verify register success
	fmt.Println(msg)
	return
}

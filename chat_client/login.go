package main

import (
	//"os"
	"fmt"
	"net"
	"encoding/json"
	"MyGitHubProject/Wechat/QQChatProject/proto"
	"MyGitHubProject/Wechat/QQChatProject/common"
)

func login(conn net.Conn, userId int, passwd string) (err error) {
	var msg proto.Message
	msg.Cmd = proto.UserLoginFlag

	var loginData proto.LoginData
	loginData.Id = userId
	loginData.Passwd = passwd

	data, err := json.Marshal(loginData)
	if err != nil {
		return
	}
	msg.Data = string(data)
    //send login data to server
	err = common.SendMsg(conn, msg)
	if err != nil {
		return
	}
	//receive login result from server
	msg, err = common.RecvMsg(conn)
	if err != nil {
		return
	}

	var loginResp proto.LoginResData
	err = json.Unmarshal([]byte(msg.Data), &loginResp)
	if err != nil {
		err = fmt.Errorf("Unmarshal data[%v] failed, Error: %v", msg.Data, err)
		return
	}
	if loginResp.Code == common.UserNotRegisterCode {
		fmt.Println("User not register, start register.")
		err = register(conn, userId, passwd)
		if err != nil {
			return
		}
        fmt.Println("Register success, now start to chat with your friends.")
		//os.Exit(0)
	}
    //need to  optimize, when user register success, return the online user.
	//fmt.Println("online user list:")
	for _, v := range loginResp.User {
		if v == userId {
			continue
		}
		fmt.Println("user logined:", v)
		user := &proto.User{UserId: v}
		onlineUserMap[user.UserId] = user
	}
	return
}

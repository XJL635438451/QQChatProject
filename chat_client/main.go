package main

import (
	"fmt"
	"net"
	"MyGitHubProject/Wechat/QQChatProject/proto"
)

var userId int
var passwd string

var msgChan chan proto.UserRecvMessageReqData

func init() {
	//Up to chat with 1000 friends
	msgChan = make(chan proto.UserRecvMessageReqData, 1000)
}

func main() {
	err := initUserInput() 
	if err != nil {
		fmt.Println(err)
		return
	}

	conn, err := net.Dial("tcp", "localhost:10000")
	if err != nil {
		fmt.Println("Error dialing", err.Error())
		return
	}

	err = login(conn, userId, passwd)
	if err != nil {
		fmt.Println("login failed, err:", err)
		return
	}

	go processServerMessage(conn)

	for {
		logic(conn)
	}
}

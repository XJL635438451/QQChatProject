package main

import (
	"fmt"
	"net"
	"os"
	"encoding/json"
	"MyGitHubProject/Wechat/QQChatProject/proto"
	"MyGitHubProject/Wechat/QQChatProject/common"
)

func processServerMessage(conn net.Conn) {
	for {
		msg, err := common.RecvMsg(conn)
		if err != nil {
			fmt.Println("read err:", err)
			os.Exit(0)
		}

		switch msg.Cmd {
		case proto.UserStatusNotifyResFlag:
			updateUserStatus(msg)
		case proto.UserRecvMessageFlag:
			recvMessageFromServer(msg)
		}
	}
}

func recvMessageFromServer(msg proto.Message) {
	var recvMsg proto.UserRecvMessageReqData
	err := json.Unmarshal([]byte(msg.Data), &recvMsg)
	if err != nil {
		fmt.Println("unmarshal failed, err:", err)
		return
	}
	fmt.Printf("%d:%s\n", recvMsg.UserId, recvMsg.Data)
	//msgChan <- recvMsg
}

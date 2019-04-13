package main

import (
	"fmt"
	"net"
	"os"
	"encoding/json"
	"strconv"
	"MyGitHubProject/Wechat/QQChatProject/proto"
	"MyGitHubProject/Wechat/QQChatProject/common"
)

func sendTextMessage(conn net.Conn, text string) (err error) {
	var msg proto.Message
	msg.Cmd = proto.UserSendMessageFlag

	var sendReq proto.UserSendMessageReqData
	sendReq.Data = text
	sendReq.UserId = userId

	data, err := json.Marshal(sendReq)
	if err != nil {
		return
	}

	msg.Data = string(data)
	
	err = common.SendMsg(conn, msg)
	if err != nil {
		return
	}

	return
}

func enterTalk(conn net.Conn) {
	for {
		//var destUserId int
		input, _ := common.GetInput("please input text(q exit chat): ")
		err := sendTextMessage(conn, input)
		if err != nil {
			fmt.Println(err)
		}
		if input == "q" {
			break
		}
	}
}

func listUnReadMsg() {
	select {
	case msg := <-msgChan:
		fmt.Printf("%d:%s\n", msg.UserId, msg.Data)
	default:
		return
	}
}

func enterMenu(conn net.Conn) {
	fmt.Println("1. list online user")
	fmt.Println("2. talk")
	fmt.Println("3. list message")
	fmt.Println("4. exit")

	var sel int
	for {
		input, err := common.GetInput("Sel: ")
		if err != nil {
			fmt.Println("Input error.")
            continue
		}
		sel, err = strconv.Atoi(input)
		if err == nil {
			break
		} else {
			fmt.Println("Please input your choice(1/2/3/4)")
		}
	}
	
	switch sel {
	case 1:
		outputUserOnline()
	case 2:
		enterTalk(conn)
	case 3:
		listUnReadMsg()
		return
	case 4:
		os.Exit(0)
	}
}

func logic(conn net.Conn) {
	enterMenu(conn)
}

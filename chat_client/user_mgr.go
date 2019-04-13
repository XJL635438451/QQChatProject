package main

import (
	"fmt"
	"encoding/json"
	"MyGitHubProject/Wechat/QQChatProject/proto"
)

var onlineUserMap map[int]*proto.User = make(map[int]*proto.User, 16)

func outputUserOnline() {
	fmt.Println("Online User List:")
	for id, _ := range onlineUserMap {
		if id == userId {
			continue
		}
		fmt.Println("user:", id)
	}
}

func updateUserStatus(msg proto.Message) {
	var userStatus proto.UserStatusNotifyData
	err := json.Unmarshal([]byte(msg.Data), &userStatus)
	if err != nil {
		fmt.Println("unmarshal failed, err:", err)
		return
	}

	user, ok := onlineUserMap[userStatus.UserId]
	if !ok {
		user = &proto.User{}
		user.UserId = userStatus.UserId
	}

	user.Status = userStatus.Status
	onlineUserMap[user.UserId] = user

	outputUserOnline()
}

package main

import"MyGitHubProject/Wechat/QQChatProject/chat_server/model"

var (
	mgr *model.UserMgr
)

func initUserMgr() {
	mgr = model.NewUserMgr(pool)
}

package proto

const (
	UserLoginFlag           = "user_login"
	UserLoginResFlag        = "user_login_res"
	UserRegisterFlag        = "user_register"
	UserStatusNotifyResFlag = "user_status_notify"
	UserSendMessageFlag     = "user_send_message"
	UserRecvMessageFlag     = "user_recv_message"
)

type Message struct {
	Cmd  string `json:"cmd"`
	Data string `json:"data"`
}

type LoginData struct {
	Id     int    `json:"user_id"`
	Passwd string `json:"passwd"`
}

type LoginResData struct {
	Code  int    `json:"code"`
	User  []int  `json:"users"`
	Error string `json:"error"`
}

type RegisterData struct {
	User User `json:"user"`
}

type UserStatusNotifyData struct {
	UserId int `json:"user_id"`
	Status int `json:"user_status"`
}

type UserSendMessageReqData struct {
	UserId int    `json:"user_id"`
	Data   string `json:"data"`
}

type UserRecvMessageReqData struct {
	UserId int    `json:"user_id"`
	Data   string `json:"data"`
}

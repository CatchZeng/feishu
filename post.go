package feishu

type PostMessage struct {
	MsgType MsgType     `json:"msgtype"`
	OpenID  string      `json:"open_id"`
	RoorID  string      `json:"root_id"`
	ChatID  string      `json:"chat_id"`
	UserID  string      `json:"user_id"`
	Email   string      `json:"email"`
	Content PostContent `json:"content"`
}

type PostContent struct {
	Post string `json:"post"`
}

func (m *PostMessage) Body() map[string]interface{} {
	m.MsgType = MsgTypeText
	return structToMap(m)
}

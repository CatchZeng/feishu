package feishu

type TextMessage struct {
	MsgType MsgType `json:"msg_type"`
	Content Content `json:"content"`
}

type Content struct {
	Text string `json:"text"`
}

func (m *TextMessage) Body() map[string]interface{} {
	m.MsgType = MsgTypeText
	return structToMap(m)
}

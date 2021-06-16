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

// NewTextMessage new message
func NewTextMessage() *TextMessage {
	msg := TextMessage{}
	return &msg
}

// SetText set text
func (m *TextMessage) SetText(text string) *TextMessage {
	m.Content = Content{
		Text: text,
	}
	return m
}

package text

import (
	"fmt"
	//"strings"
)

type Chat struct {
	messages []ChatMessage
	cur      int
}

func NewChat(msgs ...ChatMessage) Chat {
	return Chat{}
}

type ChatMessage struct {
	Bold   bool //加粗
	Italic bool //斜体

	Underlined    bool //下划线
	Strikethrough bool //删除线
	Obfuscated    bool //混淆模式，该组件会在相同宽度的字符之间随机切换。

	Text string

	Insertion string

	ClickEvent ClickEvent
}

func (message ChatMessage) JSON() (str string) {
	str += "{"
	if message.Bold {
		str += `"bold": "true",`
	}

	if message.Italic {
		str += `"italic": "true",`
	}

	if message.Underlined {
		str += `"underlined": "true",`
	}

	if message.Strikethrough {
		str += `"strikethrough": "true",`
	}

	if message.Obfuscated {
		str += `"obfuscated": "true",`
	}

	if message.Insertion != "" {
		str += fmt.Sprintf(`"insertion": "%s"`, message.Insertion)
	}

	str += fmt.Sprintf(`"text": "%s"`, message.Text)

	str += "}"
	return
}

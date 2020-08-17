package text

//"strings"

type ChatComponent interface {
	SetText(text string)
	SetColor(color color)
	SetInsertion(insertion string)
	SetBold(bold bool)
	SetItalic(italic bool)
	SetUnderlined(underlined bool)
	SetStrikethrough(strikethrough bool)
	SetObfuscated(obfuscated bool)
	SetClickEvent(event clickEvent)
	SetHoverEvent(event hoverEvent)
	Extra(chat Chat)
}

type Chat struct {
	Text          string `json:"text"`
	Insertion     string `json:"insertion"`
	Color         color  `json:"color"`
	Bold          bool   `json:"bold"`          //加粗
	Italic        bool   `json:"italic"`        //斜体
	Underlined    bool   `json:"underlined"`    //下划线
	Strikethrough bool   `json:"strikethrough"` //删除线
	Obfuscated    bool   `json:"obfuscated"`    //混淆模式，该组件会在相同宽度的字符之间随机切换。
	extra         []Chat
	click         clickEvent
	hover         hoverEvent
}

func (chat *Chat) SetText(text string) {
	chat.Text = text
}

func (chat *Chat) SetColor(color color) {
	chat.Color = color
}

func (chat *Chat) SetBold(bold bool) {
	chat.Bold = bold
}

func (chat *Chat) SetItalic(italic bool) {
	chat.Italic = italic
}

func (chat *Chat) SetUnderlined(underlined bool) {
	chat.Underlined = underlined
}

func (chat *Chat) SetStrikethrough(strikethrough bool) {
	chat.Strikethrough = strikethrough
}

func (chat *Chat) SetObfuscated(obfuscated bool) {
	chat.Obfuscated = obfuscated
}

func (chat *Chat) SetInsertion(insertion string) {
	chat.Insertion = insertion
}

func (chat *Chat) SetClickEvent(event clickEvent) {
	chat.click = event
}

func (chat *Chat) SetHoverEvent(event hoverEvent) {
	chat.hover = event
}

func (chat *Chat) Extra(c Chat) {
	chat.extra = append(chat.extra, c)
	return
}

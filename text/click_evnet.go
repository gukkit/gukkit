package text

type clickEventAction string

const (
	ClickEventOpenURLAction    = clickEventAction("open_url")
	ClickEventRunCommandAction = clickEventAction("run_command")
	ClickEventChangePageAction = clickEventAction("change_page")
)

type clickEvent struct {
	action clickEventAction
	value  interface{}
}

func NewClickEvent(action clickEventAction, value interface{}) (event clickEvent) {
	switch action {
	case ClickEventOpenURLAction, ClickEventRunCommandAction:
		if _, ok := value.(string); !ok {
			panic("ClickEventChangePageAction must value string type")
		}
	case ClickEventChangePageAction:
		if _, ok := value.(int); !ok {
			panic("ClickEventChangePageAction must value int type")
		}
	}

	event.action = action
	event.value = value
	return
}

func newChangePageClickEvent(changePage int) (event clickEvent) {
	event.action = "change_page"
	event.value = changePage
	return
}

func newOpenURLClickEvent(url string) (event clickEvent) {
	event.action = "open_url"
	event.value = url
	return
}

func newRunCommandClickEvent(cmd string) (event clickEvent) {
	event.action = "run_command"
	event.value = cmd
	return
}

package plugin

type Plugin interface {
	onEnable()
	onDisable()
	onReload()
}

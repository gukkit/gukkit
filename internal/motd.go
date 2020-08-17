package internal

import "gukkit/text"

var defaultMotdTitle = "Gukkit Server for Minecraft By iHDJ"

type Motd struct {
	Version     MotdVersion `json:"version"`
	Description text.Chat   `json:"description"`
	Favicon     string      `json:"favicon"`
}

type MotdVersion struct {
	Name     string `json:"name"`
	Protocol int    `json:"protocol"`
}

type MotdPlayers struct {
	Max    int              `json:"max"`
	Online int              `json:"online"`
	Sample []MotdPlayerInfo `json:"sample"`
}

type MotdPlayerInfo struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}

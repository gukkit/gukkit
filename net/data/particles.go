package data

const (
	AmbientEntityEffect = "minecraft:ambient_entity_effect"
	AngryVillager       = "minecraft:angry_villager"
	Barrier             = "minecraft:barrier"
	Block               = "minecraft:block"
)

type ParticleType struct {
	Name string
	ID   int
	Data []byte
}

type particle interface {
}

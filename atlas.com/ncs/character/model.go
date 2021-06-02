package character

type Model struct {
	id           uint32
	level        byte
	meso         uint32
	jobId        uint16
	strength     uint16
	dexterity    uint16
	intelligence uint16
	mapId        uint32
	gender       byte
	hair         uint32
	face         uint32
}

func (a Model) Id() uint32 {
	return a.id
}

func (a Model) Level() byte {
	return a.level
}

func (a Model) Meso() uint32 {
	return a.meso
}

func (a Model) JobId() uint16 {
	return a.jobId
}

func (a Model) Dexterity() uint16 {
	return a.dexterity
}

func (a Model) Intelligence() uint16 {
	return a.intelligence
}

func (a Model) Gender() byte {
	return a.gender
}

func (a Model) Hair() uint32 {
	return a.hair
}

func (a Model) Strength() uint16 {
	return a.strength
}

func (a Model) Face() uint32 {
	return a.face
}

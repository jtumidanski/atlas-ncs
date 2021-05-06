package character

type Model struct {
	id        uint32
	level     byte
	meso      uint32
	jobId     uint16
	dexterity uint16
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

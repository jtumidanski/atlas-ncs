package event

type Model struct {
	id    uint32
	mapId uint32
	limit uint32
}

func (m Model) Id() uint32 {
	return m.id
}

func (m Model) Limit() uint32 {
	return m.limit
}

func (m Model) MapId() uint32 {
	return m.mapId
}

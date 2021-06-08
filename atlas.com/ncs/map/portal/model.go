package portal

type Model struct {
	id   uint32
	name string
}

func (m *Model) Id() uint32 {
	return m.id
}

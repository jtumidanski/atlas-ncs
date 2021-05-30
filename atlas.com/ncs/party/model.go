package party

type Model struct {
	id      uint32
	members []Member
}

func (m Model) Members() []Member {
	return m.members
}

func (m Model) Id() uint32 {
	return m.id
}

type Member struct {
	id uint32
}

func (m Member) Id() uint32 {
	return m.id
}

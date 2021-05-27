package party

type Model struct {
	members []Member
}

func (m Model) Members() []Member {
	return m.members
}

type Member struct {
	id uint32
}

func (m Member) Id() uint32 {
	return m.id
}

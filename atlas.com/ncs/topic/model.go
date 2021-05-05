package topic

type Supplier func(name string) (*Model, error)

type Model struct {
	name string
}

func (m Model) Name() string {
	return m.name
}

package exclude

type Model struct {
	id     uint32
	itemId uint32
}

func (m Model) Id() uint32 {
	return m.id
}

func (m Model) ItemId() uint32 {
	return m.itemId
}

func NewModel(id uint32, itemId uint32) Model {
	return Model{
		id:     id,
		itemId: itemId,
	}
}

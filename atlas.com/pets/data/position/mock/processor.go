package pmock

import (
	"atlas-pets/data/position"
	"github.com/Chronicle20/atlas-model/model"
)

type Processor struct {
	GetBelowFn func(mapId uint32, x int16, y int16) model.Provider[position.Model]
}

func (m *Processor) GetBelow(mapId uint32, x int16, y int16) model.Provider[position.Model] {
	if m.GetBelowFn != nil {
		return m.GetBelowFn(mapId, x, y)
	}
	return nil
}

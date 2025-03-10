package pet

import (
	"errors"
	"strconv"
	"time"
)

type RestModel struct {
	Id              uint32    `json:"-"`
	InventoryItemId uint32    `json:"inventoryItemId"`
	TemplateId      uint32    `json:"templateId"`
	Name            string    `json:"name"`
	Level           byte      `json:"level"`
	Tameness        uint16    `json:"tameness"`
	Fullness        byte      `json:"fullness"`
	Expiration      time.Time `json:"expiration"`
	OwnerId         uint32    `json:"ownerId"`
	Lead            bool      `json:"lead"`
	Slot            byte      `json:"slot"`
	X               int16     `json:"x"`
	Y               int16     `json:"y"`
	Stance          byte      `json:"stance"`
}

func (r RestModel) GetName() string {
	return "pets"
}

func (r RestModel) GetID() string {
	return strconv.Itoa(int(r.Id))
}

func (r *RestModel) SetID(strId string) error {
	id, err := strconv.Atoi(strId)
	if err != nil {
		return err
	}
	r.Id = uint32(id)
	return nil
}

func Transform(m Model) (RestModel, error) {
	tm := GetTemporalRegistry().GetById(m.Id())
	if tm == nil {
		return RestModel{}, errors.New("temporal data not found")
	}

	return RestModel{
		Id:              m.id,
		InventoryItemId: m.inventoryItemId,
		TemplateId:      m.templateId,
		Name:            m.name,
		Level:           m.level,
		Tameness:        m.tameness,
		Fullness:        m.fullness,
		Expiration:      m.expiration,
		OwnerId:         m.ownerId,
		Lead:            m.lead,
		Slot:            m.slot,
		X:               tm.X(),
		Y:               tm.Y(),
		Stance:          tm.Stance(),
	}, nil
}

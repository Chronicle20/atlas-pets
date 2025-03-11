package pet

import (
	"errors"
	"github.com/Chronicle20/atlas-tenant"
	"gorm.io/gorm"
)

func create(db *gorm.DB) func(t tenant.Model, ownerId uint32, m Model) (Model, error) {
	return func(t tenant.Model, ownerId uint32, m Model) (Model, error) {
		e := &Entity{
			TenantId:        t.Id(),
			OwnerId:         ownerId,
			InventoryItemId: m.InventoryItemId(),
			TemplateId:      m.TemplateId(),
			Name:            m.Name(),
			Level:           m.Level(),
			Tameness:        m.Tameness(),
			Fullness:        m.Fullness(),
			Expiration:      m.Expiration(),
		}

		err := db.Create(e).Error
		if err != nil {
			return Model{}, err
		}
		return modelFromEntity(*e)
	}
}

func updateSlot(db *gorm.DB) func(t tenant.Model, petId uint64, slot int8) error {
	return func(t tenant.Model, petId uint64, slot int8) error {
		result := db.Model(&Entity{}).
			Where("tenant_id = ? AND id = ?", t.Id(), petId).
			Update("slot", slot)

		if result.Error != nil {
			return result.Error
		}

		if result.RowsAffected == 0 {
			return errors.New("no entity found or slot is already set to the given value")
		}

		return nil
	}
}

func updateTameness(db *gorm.DB) func(t tenant.Model, petId uint64, tameness uint16) error {
	return func(t tenant.Model, petId uint64, tameness uint16) error {
		result := db.Model(&Entity{}).
			Where("tenant_id = ? AND id = ?", t.Id(), petId).
			Update("tameness", tameness)

		if result.Error != nil {
			return result.Error
		}

		if result.RowsAffected == 0 {
			return errors.New("no entity found or tameness is already set to the given value")
		}

		return nil
	}
}

func deleteByInventoryItemId(t tenant.Model, inventoryItemId uint32) func(db *gorm.DB) error {
	return func(db *gorm.DB) error {
		return db.Where(&Entity{TenantId: t.Id(), InventoryItemId: inventoryItemId}).Delete(&Entity{}).Error
	}
}

func deleteForCharacter(t tenant.Model, ownerId uint32) func(db *gorm.DB) error {
	return func(db *gorm.DB) error {
		return db.Where(&Entity{TenantId: t.Id(), OwnerId: ownerId}).Delete(&Entity{}).Error
	}
}

func modelFromEntity(e Entity) (Model, error) {
	return NewModelBuilder(e.Id, e.InventoryItemId, e.TemplateId, e.Name, e.OwnerId).
		SetLevel(e.Level).
		SetTameness(e.Tameness).
		SetFullness(e.Fullness).
		SetExpiration(e.Expiration).
		SetSlot(e.Slot).
		Build(), nil
}

package pet

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

func Migration(db *gorm.DB) error {
	return db.AutoMigrate(&Entity{})
}

type Entity struct {
	TenantId        uuid.UUID `gorm:"not null;"`
	CharacterId     uint32    `gorm:"not null;"`
	Id              uint32    `gorm:"primary_key;auto_increment"`
	InventoryItemId uint32    `gorm:"not null"`
	TemplateId      uint32    `gorm:"not null"`
	Name            string    `gorm:"size:13;not null"`
	Level           byte      `gorm:"not null;default:1"`
	Tameness        uint16    `gorm:"not null;default:0"`
	Fullness        byte      `gorm:"not null;default:100"`
	Expiration      time.Time `gorm:"not null;"`
}

func (e Entity) TableName() string {
	return "pets"
}

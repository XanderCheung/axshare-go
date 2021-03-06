package models

import (
	"axshare_go/internal/db"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Axure struct {
	gorm.Model
	Name         string       `json:"name"`
	SecretKey    string       `json:"secret_key" binding:"required"`
	AxureGroupId uint         `json:"axure_group_id"`
	AxureGroup   AxureGroup   `gorm:"foreignKey:AxureGroupId"`
	Attachments  []Attachment `json:"attachments" gorm:"foreignKey:AxureId"`
}

func FindAxure(id uint) Axure {
	axure := Axure{}
	db.AxshareDb.Find(&axure, id)
	return axure
}

func (c *Axure) BeforeCreate(tx *gorm.DB) (err error) {
	c.genSecretKey()
	return
}

// 原型静态web链接
func (c *Axure) WebLink() string {
	attachment := c.LatestAttachment()
	return attachment.WebLink()
}

// 原型永久地址
func (c *Axure) PermanentLink() string {
	return fmt.Sprintf("%s/#/axures/%d?key=%s", CacheConfig.WebDomain, c.ID, c.SecretKey)
}

func (c *Axure) LatestAttachment() Attachment {
	attachment := Attachment{}
	db.AxshareDb.Model(&attachment).
		Where("axure_id = ?", c.ID).
		Last(&attachment)
	return attachment
}

// generate secret key
func (c *Axure) genSecretKey() {
	c.SecretKey = uuid.NewV4().String()
}

package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"strings"
)

type Axure struct {
	gorm.Model
	Name         string       `json:"name"`
	Desc         string       `json:"desc"`
	Link         string       `json:"link"`
	Key          string       `json:"key"`
	AxureGroupId uint         `gorm:"index"json:"axure_group_id"`
	Attachments  []Attachment `json:"attachments"`
}

// 原型静态web链接
func (c *Axure) WebLink() string {
	if !c.IsReleased() {
		return ""
	}
	webHost := viper.GetString("web_host")
	return webHost + c.Link
}

// 原型永久地址
func (c *Axure) PermanentLink() string {
	adminHost := viper.GetString("admin_host")
	permanentLink := strings.Join([]string{
		adminHost, "/axures/", fmt.Sprint(c.ID), "?key=", c.Key}, "")
	return permanentLink
}

// 文件是否解压
func (c *Axure) IsReleased() bool {
	isReleased := len(c.Link) > 0
	return isReleased
}

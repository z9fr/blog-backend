package types

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Title       string         `gorm:"column:title" json:"title"`
	Descrption  string         `gorm:"column:description" json:"description"`
	Body        string         `gorm:"column:body" json:"body"`
	HeaderImage string         `gorm:"column:headerImage" json:"headerImage"`
	CreatedBy   string         `gorm:"column:created_by" json:"created_by" `
	Slug        string         `gorm:"column:slug" json:"slug"`
	IsPublic    bool           `gorm:"column:is_public" json:"is_public"`
	Tags        pq.StringArray `gorm:"type:varchar(64)[]" json:"tags" swaggertype:"string"`
}

type PostCreateRequest struct {
	Title       string         `json:"title"`
	Description string         `json:"description"`
	Body        string         `json:"body"`
	HeaderImage string         `json:"headerImage"`
	Tags        pq.StringArray `gorm:"type:varchar(64)[]" json:"tags" swaggertype:"string"`
}

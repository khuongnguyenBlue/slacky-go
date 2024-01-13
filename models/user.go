package models

import models_base "github.com/khuongnguyenBlue/slacky/models/base"

type User struct {
	models_base.BaseModel
	Email string `gorm:"uniqueIndex"`
	Hash string `gorm:"not null"`

	WorkspaceMembers []WorkspaceMember
}
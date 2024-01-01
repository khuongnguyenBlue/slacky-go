package models

import models_base "github.com/khuongnguyenBlue/slacky/models/base"

type Channel struct {
	models_base.BaseModel
	WorkspaceID uint   `gorm:"index:idx_channels_workspace_name,unique;not null"`
	Name        string `gorm:"index:idx_channels_workspace_name,unique;not null"`
	Type        string `gorm:"not null"`
	Status      string `gorm:"not null"`
	Description string

	Workspace Workspace
	Members   []WorkspaceMember `gorm:"many2many:channel_members"`
}

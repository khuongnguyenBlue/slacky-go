package models

import models_base "github.com/khuongnguyenBlue/slacky/models/base"

type WorkspaceMember struct {
	models_base.BaseModel
	UserID      uint   `gorm:"index;not null"`
	WorkspaceID uint   `gorm:"index:idx_workspace_members_workspace_tag_name,unique;not null"`
	TagName     string `gorm:"index:idx_workspace_members_workspace_tag_name,unique;not null"`
	DisplayName string `gorm:"not null"`
	NameToken   string `gorm:"->;type:tsvector GENERATED ALWAYS AS (to_tsvector('english', tag_name || ' ' || display_name)) STORED;index:,type:gin"`
	Role        string `gorm:"not null"`
	Status      string `gorm:"not null"`

	User      User
	Workspace Workspace
	Channels  []Channel `gorm:"many2many:channel_members"`
	Messages  []Message `gorm:"foreignKey:SenderID"`
	Threads   []Thread  `gorm:"foreignKey:SenderID"`
}

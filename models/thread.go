package models

import models_base "github.com/khuongnguyenBlue/slacky/models/base"

type Thread struct {
	models_base.BaseModel
	MessageID    uint   `gorm:"index"`
	SenderID     uint   `gorm:"index"`
	Content      string `gorm:"not null"`
	ContentToken string `gorm:"->;type:tsvector GENERATED ALWAYS AS (to_tsvector('english', content)) STORED;index:,type:gin"`

	Message Message
	Sender  WorkspaceMember
}

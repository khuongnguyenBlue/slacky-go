package models

import models_base "github.com/khuongnguyenBlue/slacky/models/base"

type Message struct {
	models_base.BaseModel
	ChannelID uint   `gorm:"index:idx_messages_channel_sender"`
	SenderID  uint   `gorm:"index:idx_messages_channel_sender;index"`
	Content   string `gorm:"not null"`
	ContentToken string `gorm:"->;type:tsvector GENERATED ALWAYS AS (to_tsvector('english', content)) STORED;index:,type:gin"`

	Channel Channel
	Sender  WorkspaceMember
}

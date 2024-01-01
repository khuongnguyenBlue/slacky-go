package models

import models_base "github.com/khuongnguyenBlue/slacky/models/base"

type Workspace struct {
	models_base.BaseModel
	Name string `gorm:"not null;uniqueIndex"`

	Members []WorkspaceMember
}

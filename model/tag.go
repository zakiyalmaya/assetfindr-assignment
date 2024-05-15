package model

type Tag struct {
	ID    uint    `gorm:"primaryKey" json:"id"`
	Label string  `gorm:"index" json:"label" validate:"required"`
	Posts []*Post `gorm:"many2many:post_tags;" json:"posts,omitempty"`
}

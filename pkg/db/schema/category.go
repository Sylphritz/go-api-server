package schema

type Category struct {
	BaseEntity
	Name        string `gorm:"uniqueIndex;not null" json:"name"`
	Description string `json:"description"`
	Posts       []Post `json:"-"`
	BlogID      uint   `json:"blogId"`
}

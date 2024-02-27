package schema

type Blog struct {
	BaseEntity
	Title       string     `gorm:"uniqueIndex;not null" json:"title"`
	Description string     `json:"description"`
	Posts       []Post     `json:"-"`
	Categories  []Category `json:"-"`
}

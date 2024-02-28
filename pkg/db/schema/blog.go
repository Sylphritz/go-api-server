package schema

type Blog struct {
	BaseEntity
	Title       string     `gorm:"uniqueIndex:idx_blog_title_user_id;not null" json:"title"`
	UserID      uint       `gorm:"uniqueIndex:idx_blog_title_user_id;not null" json:"userId"`
	Description string     `json:"description"`
	Posts       []Post     `json:"-"`
	Categories  []Category `json:"-"`
}

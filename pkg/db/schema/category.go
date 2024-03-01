package schema

type Category struct {
	BaseEntity
	Name        string `gorm:"not null" json:"name"`
	Slug        string `gorm:"uniqueIndex:idx_category_slug_blog_id;not null" json:"slug"`
	Description string `json:"description"`
	Posts       []Post `json:"-"`
	BlogID      uint   `gorm:"uniqueIndex:idx_category_slug_blog_id;not null" json:"blogId"`
}

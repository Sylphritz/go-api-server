package schema

type Category struct {
	BaseEntity
	Name        string `gorm:"uniqueIndex:idx_category_name_blog_id;not null" json:"name"`
	Description string `json:"description"`
	Posts       []Post `json:"-"`
	BlogID      uint   `gorm:"uniqueIndex:idx_category_name_blog_id;not null" json:"blogId"`
}

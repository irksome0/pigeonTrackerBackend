package models

type Blog struct {
	Id          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Image       string `json:"image"`
	UserID      string `json:"user_id"`
	User        User   `json:"user";gorm:"foreignkey:UserID"`
}

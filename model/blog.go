package model

import "time"

type GenericModel struct {
	Id        int       `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

type Post struct {
	GenericModel
	Id      uint64 `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func GetAllPosts() ([]Post, error) {
	var post []Post

	tx := db.Find(&post)

	if tx.Error != nil {
		return post, tx.Error
	}

	return post, nil
}

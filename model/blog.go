package model

type Article struct {
	Id      uint64 `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func GetAllPosts() ([]Article, error) {
	var articles []Article

	tx := db.Find(&articles)

	if tx.Error != nil {
		return articles, tx.Error
	}

	return articles, nil
}

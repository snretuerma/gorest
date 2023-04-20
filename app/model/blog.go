package model

func GetAllPosts() ([]Post, error) {
	var posts []Post

	tx := db.Find(&posts)

	if tx.Error != nil {
		return posts, tx.Error
	}

	return posts, nil
}

func GetPost(id string) (Post, error) {
	var post Post

	tx := db.First(&post, id)

	if tx.Error != nil {
		return post, tx.Error
	}

	return post, nil
}

func CreatePost(post Post) error {

	result := db.Create(&post)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func UpdatePost(id uint64, request Post) error {

	var post Post

	tx := db.First(&post, id)

	if tx.Error != nil {
		return tx.Error
	}

	result := db.Model(&post).Updates(&request)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func DeletePost(id uint64) error {
	var post Post

	tx := db.Delete(&post, id)

	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

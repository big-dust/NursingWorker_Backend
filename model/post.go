package model

type Post struct {
	ID      uint   `gorm:"primarykey;column:id"`
	OpenID  string `json:"-"`
	Title   string `form:"title" gorm:"column:title"`
	Content string `gorm:"column:content" form:"content"`
}

type Videos struct {
	ID     uint `gorm:"primarykey"`
	PostID string
	Path   string
}

type Images struct {
	ID     uint `gorm:"primarykey"`
	PostID string
	Path   string
}

func Create(post Post) (Post, error) {
	if err := DB.Create(&post).Error; err != nil {
		return Post{}, err
	}
	return post, nil
}

func CreateImage(postID string, path string) error {
	var image Images
	image.Path = path
	image.PostID = postID
	if err := DB.Create(&image).Error; err != nil {
		return err
	}
	return nil
}

func CreateVideo(postID string, path string) error {
	var video Videos
	video.PostID = postID
	video.Path = path
	if err := DB.Create(&video).Error; err != nil {
		return err
	}
	return nil
}

func IsMyPost(postID string, openID string) bool {
	var tmp Post
	if err := DB.Model(&Post{}).Select("id").Where("id = ? AND open_id= ?", postID, openID).Take(&tmp).Error; err != nil {
		return false
	}
	return true
}

func PtDt(postID string) {
	DB.Delete(&Post{}, "id = ?", postID)
}

func HisPost(openID string) ([]Post, error) {
	var posts []Post
	err := DB.Model(&Post{}).Where("open_id = ?", openID).Find(&posts).Error
	return posts, err
}

func HisLike(userID string) ([]Post, error) {
	var postIDs []string
	DB.Model(&Like{}).Select("post_id").Where("user_id= ?", userID).Find(&postIDs)
	var posts []Post
	err := DB.Model(&Post{}).Where("id in (?)", postIDs).Find(&posts).Error
	return posts, err
}

func HisColl(userID string) ([]Post, error) {
	var postIDs []string
	DB.Model(&Collection{}).Select("post_id").Where("user_id= ?", userID).Find(&postIDs)
	var posts []Post
	err := DB.Model(&Post{}).Where("id in (?)", postIDs).Find(&posts).Error
	return posts, err
}

func ReCmt(old []string, number string) ([]Post, error) {
	var posts []Post
	result := DB.Raw("SELECT * FROM posts ORDER BY RAND() LIMIT ?", number).Scan(&posts)
	if result.Error != nil {
		return nil, result.Error
	}
	return posts, nil
}

package models

import (
  "time"
)

type Post struct {
  Id       uint32      `gorm:"primary_key;auto_increment" json:"id"`
  ImageUrl  string     `gorm:"type:varchar(255)" json:"image_url"`
  Subtitle  string     `gorm:"type:varchar(50)" json:"subtitle"`
  UserId    uint32     `json:"user_id"`
  User      User       `json:"user"`
  CreatedAt time.Time  `json:"created_at"`
  UpdatedAt time.Time  `json:"updated_at"`
  Feedbacks []Feedback `gorm:"ForeigKey:PostId" json:"feedbacks"`
}

func NewPost(post Post) error {
  db := Connect()
  defer db.Close()
  rs := db.Create(&post)
  return rs.Error
}

func GetPosts() []Post {
  db := Connect()
  defer db.Close()
  var posts []Post
  db.Order("id asc").Find(&posts)
  for i, _ := range posts {
    db.Model(posts[i]).Related(&posts[i].User)
    posts[i].Feedbacks = GetFeedbacksByPost(posts[i])
  }
  return posts
}

func UpdatePost(post Post) (int64, error) {
  db := Connect()
  defer db.Close()
  rs := db.Model(&post).Where("id = ?", post.Id).UpdateColumns(
    map[string]interface{}{
      "image_url": post.ImageUrl,
      "subtitle": post.Subtitle,
    },
  )
  return rs.RowsAffected, rs.Error
}
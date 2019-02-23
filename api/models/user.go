package models

import (
  "time"
  "fmt"
  "api/security"
)

type User struct {
  Id        uint32     `gorm:"primary_key;auto_increment" json:"id"`
  Nickname  string     `gorm:"type:varchar(20);unique_index;not null" json:"nickname"`
  Email     string     `gorm:"type:varchar(40);unique_index;not null" json:"email"`
  Password  string     `gorm:"size:60;not null" json:"password"`
  CreatedAt time.Time  `gorm:"default:current_timestamp()" json:"created_at"`
  UpdatedAt time.Time  `gorm:"default:current_timestamp()" json:"updated_at"`
  Posts     []Post     `gorm:"ForeignKey:UserId" json:"posts"`
  Feedbacks []Feedback `gorm:"ForeignKey:UserId" json:"feedbacks"`
}

func NewUser(user User) error {
  hashedPassword, err := security.Hash(user.Password)
  if err != nil {
    return err
  }
  user.Password = fmt.Sprintf("%s", hashedPassword)
  db := Connect()
  defer db.Close()
  rs := db.Create(&user)
  return rs.Error
}

func UpdateUser(user User) (int64, error) {
  db := Connect()
  defer db.Close()
  rs := db.Model(&user).Where("id = ?", user.Id).UpdateColumns(
    map[string]interface{}{
      "nickname": user.Nickname,
      "email": user.Email,
    },
  )
  return rs.RowsAffected, rs.Error
}
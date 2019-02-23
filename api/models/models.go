package models

import "github.com/jinzhu/gorm"

const (
  USERS     = "users"
  POSTS     = "posts"
  FEEDBACKS = "feedbacks"
)

func GetAll(table string) interface{} {
  db := Connect()
  defer db.Close()
  switch table {
    case USERS:
      return db.Order("id asc").Find(&[]User{}).Value
    case POSTS:
      return db.Order("id asc").Find(&[]Post{}).Value
    case FEEDBACKS:
      return db.Order("id asc").Find(&[]Feedback{}).Value
  }
  return nil
}

func GetById(table string, id uint64) interface{} {
  db := Connect()
  defer db.Close()
  switch table {
    case USERS:
      return db.Where("id = ?", id).Find(&User{}).Value
    case POSTS:
      return db.Where("id = ?", id).Find(&Post{}).Value
    case FEEDBACKS:
      return db.Where("id = ?", id).Find(&Feedback{}).Value
  }
  return nil
}

func Delete(table string, id uint64) (int64, error) {
  db := Connect()
  defer db.Close()
  var rs *gorm.DB
  switch table {
  case USERS:
    rs = db.Where("id = ?", id).Delete(&User{})
    break
  case POSTS:
    rs = db.Where("id = ?", id).Delete(&Post{})
    break
  case FEEDBACKS:
    rs = db.Where("id = ?", id).Delete(&Feedback{})
  }
  return rs.RowsAffected, rs.Error
}

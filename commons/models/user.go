package models

// Gorm 配置数据库迁移对象

// User 用户模型
type User struct {
	BaseModel
	Name   *string `gorm:"column:name;size:12" json:"name"`
	Email  string  `gorm:"column:email;size:20;default:'han';uniqueIndex:index_email_card_id" json:"email"`
	CardId string  `gorm:"column:card_id;size:20;not null;uniqueIndex:index_email_card_id" json:"card_id"`
	Age    *int    `gorm:"column:age;size:5;comment:'age要小于360'" json:"age"`
}

func (User) TableName() string {
	return "user"
}
func (User) TableComment() string {
	return "这是user表"
}

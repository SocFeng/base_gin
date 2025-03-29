package dal

import (
	"base_gin/commons/dataBase"
	"base_gin/commons/models"
	"fmt"
)

func GetAllUser() (resAllUser []models.User) {
	tx := dataBase.GlobalGormDB.Find(&resAllUser)
	if tx.Error != nil {
		return nil
	}
	return resAllUser

}
func InsertOneUser(user *models.User) error {

	tx := dataBase.GlobalGormDB.Create(user)
	if tx.Error != nil {
		return fmt.Errorf("插入出错")
	}
	return nil
}

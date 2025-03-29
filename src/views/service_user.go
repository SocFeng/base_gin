package views

import (
	"base_gin/commons/dataBase"
	"base_gin/commons/logs"
	"base_gin/commons/middleware"
	"base_gin/commons/models"
	dal "base_gin/src/dal/user"
	"fmt"
	"github.com/gin-gonic/gin"
	"os/user"
)

func UserName(c *gin.Context) {

	c.JSON(200, gin.H{
		"data": "This is UserName!",
	})
}

// InsertOneUserData 插入单条User记录
func InsertOneUserData(c *gin.Context) {
	// 创建一个用户信息
	user := models.User{Name: models.Ptr("fenghan"), Email: "2418028749@qq.com", CardId: "123456789"}
	// 执行插入
	tx := dataBase.GlobalGormDB.Create(&user)
	if tx.Error != nil {
		logs.AppError(tx.Error)
	}
	c.JSON(200, gin.H{
		"msg":   "数据插入成功！",
		"title": "InsertOneUserData",
		"data":  tx.RowsAffected,
	})
}

// InsertManyUserData 插入多条user记录
func InsertManyUserData(c *gin.Context) {

	users := []models.User{
		{Name: models.Ptr("zhangsan"), Email: "123@qq.com", CardId: "123"},
		{Name: models.Ptr("lisi"), Email: "456@qq.com", CardId: "123", Age: models.Ptr(20)},
		{Name: models.Ptr("wanger"), Email: "789@qq.com", CardId: "789"},
	}
	tx := dataBase.GlobalGormDB.Create(&users)
	if tx.Error != nil {
		logs.AppError(tx.Error)
	}

	c.JSON(200, gin.H{
		"msg":   "数据插入成功",
		"title": "InsertOneUserData",
		"data":  tx.RowsAffected,
	})
}

// InsertUserDataReturnId  插入数据并返回id
func InsertUserDataReturnId(c *gin.Context) {

	users := []models.User{
		{Name: models.Ptr("liu4"), Email: "1111@qq.com", CardId: "999", Age: nil},
		{Email: "2222@qq.com", CardId: "555", Age: models.Ptr(20)},
		{Name: models.Ptr("mei7"), Email: "333@qq.com", CardId: "333"},
	}
	_ = dataBase.GlobalGormDB.Create(&users)
	var ids []int
	for _, user := range users {
		ids = append(ids, user.ID)
	}
	c.JSON(200, gin.H{
		"msg":   "数据插入成功",
		"title": "InsertManyUserDataReturnId",
		"data":  ids,
	})
}

func QueryBaseUserData(c *gin.Context) {

	// 查询全部users
	queryAllUser := dal.GetAllUser()

	// 查询单一用户
	var userIdOne user.User
	//dataBase.GlobalGormDB.First(&user, 2) // 查询 id = 2的一条数据
	tx := dataBase.GlobalGormDB.Where("id = ?", 2).First(&userIdOne)
	if tx.Error != nil {
		fmt.Println("ERRROR")
	}

	c.JSON(200, gin.H{
		"msg":          "数据插入成功",
		"title":        "InsertManyUserDataReturnId",
		"queryAllUser": queryAllUser,
		"userIdOne":    userIdOne,
	})
}

func QueryWhereUserData(c *gin.Context) {

	// Where 条件查询
	var queryWhereNameUsers []user.User
	// 查询所有 name = fenghan的数据
	_ = dataBase.GlobalGormDB.Where("name = ?", "fenghan").Find(&queryWhereNameUsers)

	// 多字段查询
	var queryWhereNameEmailUsers []user.User
	_ = dataBase.GlobalGormDB.Where("name = ? and email = ?", "fenghan", "1111@qq.com").Find(&queryWhereNameEmailUsers)

	// 飞非空查询
	var queryWhereNameNotNullUsers []user.User
	_ = dataBase.GlobalGormDB.Where("age IS NOT NULL").Find(&queryWhereNameNotNullUsers)

	// IN 条件查询
	ids := []int{1, 3, 4}
	var queryInIdsUsers []user.User
	dataBase.GlobalGormDB.Where("id IN (?)", ids).Find(&queryInIdsUsers)

	// 分页查询
	pageSize := 1
	pageNum := 2
	var queryPageUsers []user.User
	// 先 name 倒叙排列 在取分页数据 最后赋值给 queryPageUsers
	dataBase.GlobalGormDB.Order("name desc").Limit(pageNum).Offset(pageSize).Find(&queryPageUsers)

	// 查询指定的字段
	type UserInfo struct {
		Name *string `json:"name"`
		Age  *int    `json:"age"`
	}
	var userInfo []UserInfo
	// 指定查询的模型 --> 指定条件 --> 筛选字段 --> 排序 --> 赋值
	dataBase.GlobalGormDB.Model(user.User{}).Where("id in (?)", []int{2, 3}).Select("name", "age").Order("age desc").Find(&userInfo)

	// 统计数量
	var count int64 // 数量只能是int64
	dataBase.GlobalGormDB.Model(user.User{}).Where("age IS NOT NULL").Count(&count)

	var queryAllUserMap []map[string]interface{}
	var queryAllUserMapSelect []map[string]interface{}

	dataBase.GlobalGormDB.Model(user.User{}).Select("*").Where("name = ?", "fenghan").Find(&queryAllUserMap)
	dataBase.GlobalGormDB.Model(user.User{}).Select("id", "name", "age", "create_time").Where("name = ?", "fenghan").Find(&queryAllUserMapSelect)
	c.JSON(200, gin.H{
		"msg":                        "数据插入成功",
		"queryWhereNameUsers":        queryWhereNameUsers,
		"queryWhereNameEmailUsers":   queryWhereNameEmailUsers,
		"queryWhereNameNotNullUsers": queryWhereNameNotNullUsers,
		"queryInIdsUsers":            queryInIdsUsers,
		"queryPageUsers":             queryPageUsers,
		"userInfo":                   userInfo,
		"count":                      count,
		"queryAllUserMap":            queryAllUserMap,
		"queryAllUserMapSelect":      queryAllUserMapSelect,
	})
}

func QueryUsersCountSQL(c *gin.Context) {
	resData := dataBase.QueryToMap(`SELECT count(*) as num FROM user`)
	ids := []int{1, 2, 3, 4}
	inClause, inArgs := dataBase.BuildInQuery(ids)
	query := fmt.Sprintf("SELECT * FROM user WHERE id IN %s AND name = ?", inClause)
	inArgs = append(inArgs, "fenghan")
	resDataIn := dataBase.QueryToMap(query, inArgs...)
	c.JSON(200, gin.H{
		"data":        resData,
		"resDataExec": resDataIn,
	})
}

func GetPostData(c *gin.Context) {
	data := middleware.GetRequestData(c)
	if data == nil {
		c.JSON(200, gin.H{
			"msg": "参数异常!",
		})
	}

	fmt.Println("req", data)

	reqJsonData := data["json"]
	reqQueryData := data["query"]
	reqFormData := data["form"]
	reqFileData := data["files"]
	c.JSON(200, gin.H{
		"msg":          "success",
		"reqJsonData":  reqJsonData,
		"reqQueryData": reqQueryData,
		"reqFormData":  reqFormData,
		"reqFileData":  reqFileData,
	})
}

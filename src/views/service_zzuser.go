package views

import (
	"base_gin/src/dal"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func CreateOneZzUser(c *gin.Context) {
	now := time.Now()

	affectedRows, lastInsertId := dal.DalCreateOneZzUser([]any{"冯含", "fenghan1", 20, "2418028749@aa.com", 33, now, now})
	c.JSON(200, gin.H{
		"msg":          "ok",
		"affectedRows": affectedRows,
		"lastInsertId": lastInsertId,
	})
}

func CreateManyZzUser(c *gin.Context) {
	now := time.Now()
	manyZzUser := [][]any{
		{"张三2", "zhangsna2", 60, nil, 22, now, now},
		{"李四2", "lisi2", 42, "2418028749@aa.com2", nil, now, now},
		{"王二2", "wanger2", 18, "2418028749@aa.com2", 88, now, now},
	}
	affectedRows, lastInsertId := dal.DalCreateManyZzUser(manyZzUser)
	c.JSON(200, gin.H{
		"msg":          "ok",
		"affectedRows": affectedRows,
		"lastInsertId": lastInsertId,
	})
}
func QueryAllZzUser(c *gin.Context) {
	queryAllZzUser := dal.DalQueryAllZzUser()

	c.JSON(200, gin.H{
		"msg":            "ok",
		"queryAllZzUser": queryAllZzUser,
	})
}

func QueryWhereZzUser(c *gin.Context) {
	ids := []any{1, 2, 3, 4, 5, 6, 7, 8, 9}
	maxAge := 60

	queryWhereZzUser := dal.DalQueryWhereZzUser(ids, maxAge)
	c.JSON(200, gin.H{
		"msg":              "ok",
		"queryWhereZzUser": queryWhereZzUser,
	})
}

func ChangeWhereZzUser(c *gin.Context) {
	ids := []any{4, 5}
	affectedRows := dal.DalChangeWhereZzUser(ids, "CC@aa.com")
	c.JSON(200, gin.H{
		"msg":          "ok",
		"affectedRows": affectedRows,
	})
}

func DeleteWhereZzUser(c *gin.Context) {

	affectedRows := dal.DalDeleteWhereZzUser(1)
	c.JSON(200, gin.H{
		"msg":          "ok",
		"affectedRows": affectedRows,
	})

}

func QueryJoinZzUser(c *gin.Context) {
	ids := []any{2, 3, 4, 5, 6, 7, 8}
	ageMax := 60
	infoCount := 60

	queryJoinZzUser := dal.DalQueryJoinUser(ids, ageMax, infoCount)
	for idx, val := range queryJoinZzUser {
		fmt.Println(idx, val)
		if val["count"] == nil {
			fmt.Println("COUNT is NULL")
		}
	}
	c.JSON(200, gin.H{
		"msg":             "ok",
		"queryJoinZzUser": queryJoinZzUser,
	})
}

func TxExecTimeUser(c *gin.Context) {

	err := dal.DalTestExecTx()
	if err != nil {
		// 出错退出
		c.JSON(400, gin.H{
			"msg":   "执行出错",
			"error": err.Error(),
		})
		// 必须要写 ee
		return
	}
	// 正常退出
	c.JSON(200, gin.H{
		"msg": "执行成功",
	})

}

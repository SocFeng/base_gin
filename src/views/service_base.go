package views

import (
	"base_gin/commons/cache"
	"base_gin/commons/dataBase"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func BseReturn(c *gin.Context) {

	resData := dataBase.QueryToMap(`SELECT * FROM newtable`)
	fmt.Println(resData)

	_ = cache.GlobalRedis.Set(c, "name", "fanghan", 5*time.Second).Err()
	time.Sleep(6 * time.Second)
	val, err := cache.GlobalRedis.Get(c, "name").Result()
	if err != nil {
		fmt.Println(err, "没有拿到数据")
	}

	c.JSON(200, gin.H{
		"data":  resData,
		"cache": val,
	})
}

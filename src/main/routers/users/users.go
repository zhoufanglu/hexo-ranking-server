package users

import (
	"fmt"
	"fuck-go/src/main/db"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"time"
)

// Users 结构体定义  对应数据库名称
type Users struct {
	Id       string `gorm:"column:id" json:"id"`
	Name     string `gorm:"column:name" json:"name"`
	Qq       string `gorm:"column:qq" json:"qq"`
	CreateAt string `gorm:"column:create_at" json:"createAt"`
}

type requestInsertType struct {
	Name string `json:"name"`
	Qq   string `json:"qq"`
}

type requestDeleteType struct {
	Id string `json:"id"`
}

func InsertUser(c *gin.Context) {
	// ?获取入参
	var userRequest requestInsertType
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// ?判断名字有没有重复
	users := SelectUserBySql()
	if containsName(users, userRequest.Name) {
		c.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "已有存在的名字",
		})
		return
	}

	// ?拼接 插入
	var user Users
	user.CreateAt = time.Now().Format("2006-01-02 15:04:05")
	user.Id = uuid.New().String()
	user.Name = userRequest.Name
	user.Qq = userRequest.Qq

	// 使用 GORM 创建记录
	result := db.Db.Create(&user)

	if result.Error != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "插入成功",
	})
}

func GetUsers(c *gin.Context) {
	// 查询数据库
	var users []Users
	users = SelectUserBySql()

	// 返回用户信息
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": users,
	})
}

func DeleteUser(c *gin.Context) {
	// ?获取入参
	var userRequest requestDeleteType
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Db.Exec(fmt.Sprintf("delete from Users where id='%s'", userRequest.Id))
	db.Db.Exec(fmt.Sprintf("DELETE FROM Records WHERE user_id='%s'", userRequest.Id))

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "删除成功",
	})
}

func SelectUserBySql() []Users {
	var users []Users
	if err := db.Db.Order("create_at ASC").Find(&users).Error; err != nil {
		fmt.Println("Error querying Users table:", err)
	}
	return users
}

// 判断 name 是否存在于 Users 结构体切片中
func containsName(users []Users, name string) bool {
	for _, user := range users {
		if user.Name == name {
			return true
		}
	}
	return false
}

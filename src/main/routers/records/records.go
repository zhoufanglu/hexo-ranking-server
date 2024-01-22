package records

import (
	"fmt"
	"fuck-go/src/main/db"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"strings"
)

// Records 结构体定义  对应数据库名称
type Records struct {
	Id       string `gorm:"column:id" json:"id"`
	UserId   string `gorm:"column:user_id" json:"userId"`
	CreateAt string `gorm:"column:create_at" json:"createAt"`
}

type requestInsertType struct {
	UserId   string `json:"userId"`
	CreateAt string `json:"createAt"`
}

type UserDateListType struct {
	Id       string   `json:"userId"`
	Name     string   `json:"name"`
	Dates    string   `json:"dates"`
	DateList []string `json:"dateList"`
}

func GetRecords(c *gin.Context) {
	// ?获取入参
	createAt := c.Query("createAt")
	fmt.Print(30, createAt)

	var records []UserDateListType
	// Execute the SQL query
	result := db.Db.Raw(`
		SELECT
			U.id,
			U.name,
			GROUP_CONCAT(DATE_FORMAT(R.create_at, '%Y-%m-%d')) AS dates
		FROM
			Users U
		LEFT JOIN
			Records R ON U.id = R.user_id AND DATE_FORMAT(R.create_at, '%Y-%m') = '2024-01'
		GROUP BY
			U.id, U.name
		ORDER BY
    		U.create_at ASC;
	`).Scan(&records)

	// Check for errors
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	// 转换日期字符串为数组
	for i := range records {
		records[i].DateList = strings.Split(records[i].Dates, ",")
	}

	// 返回信息
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": records,
	})
}

func InsertRecord(c *gin.Context) {
	// ?获取入参
	var recordRequest requestInsertType
	if err := c.ShouldBindJSON(&recordRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// ?拼接 插入
	var record Records
	record.UserId = recordRequest.UserId
	record.CreateAt = recordRequest.CreateAt
	record.Id = uuid.New().String()

	// 使用 GORM 创建记录
	result := db.Db.Create(&record)

	if result.Error != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "插入成功",
	})
}

type requestDeleteType struct {
	Id string `json:"id"`
}

func DeleteRecord(c *gin.Context) {
	// ?获取入参
	var recordRequest requestDeleteType
	if err := c.ShouldBindJSON(&recordRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Db.Exec(fmt.Sprintf("delete from Records where id='%s'", recordRequest.Id))

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "删除成功",
	})
}

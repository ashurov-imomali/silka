package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
	"net/http"
	"time"
)

var db *gorm.DB

func main() {
	dbSettings := `host=localhost port=5432 user=admin password=admin dbname=mydb sslmode=disable`
	db1, err := gorm.Open("postgres", dbSettings)
	db = db1
	if err != nil {
		log.Println(err)
		return
	}
	engine := gin.New()

	engine.POST("/ping", func(c *gin.Context) {
		c.YAML(200, gin.H{"message": "pong"})
	})

	engine.GET("/struct", GetTest)

	err = http.ListenAndServe("localhost:8080", engine)
	if err != nil {
		log.Println(err)
		return
	}

}

type TestStruct struct {
	ID        int64      `gorm:"column:id" json:"id"`
	Name      string     `gorm:"column:name" json:"name"`
	CreatedBy int64      `gorm:"column:created_by" json:"created_by"`
	CreatedAt time.Time  `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time  `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at" json:"deleted_at"`
}

func (TestStruct) TableName() string {
	return "test"
}

type Users struct {
	ID   int64  `gorm:"column:id" json:"id"`
	Name string `gorm:"column:name" json:"name"`
}

func (Users) TableName() string {
	return "users"
}

type ResponseModel struct {
	TestStruct
	UserName string `gorm:"column:user_name" json:"created_by"`
}

func GetTest(c *gin.Context) {
	var model []ResponseModel
	dateFrom := c.Query("date_from")
	dateTo := c.Query("date_to")

	query := db.Table("test t").Select("t.*, u.name as user_name").
		Joins("join users u on u.id = t.created_by")

	if dateTo != "" && dateFrom != "" {
		query = query.Where("t.created_at >= ? and t.	created_at <= ?", dateFrom, dateTo)
	}

	err := query.Scan(&model).Error
	if err != nil {
		log.Println(err)
		return
	}
	c.JSON(205, model)
}

package main

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Data struct {
	gorm.Model
	Name  string
	Price uint64
}

func main() {
	initdb()
	r := gin.Default()

	r.GET("/", read)
	r.POST("/add/:name/:price", create)
	r.PUT("/update/:name", update)
	r.DELETE("/delete/:id", delete)

	r.Run()

}

//init db
func initdb() *gorm.DB {
	db, e := gorm.Open(sqlite.Open("data.db"), &gorm.Config{})
	if e != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Data{})
	return db
}

//CRUD
func delete(c *gin.Context) {
	db := initdb()
	id := c.Param("id")
	db.Delete(&Data{}, id)
	c.JSON(200, gin.H{
		"info": "delete id " + id,
	})

}

func create(c *gin.Context) {
	name := c.Param("name")
	price, err := strconv.ParseUint(c.Param("price"), 10, 64)
	if err != nil {
		// panic("Error Convert")
		c.JSON(503, gin.H{
			"error": "error convert",
		})
		return

	}
	fmt.Println(price)

	db := initdb()
	db.Create(&Data{
		Name:  name,
		Price: price,
	})

	c.JSON(200, gin.H{
		"info": "data add",
	})
}

func read(c *gin.Context) {
	db := initdb()
	var data []Data
	db.Find(&data)

	c.JSON(200, data)

}

func update(c *gin.Context) {
	name := c.Param("name")
	db := initdb()
	db.Model(&Data{}).Where("ID = ?", 2).Update("name", name)
	c.JSON(200, gin.H{
		"update": "update name 2",
	})

}

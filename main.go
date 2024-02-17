package main

import (
	"learn_golang/module/restaurant/transport/ginrestaurant"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Restaurant struct {
	Id   int    `json:"id" gorm:"column:id;"`
	Name string `json:"name" gorm:"column:name"`
	Addr string `json:"addr" gorm:"column:addr"`
}

func (Restaurant) TableName() string { return "restaurants" }

type RestaurantUpdate struct {
	Name *string `json:"name" gorm:"column:name"`
	Addr *string `json:"addr" gorm:"column:addr"`
}

func (RestaurantUpdate) TableName() string { return Restaurant{}.TableName() }

func main() {
	// dsn := os.Getenv("CONNECTION_STRING")
	dsn := "food_delivery:buonlamchiemoi147@tcp(127.0.0.1:3307)/food_delivery?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(db)
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"data": "pong",
		})
	})

	//POST
	v1 := r.Group("/v1")
	restaurants := v1.Group("/restaurants")

	restaurants.POST("", ginrestaurant.CreateRestaurant(db))

	restaurants.GET("/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}
		var data Restaurant

		db.Where("id=?", id).First(&data)
		c.JSON(http.StatusOK, gin.H{
			"data": data,
		})
	})

	restaurants.GET("", func(c *gin.Context) {
		var data []Restaurant

		type Paging struct {
			Page  int `json:"page" form:"page"`
			Limit int `json:"limit" form:"limit"`
		}
		var pagingData Paging

		if err := c.ShouldBind(&pagingData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}
		if pagingData.Page <= 0 {
			pagingData.Page = 1
		}
		if pagingData.Limit <= 0 {
			pagingData.Limit = 5
		}

		db.Offset((pagingData.Page - 1) * pagingData.Limit).
			Order("id desc").Limit(pagingData.Limit).
			Find(&data)

		c.JSON(http.StatusOK, gin.H{
			"data": data,
		})
	})

	restaurants.PATCH("/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}
		var data RestaurantUpdate
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}
		db.Where("id=?", id).Updates(&data)
		c.JSON(http.StatusOK, gin.H{
			"data": data,
		})
	})

	restaurants.DELETE("/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		db.Table(Restaurant{}.TableName()).Where("id=?", id).Delete(nil)
		c.JSON(http.StatusOK, gin.H{
			"data": 1,
		})
	})

	r.Run()

	// if err := db.Create(&newRestaurant).Error; err != nil {
	// 	log.Println(err)
	// }
	// newRestaurant := Restaurant{Name: "ABC", Addr: "Doan xem"}

	// var myRestaurant Restaurant
	// if err := db.Where("id = ?", 2).First(&myRestaurant).Error; err != nil {
	// 	log.Println(err)
	// }

	// myRestaurant.Name = "Duy"
	// if err := db.Where("id = ?", 2).Updates(&myRestaurant).Error; err != nil {
	// 	log.Println(err)
	// }
	// if err := db.Table(Restaurant{}.TableName()).Where("id = ?", 1).Delete(nil).Error; err != nil {
	// 	log.Println(err)
	// }

	// //update rong
	// newName := ""
	// updateData := RestaurantUpdate{Name: &newName}
	// if err := db.Where("id = ?", 2).Updates(&updateData).Error; err != nil {
	// 	log.Println(err)
	// }
}

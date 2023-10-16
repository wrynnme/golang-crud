package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type (
	User struct {
		username string `json:"username"`
		password string `json:"password"`
		fullName string `json:"fullName"`
	}
)

func getAll(c echo.Context) error {
	dsn := "root:@tcp(127.0.0.1:3306)/golang-crud?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	result := []map[string]interface{}{}
	db.Table("users").Find(&result)

	return c.JSON(http.StatusOK, result)
}

func getByID(c echo.Context) error {
	dsn := "root:@tcp(127.0.0.1:3306)/golang-crud?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	id := c.Param("id") // DATA FROM PARAMETER

	result := map[string]interface{}{}
	db.Table("users").Where("id = ?", id).Find(&result)

	return c.JSON(http.StatusOK, result)
}

func create(c echo.Context) (err error) {
	dsn := "root:@tcp(127.0.0.1:3306)/golang-crud?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	u := new(User)

	if err = c.Bind(u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	user := User{username: u.username, password: u.password, fullName: u.fullName}

	result := db.Create(&user)

	return c.JSON(http.StatusOK, result)
}

func main() {
	e := echo.New()

	e.GET("/", getAll)
	e.GET("/:id", getByID)
	e.POST("/", create)

	e.Logger.Fatal(e.Start(":1323"))
}

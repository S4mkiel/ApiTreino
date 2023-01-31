package main

import (
	"log"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type User struct {
	gorm.Model
	Username string `json:"omitempty"`
	Name     string
	Age      uint8
	CompanyID uint `json:"ForeignKey:CompanyRefer"`
	CompanyRefer Company `json:"AssociationForeignKey:CompanyID"`
}

type Company struct {
	gorm.Model
	Name    string `json:"omitempty"`
	Andress string
}

func main() {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		log.Fatal("Error to open database",err)
	}
	if err := db.AutoMigrate(&User{}, &Company{}).Error; err != nil {
		log.Fatal("Error AutoMigrate",err)
	}
	defer db.Close()

	app := fiber.New()


	app.Get("/users", func (c *fiber.Ctx){
		var users User
		db.Find(&users)
		c.JSON(users)
	})

	app.Get("/users/:id", func(c *fiber.Ctx){
		id := c.Params("id")
  	var user User
  	db.First(&user, id)
  	c.JSON(user)
	})

	app.Post("/users", func (c *fiber.Ctx){
		user := new(User)
		if err := c.BodyParser(user); err != nil{
			c.Status(503).Send(err)
			return
	}
		db.Create(&user)
		c.JSON(user)
	})
}
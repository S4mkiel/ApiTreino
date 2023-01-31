package main

import (
	"log"
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type User struct {
	gorm.Model
	Username string `json:"username,omitempty" gorm:"unique_index"`
	Name     string
	Age      uint8
	CompanyID uint `json:"companyid,omitempty" gorm:"ForeignKey:CompanyRefer"`
	CompanyRefer Company `json:"comapnyrefer,omitempty" gorm:"ForeignKey:CompanyID;AssociationForeignKey:ID"`
}

type Company struct {
	gorm.Model
	Name    string `json:"name,omitempty" gorm:"unique_index"`
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


	app.Get("/users", func(c *fiber.Ctx) error {
		var users []User
		db.Find(&users)
		return c.JSON(users)
	})

	app.Get("/users/:id", func(c *fiber.Ctx) error {
		var user User
		db.First(&user, c.Params("id"))
		return c.JSON(user)
	})

	app.Post("/users", func(c *fiber.Ctx) error {
		var user User
		if err := c.BodyParser(&user); err != nil {
			return err
		}
		db.Create(&user)
		return c.JSON(user)
	})
	
	app.Listen(":3000")
}
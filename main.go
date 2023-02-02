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
	CompanyRefer Company `json:"companyrefer,omitempty" gorm:"ForeignKey:CompanyID;AssociationForeignKey:ID"`
}

type Company struct {
	gorm.Model
	Name    string `json:"name,omitempty" gorm:"unique_index"`
	Andress string
}

func main() {
	db, err := gorm.Open("sqlite3", "ApiTreino.db")
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
		if err :=db.Create(&user).Error; err != nil{
			return err
		}
		return c.JSON(user)
	})

	app.Put("/users/:id", func(c *fiber.Ctx) error {
		var user User
		if err := c.BodyParser(&user); err != nil {
			return err
		}
		if err := db.Model(&user).Updates(user).Error; err != nil{
			return err
		}
		return c.JSON(user)
	})

	app.Delete("/users/:id", func(c *fiber.Ctx) error {
		var user User
		db.First(&user, c.Params("id"))
		if err := db.Delete(&user).Error; err != nil{
			return err
		}
		return c.JSON(user)
	})

	app.Get("/companies", func(c *fiber.Ctx) error {
		var companies []Company
		db.Find(&companies)
		return c.JSON(companies)
	})
	
	app.Get("/companies/:id", func(c *fiber.Ctx) error {
		var company Company
		db.First(&company, c.Params("id"))
		return c.JSON(company)
	})
	
	app.Post("/companies", func(c *fiber.Ctx) error {
		var company Company
		if err := c.BodyParser(&company); err != nil {
			return err
		}
		if err:= db.Create(&company).Error; err != nil{
			return err
		}
		return c.JSON(company)
	})

	app.Put("/companies/:id", func(c *fiber.Ctx) error {
		var company Company
		if err := c.BodyParser(&company); err != nil {
			return err
		}
		if err := db.Model(&company).Updates(company).Error; err != nil{
			return err
		}
		return c.JSON(company)
	})

	app.Delete("/companies/:id", func(c *fiber.Ctx) error {
		var company Company
		db.First(&company, c.Params("id"))
		if err := db.Delete(&company).Error; err != nil{
			return err
		}
		return c.JSON(company)
	})

	app.Listen(":3000")
}
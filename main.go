package main

import (
	"fmt"
	"gowork/database"
	"gowork/routes"
	"log"

	m "gowork/models"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func initDatabase() {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=Local",
		"root",
		"",
		"127.0.0.1",
		"3306",
		"user",
	)
	var err error
	database.DBConn, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Println("Database connected!")
	// log.Println(db)

	database.DBConn.AutoMigrate(&m.Users{})
	fmt.Println("Migrated DB")
}

func main() {
	app := fiber.New()
	validate := validator.New()

	routes.UserRoute(app)

	initDatabase()

	// Provide a minimal config
	app.Use(basicauth.New(basicauth.Config{
		Users: map[string]string{
			"testgo": "772565",
		},
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("WELCOME TO PROJECT")
	})

	// GET http://example.com/user/fenny
	app.Get("/user/:name", func(c *fiber.Ctx) error {
		name := c.Params("name") // "fenny"
		log.Println(name)
		return c.Status(200).JSON(name)

		// ...
	})

	// Parameters
	app.Get("/user/:name/books/:title", func(c *fiber.Ctx) error {
		fmt.Fprintf(c, "%s\n", c.Params("name"))
		fmt.Fprintf(c, "%s\n", c.Params("title"))
		return nil
	})

	api := app.Group("/api") // /api
	v1 := api.Group("/v1")   // /api/v1
	v1.Get("/list", func(c *fiber.Ctx) error {
		return c.SendString("api/v1/list")
	}) // /api/v1/list
	v1.Get("/user", func(c *fiber.Ctx) error {
		return c.SendString("api/v1/user")
	}) // /api/v1/user

	v2 := api.Group("/v2") // /api/v2
	v2.Get("/list", func(c *fiber.Ctx) error {
		return c.SendString("api/v2/user")
	}) // /api/v2/list
	v2.Get("/user", func(c *fiber.Ctx) error {
		return c.SendString("api/v2/user")
	}) // /api/v2/user

	v1.Post("/adduser", func(c *fiber.Ctx) error {
		//Connect to database
		type User struct {
			Num      int    `json:"num" validate:"required,number"`
			Name     string `json:"name" validate:"required,min=3,max=32"`
			IsActive *bool  `json:"isactive" validate:"required"`
			Email    string `json:"email,omitempty" validate:"required,email,min=3,max=32"`
		}
		user := new(User)

		if err := c.BodyParser(&user); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": err.Error(),
			})
		}
		errors := validate.Struct(user)
		if errors != nil {
			return c.Status(fiber.StatusBadRequest).JSON(errors.Error())
		}

		//Do something else here

		//Return user
		return c.JSON(user)
	})

	app.Listen(":3000")
}

package controllers

import (
	"log"
	"strings"

	"gowork/database"
	m "gowork/models"

	"github.com/gofiber/fiber/v2"
)

func TestLogin(c *fiber.Ctx) error {
	p := new(m.Person)

	if err := c.BodyParser(p); err != nil {
		return err
	}

	log.Println(p.Name) // john
	log.Println(p.Pass) // doe
	return c.SendString("Login Success")
}

func GetDogs(c *fiber.Ctx) error {
	db := database.DBConn
	var dogs []m.Dogs

	db.Find(&dogs)
	return c.Status(200).JSON(dogs)
}

func GetDog(c *fiber.Ctx) error {
	db := database.DBConn
	search := strings.TrimSpace(c.Query("search"))
	var dog []m.Dogs

	result := db.Find(&dog, "dog_id = ?", search)

	// returns found records count, equals `len(users)
	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}
	return c.Status(200).JSON(&dog)
}

func AddDog(c *fiber.Ctx) error {
	db := database.DBConn
	var dog m.Dogs

	if err := c.BodyParser(&dog); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	db.Create(&dog)
	return c.Status(201).JSON(dog)
}
func UpdateDog(c *fiber.Ctx) error {
	db := database.DBConn
	var dog m.Dogs
	id := c.Params("id")

	if err := c.BodyParser(&dog); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	db.Where("id = ?", id).Updates(&dog)
	return c.Status(200).JSON(dog)
}

func RemoveDog(c *fiber.Ctx) error {
	db := database.DBConn
	var dog m.Dogs
	id := c.Params("id")

	result := db.Delete(&dog, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.SendStatus(200)
}

func GetUsers(c *fiber.Ctx) error {
	db := database.DBConn
	var users []m.Users

	db.Find(&users)
	return c.Status(200).JSON(users)
}

func GetUser(c *fiber.Ctx) error {
	db := database.DBConn
	search := strings.TrimSpace(c.Query("search"))
	var users []m.Users

	result := db.Find(&users, "employee_id = ?", search)

	// returns found records count, equals `len(users)
	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}
	return c.Status(200).JSON(&users)
}

func AddUser(c *fiber.Ctx) error {
	db := database.DBConn
	var users m.Users

	if err := c.BodyParser(&users); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	db.Create(&users)
	return c.Status(201).JSON(users)
}

func UpdateUser(c *fiber.Ctx) error {
	db := database.DBConn
	var users m.Users
	id := c.Params("id")

	if err := c.BodyParser(&users); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	db.Where("id = ?", id).Updates(&users)
	return c.Status(200).JSON(users)
}

func RemoveUser(c *fiber.Ctx) error {
	db := database.DBConn
	var users m.Users
	id := c.Params("id")

	result := db.Delete(&users, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.SendStatus(200)
}

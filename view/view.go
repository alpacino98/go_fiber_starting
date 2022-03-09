package view

import (
	"github.com/alpacino98/go_fiber_starting/database"
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
)

type View struct {
	gorm.Model
	Title  string `json:"title"`
	Author string `json:"author"`
	Rating int    `json:"rating"`
}

func GetViews(c *fiber.Ctx) error {
	db := database.DBConn
	var views []View
	db.Find(&views)
	return c.JSON(views)
	// c.Send("All views")
}

func GetView(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn
	var view View
	db.Find(&view, id)
	return c.JSON(view)
}

func NewView(c *fiber.Ctx) error {
	db := database.DBConn
	view := new(View)

	err := c.BodyParser(view)
	if err != nil {
		return c.Status(500).SendString("Invalid body")
	}

	// view.Title = "1984"
	// view.Author = "Alpha"
	// view.Rating = 15
	db.Create(&view)
	return c.JSON(view)
}

func DeleteView(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn
	var view View

	db.First(&view, id)
	if view.Title == "" {
		return c.Status(500).SendString("View not found.")
	}

	db.Delete(&view)
	return c.SendString("Deleted view successfuly.")
}

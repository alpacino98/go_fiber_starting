package view

import (
	"github.com/alpacino98/go_fiber_starting/database"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
)

type View struct {
	gorm.Model
	Title  string `json:"title"`
	Author string `json:"author"`
	Rating int    `json:"rating"`
}

func GetViews(c *fiber.Ctx) {
	db := database.DBConn
	var views []View
	db.Find(&views)
	c.JSON(views)
	// c.Send("All views")
}

func GetView(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn
	var view View
	db.Find(&view, id)
	c.JSON(view)
}

func NewView(c *fiber.Ctx) {
	db := database.DBConn
	view := new(View)

	err := c.BodyParser(view)
	if err != nil {
		c.Status(500).Send(err)
	}

	// view.Title = "1984"
	// view.Author = "Alpha"
	// view.Rating = 15
	db.Create(&view)
	c.JSON(view)
}

func DeleteView(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn
	var view View

	db.First(&view, id)
	if view.Title == "" {
		c.Status(500).Send("View not found.")
		return
	}

	db.Delete(&view)
	c.Send("Deleted view successfuly.")
}

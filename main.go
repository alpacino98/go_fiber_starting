package main

import (
	"fmt"

	"github.com/alpacino98/go_fiber_starting/database"
	"github.com/alpacino98/go_fiber_starting/view"
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/views", view.GetViews)
	app.Get("/api/v1/view/:id", view.GetView)
	app.Post("/api/v1/view", view.NewView)
	app.Delete("/api/v1/view/:id", view.DeleteView)
}

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "view.db")

	if err != nil {
		panic("Failed to connect to database.")
	}
	fmt.Println("Database successfully opened.")

	database.DBConn.AutoMigrate(&view.View{})
	fmt.Println("Database migrated")
}

func main() {
	app := fiber.New()
	initDatabase()
	defer database.DBConn.Close()

	// app.Use(cors.New())
	setupRoutes(app)

	app.Listen(":3000")
}

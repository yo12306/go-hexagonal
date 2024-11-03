package main

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/yo12306/go-hexagonal/adapters"
	"github.com/yo12306/go-hexagonal/core"
)

func main() {
	app := fiber.New()

	db, err := gorm.Open(sqlite.Open("orders.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	orderRepo := adapters.NewGormOrderRepository(db)

	orderService := core.NewOrderService(orderRepo)

	orderHandler := adapters.NewHttpOrderHandler(orderService)

	app.Post("/order", orderHandler.CreateOrder)

	db.AutoMigrate(&core.Order{})

	app.Listen(":8000")
}
